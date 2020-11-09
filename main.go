package main

import(
	"flag"
	"os"

	"github.com/gin-gonic/gin"
	
	"ren_chat/router"
	"ren_chat/config"
)

//获取终端端口号
func getAddr()string{
	cnf := config.NewConfig("global","json")
	host := cnf.GetString("server.host")
	port := cnf.GetString("server.port")

	flag.StringVar(&host,"h",host,"监听地址")
	flag.StringVar(&port,"p",port,"监听端口")
	flag.Parse()
	return host+":"+port
}

//入口函数
func main(){
	//设置运行环境
	os.Setenv("REGIN_RUNMODE","dev")
	//获取命令行端口
	addr := getAddr()
	svr := gin.Default()
	svr.StaticFile("/favicon.ico","./public/static/imgs/favicon.ico")
	//注册路由
	router.Register(svr)
	//启动一个http服务器
	svr.Run(addr)
}