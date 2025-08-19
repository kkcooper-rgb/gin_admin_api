package main

import (
	"go_admin_api/core"
	"go_admin_api/global"
)

func main() {
	// fmt.Println("admin博客")
	// fmt.Println("配置", config.Config.System)
	// fmt.Println("日志", config.Config.Logger)
	global.Log = core.InitLogger()
	global.Log.Warnln("go的日志")
	global.Log.Error("go的日志")
	global.Log.Infof("go的日志")
}
