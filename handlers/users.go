package handlers

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"
	"time"

	"ecommerce/commonFunctions"
	db "ecommerce/db/sqlc"
	aerrors "ecommerce/dev.error"
	"ecommerce/mail"
	"ecommerce/types"

	"ecommerce/util"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type newPassword struct {
	Password string `json:"password"`
}
type requestForgetPassword struct {
	AuthCode string `json:"auth_code"`
	Email    string `json:"email"`
}

// structure that gets user request from end points
type createUserRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
	FullName string `json:"full_name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
}

// structure to store data that is to be sent to end points
type userResponse struct {
	ID                  int32     `json:"id"`
	Username            string    `json:"username"`
	FullName            string    `json:"full_name"`
	Email               string    `json:"email"`
	PasswordChangedDate time.Time `json:"password_changed_date"`
	CreatedDate         time.Time `json:"created_date"`
}

// translate the user to userResponse for removing password
func translateUser(user db.User) userResponse {
	return userResponse{
		ID:                  user.ID,
		Username:            user.Username,
		FullName:            user.FullName,
		Email:               user.Email,
		PasswordChangedDate: user.PasswordChangedDate,
		CreatedDate:         user.CreatedDate,
	}
}

type emailResponse struct {
	ID          int32     `json:"id"`
	Username    string    `json:"username"`
	Email       string    `json:"email"`
	CreatedDate time.Time `json:"created_date"`
}

func transUser(user db.User) emailResponse {
	return emailResponse{
		ID:          user.ID,
		Username:    user.Username,
		Email:       user.Email,
		CreatedDate: user.CreatedDate,
	}
}

// function to get user details and send it to end point as slice of json
func (server *Server) getUsers(ctx *gin.Context) {

	//get user
	users, err := server.store.GetUsers(ctx)
	if err != nil {
		// set actual error
		commonFunctions.SendError(http.StatusNotFound, ctx, err, aerrors.ErrUserNotFound)
		return
	}

	// slice variable to store user details
	responses := []userResponse{}

	// converting users to userResponse
	for i := 0; i < len(users); i++ {

		responses = append(responses, translateUser(users[i]))
	}
	// sending user details to end point as slice of json
	ctx.JSON(http.StatusOK, responses)
}

// function to create users and store in detail
func (server *Server) createUser(ctx *gin.Context) {

	var req createUserRequest

	//bind
	if err := ctx.ShouldBindJSON(&req); err != nil {
		commonFunctions.SendError(http.StatusBadRequest, ctx, err, aerrors.ErrUserBinding)
		return
	}

	// converting plain text to hashed password
	hashedPassword, err := util.HashPassword(req.Password)
	hashedFullName, err := util.HashPassword(req.FullName)
	if err != nil {

		// set actual error
		commonFunctions.SendError(http.StatusInternalServerError, ctx, err, aerrors.ErrHashingFailed)
		return
	}
	//Compare the submitted captcha value with the one in memory
	arg := db.CreateUserParams{
		Username: req.Username,
		Password: hashedPassword,
		FullName: hashedFullName,
		Email:    req.Email,
	}
	// To store data into database
	user, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		commonFunctions.SendError(http.StatusInternalServerError, ctx, err, aerrors.ErrUserInsertion)
		return
	}
	//variable to store only selected details of user
	rsp := translateUser(user)

	//sendin gto endpoint as json to endpoint
	ctx.JSON(http.StatusOK, rsp)
	// variable that stores  user details to be stored in database

}

/* login api */
type loginUserRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
	AuthCode string `json:"auth_code"`
}

type loginUserResponse struct {
	SessionID             uuid.UUID    `json:"session_id"`
	AccessToken           string       `json:"access_token"`
	AccessTokenExpiresAt  time.Time    `json:"access_token_expires_at"`
	RefreshToken          string       `json:"refresh_token"`
	RefreshTokenExpiresAt time.Time    `json:"refresh_token_expires_at"`
	User                  userResponse `json:"user"`
}

func (server *Server) loginUser(ctx *gin.Context) {
	fmt.Println("request received")
	var req loginUserRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		fmt.Println(err)
		commonFunctions.SendError(http.StatusBadRequest, ctx, err, aerrors.ErrUserBinding)
		return
	}
	if req.AuthCode == "" {
		user, err := server.store.GetUser(ctx, req.Username)
		if err != nil {
			fmt.Println(err)
			commonFunctions.SendError(http.StatusNotFound, ctx, err, aerrors.ErrUserNotFound)
			return
		}

		err = util.CheckPassword(req.Password, user.Password)
		if err != nil {
			fmt.Println(err)
			commonFunctions.SendError(http.StatusUnauthorized, ctx, err, aerrors.ErrWrongPassword)
			return
		}
		authCode := util.RandomNumericString(6)
		err = server.store.AuthCodeUser(ctx, db.AuthCodeUserParams{
			AuthCode: sql.NullString{String: authCode, Valid: true},
			ID:       user.ID,
		})
		if err != nil {

			fmt.Println(err)
			//change the error response message
			commonFunctions.SendError(http.StatusInternalServerError, ctx, err, aerrors.ErrAuthCodeNotInserted)
			return
		}
		mail.SendMail(user.Email, authCode)
		//error for authlogin

		ctx.JSON(http.StatusOK, translateUser(user))
	} else {
		user, err := server.store.GetUser(ctx, req.Username)
		if err != nil {
			fmt.Println(err)
			commonFunctions.SendError(http.StatusNotFound, ctx, err, aerrors.ErrUserNotFound)
			return
		}

		err = util.CheckPassword(req.Password, user.Password)
		if err != nil {
			fmt.Println(err)
			commonFunctions.SendError(http.StatusUnauthorized, ctx, err, aerrors.ErrWrongPassword)
			return
		}

		if req.AuthCode != user.AuthCode.String {
			err := errors.New("OTP is not matched")
			fmt.Println(req.AuthCode)
			//Authentication not mached
			commonFunctions.SendError(http.StatusUnauthorized, ctx, err, aerrors.ErrAuthCodeNotMatched)
			return
		}
		err = server.store.AuthCodeUser(ctx, db.AuthCodeUserParams{
			AuthCode: sql.NullString{Valid: false},
			ID:       user.ID,
		})
		if err != nil {

			fmt.Println(err)
			//change the error response message
			commonFunctions.SendError(http.StatusInternalServerError, ctx, err, aerrors.ErrAuthCodeNotInserted)
			return
		}

		accessToken, accessPayload, err := server.tokenMaker.CreateToken(
			user.ID,
			user.Username,
			server.config.AccessTokenDuration,
		)
		if err != nil {

			fmt.Println(err)
			commonFunctions.SendError(http.StatusInternalServerError, ctx, err, aerrors.ErrTokenCreate)
			return
		}

		refreshToken, refreshPayload, err := server.tokenMaker.CreateToken(
			user.ID,
			user.Username,
			server.config.RefreshTokenDuration,
		)
		if err != nil {

			fmt.Println(err)
			commonFunctions.SendError(http.StatusInternalServerError, ctx, err, aerrors.ErrTokenRefreshFailed)
			return
		}

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}

		//save the sesion in the database
		rsp := loginUserResponse{
			SessionID:             accessPayload.ID,
			AccessToken:           accessToken,
			AccessTokenExpiresAt:  accessPayload.ExpiredAt,
			RefreshToken:          refreshToken,
			RefreshTokenExpiresAt: refreshPayload.ExpiredAt,
			User:                  translateUser(user),
		}
		ctx.JSON(http.StatusOK, rsp)
	}
}

type EmailResponse struct {
	// SessionID   uuid.UUID     `json:"session_id"`
	// AccessToken string        `json:"access_token"`
	User emailResponse `json:"user"`
}

func (server *Server) ForgetPWD(ctx *gin.Context) {
	var req requestForgetPassword
	if err := ctx.ShouldBindJSON(&req); err != nil {
		fmt.Println(err)
		commonFunctions.SendError(http.StatusBadRequest, ctx, err, aerrors.ErrUserBinding)
		return
	}
	if req.AuthCode == "" {

		user, err := server.store.GetUserEmail(ctx, req.Email)
		if err != nil {
			fmt.Println(err)
			commonFunctions.SendError(http.StatusNotFound, ctx, err, aerrors.ErrUserNotFound)
			return
		}
		authCode := util.RandomNumericString(4)
		err = server.store.PasswordTokenUser(ctx, db.PasswordTokenUserParams{
			PasswordVerificationToken: sql.NullString{String: authCode, Valid: true},
			ID:                        user.ID,
		})
		if err != nil {

			fmt.Println(err)
			//change the error response message
			commonFunctions.SendError(http.StatusInternalServerError, ctx, err, aerrors.ErrAuthCodeNotInserted)
			return
		}
		mail.SendMail(user.Email, authCode)
		ctx.JSON(http.StatusOK, transUser((user)))

	} else {
		user, err := server.store.GetUserEmail(ctx, req.Email)
		if err != nil {
			fmt.Println(err)
			commonFunctions.SendError(http.StatusNotFound, ctx, err, aerrors.ErrUserNotFound)
			return
		}
		// accessToken, accessPayload, err := server.tokenMaker.CreateToken(
		// 	user.ID,
		// 	user.Username,
		// 	server.config.AccessTokenDuration,
		// )
		// if err != nil {

		// 	fmt.Println(err)
		// 	commonFunctions.SendError(http.StatusInternalServerError, ctx, err, aerrors.ErrTokenCreate)
		// 	return
		// }
		if user.PasswordVerificationToken.String == req.AuthCode {
			ctx.JSON(http.StatusOK, user)
		} else {
			ctx.JSON(404, "Invalid AuthCode")
		}
		rsp := EmailResponse{
			// SessionID:   accessPayload.ID,
			// AccessToken: accessToken,
			User: transUser(user),
		}
		ctx.JSON(http.StatusOK, rsp)
	}

}

func (server *Server) ChangePWD(ctx *gin.Context) {
	user_id, err := strconv.Atoi(ctx.Param("id"))
	var password newPassword
	err = ctx.BindJSON(&password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "Bad request")
		return
	}
	// id := int32(22)
	user, err := server.store.GetUserById(ctx, int32(user_id))
	if err != nil {
		ctx.JSON(404, "User not found")
		return
	}

	err = util.CheckPassword(password.Password, user.Password)
	if err != nil {
		hashedPassword, err := util.HashPassword(password.Password)
		if err != nil {
			ctx.JSON(500, "failed to encrypt")
			return
		}

		err = server.store.ChangePasswordUser(ctx, db.ChangePasswordUserParams{
			ID:       int32(user_id),
			Password: hashedPassword,
		})
		if err != nil {
			ctx.JSON(500, "Failed to change password")
			return
		} else {
			ctx.JSON(200, "Successfully changed password")
			return
		}
	} else {
		ctx.JSON(http.StatusBadRequest, "Please type different password which doesnt match with old one")
		return
	}

}

func (server *Server) ChangePWDById(ctx *gin.Context) {
	var password newPassword
	err := ctx.BindJSON(&password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "Bad request")
		return
	}
	// id := int32(22)
	user, err := server.store.GetUserById(ctx, types.GetCurrentUserId(ctx))
	if err != nil {
		ctx.JSON(404, "User not found")
		return
	}

	err = util.CheckPassword(password.Password, user.Password)
	if err != nil {
		hashedPassword, err := util.HashPassword(password.Password)
		if err != nil {
			ctx.JSON(500, "failed to encrypt")
			return
		}

		err = server.store.ChangePasswordUser(ctx, db.ChangePasswordUserParams{
			ID:       types.GetCurrentUserId(ctx),
			Password: hashedPassword,
		})
		if err != nil {
			ctx.JSON(500, "Failed to change password")
			return
		} else {
			ctx.JSON(200, "Successfully changed password")
			return
		}
	} else {
		ctx.JSON(http.StatusBadRequest, "Please type different password which doesnt match with old one")
		return
	}

}
