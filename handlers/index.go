package handlers

import(
	"github.com/gin-gonic/gin"
)

func Index(ctx *gin.Context){
	ctx.JSON(200,gin.H{
		"code":200,
		"msg":"success",
	})
}