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
	// router.InitRouter()
	// 初始化路由
	r := router.InitRouter()

	// 启动 Gin 服务，可指定监听地址和端口，如 r.Run(":8080")，不指定默认监听 8080 端口
	if err := r.Run(); err != nil {
		// global.Log.Error("服务启动失败: %v", err)
		// 也可以选择退出程序，比如 os.Exit(1)
	}
}
