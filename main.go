package main

import (
	"go_admin_api/core"
	"go_admin_api/global"
	"go_admin_api/router"
)

func main() {
	// fmt.Println("admin博客")
	// fmt.Println("配置", config.Config.System)
	// fmt.Println("日志", config.Config.Logger)
	global.Log = core.InitLogger()
	// global.Log.Warnln("go的日志")
	// global.Log.Error("go的日志")
	// global.Log.Infof("mysql配置$s", config.Config.Mysql)
	// 初始化mysql
	core.MysqlInit()
	// 初始化redis
	core.RedisInit()
	router.InitRouter()
}
