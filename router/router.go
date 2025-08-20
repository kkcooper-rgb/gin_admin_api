package router

import (
	"go_admin_api/api"
	"go_admin_api/config"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	// 设置启动模式
	gin.SetMode(config.Config.System.Env)
	router := gin.New()
	// 宕机时恢复
	router.Use(gin.Recovery())
	// register注册
	register(router)
	return router
}

// 路由接口
func register(router *gin.Engine) {
	router.GET("/api/success", api.Success)
	// router.GET("/api/failed", api.Failed)
}
