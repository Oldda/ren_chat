package router

import(
	"github.com/gin-gonic/gin"
	"ren_chat/handlers"
)

func Register(svr *gin.Engine){
	svr.GET("/",handlers.Index) //首页
	svr.GET("/chat",handlers.ChatRoom) //聊天室ws
	svr.GET("/customer",handlers.Customer) //客服系统
}