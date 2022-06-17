package handlers

import (
	"ecommerce/token"
	db "ecommerce/db/sqlc"
	"ecommerce/middleware"
	"ecommerce/models"

	"fmt"

	"github.com/gin-gonic/gin"

	ginSwagger "github.com/swaggo/gin-swagger"

	// gin-swagger middleware
	swaggerFiles "github.com/swaggo/files"
)

// swagger embed files
type Server struct {
	config     models.Configuration
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

func NewServer(config models.Configuration, store db.Store) (*Server, error) {

	fmt.Println("TokenSymmetricKey = ", config.TokenSymmetricKey)
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	server.setupRouter()
	return server, nil
}
func (server *Server) setupRouter() {
	router := gin.Default()
	router.Use(middleware.CORSMiddleware())
	v1 := router.Group("/api/v1")
	{
		// v1.Use(middleware.AuthMiddleware(server.tokenMaker, server.store))
		// v1.Use(middleware.AuditMiddleware(server.store))

		_users := v1.Group("/users")
		{
			_users.GET("", server.getUsers)
		}
		v1.POST("/register", server.createUser)
		v1.POST("/login", server.loginUser)
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	server.router = router

}

// Start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
