package commonFunctions

import (
	aerrors "ecommerce/dev.error"
	//"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

//Throwing custom error to End point
func SendError(code int, ctx *gin.Context, err error, customError *aerrors.Error) {
	//fmt.Println(err)
	aerrors.SetActualError(err, customError)
	//fmt.Println("aerr--", *customError)
	ctx.AbortWithStatusJSON(http.StatusNotFound, customError)
}
