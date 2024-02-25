package main

import (
	"github.com/576690/qy8/backend/core"
	"github.com/576690/qy8/backend/global"
	"github.com/576690/qy8/backend/router"
)

func main() {
	//读取配置文件
	core.InitConf()
	//fmt.Println(global.Config)

	//初始化日志
	global.Log = core.InitLogger()
	global.Log.Warnln("嘿嘿")
	global.Log.Error("哈哈")
	global.Log.Infoln("嘻嘻")

	//链接数据库
	global.DB = core.InitGorm()
	//fmt.Println(global.DB)

	core.InitLogger()

	router := router.InitRouter()

	addr := global.Config.System.Addr()
	global.Log.Infof("qy8运行在: %s", addr)
	router.Run(addr)

	// server := http.Server{
	// 	Addr: "127.0.0.1:8080",
	// }

	// router.Router()

	// error := server.ListenAndServe()
	// if error != nil {
	// 	log.Println(error)
	// }
	// // 启动服务器

}
