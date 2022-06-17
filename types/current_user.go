package types

import (
	"asarfi/common/token"

	"github.com/gin-gonic/gin"
)

// func GetCurrentUserId(ctx *gin.Context) int32 {
// 	payload, _ := ctx.Get("current_user")
// 	return payload.(*token.Payload).UserId
// }
func GetCurrentUserId(ctx *gin.Context) int32 {
	payload, _ := ctx.Get("current_user")
	return payload.(*token.Payload).UserId
}
