package handlers

import(
  "log"
  ws_svr "ren_chat/server/ws"
  "github.com/gin-gonic/gin"
)
//ws服务器对象
var ws *ws_svr.WsServer
//一个uid对应多个链接
var uidToConns = make(map[string]map[*ws_svr.WsConn]*ws_svr.WsConn)
//一个链接只对应一个Uid
var connToUid = make(map[*ws_svr.WsConn]string)


func ChatRoom(ctx *gin.Context){
  //用户id
  user_id := ctx.Query("user_id")
  if user_id == ""{
    ctx.JSON(NewJsonResponseData(NEED_PARAMS_LOST,gin.H{"Reminder":"请传入您的UserId"}))
    return
  }
  //开启ws服务器
  ws = ws_svr.NewWsServer()
  go ws.Run()

  client := ws_svr.NewWsConn(ws)

  client.OnOpen(ctx.Writer,ctx.Request,func(cli *ws_svr.WsConn){
    //绑定用户和链接
    bindUidAndConns(user_id,cli)
  })

  client.OnMessage(func(cli *ws_svr.WsConn,msg []byte){
    wsMessageHandler(cli,msg)
  })

  client.OnClose(func(cli *ws_svr.WsConn,err error){
    unbindUidAndConns(user_id,cli)
  })
}

//绑定用户和链接
func bindUidAndConns(uid string,cli *ws_svr.WsConn){
  if _,ok := uidToConns[uid];ok{
    uidToConns[uid][cli] = cli
    return
  }
  ctoc := make(map[*ws_svr.WsConn]*ws_svr.WsConn)
  ctoc[cli] = cli
  uidToConns[uid] = ctoc
  connToUid[cli] = uid
}

//解绑用户和链接
func unbindUidAndConns(uid string,cli *ws_svr.WsConn){
  delete(uidToConns[uid],cli)
  delete(connToUid,cli)
}

//消息处理函数
func wsMessageHandler(conn *ws_svr.WsConn,msg []byte){
  //获取在线链接列表
  log.Println(string(msg))
}