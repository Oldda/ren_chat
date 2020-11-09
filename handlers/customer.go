package handlers

import(
	"log"
  	ws_svr "ren_chat/server/ws"
  	"github.com/gin-gonic/gin"
)

var customerServer *ws_svr.WsServer

//客服系统
func Customer(ctx *gin.Context){
	//开启ws服务器
  	customerServer = ws_svr.NewWsServer()
  	go customerServer.Run()

  	conn := ws_svr.NewWsConn(customerServer)

  	conn.OnOpen(ctx.Writer,ctx.Request,func(conn *ws_svr.WsConn){
    	log.Println("open")
  	})

  	conn.OnMessage(func(conn *ws_svr.WsConn,msg []byte){
    	log.Println("message")
  	})

  	conn.OnClose(func(conn *ws_svr.WsConn,err error){
    	log.Println("close")
  	})
}