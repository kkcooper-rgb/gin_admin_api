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
	router.GET("/api/failed", api.Failed)

	router.POST("/api/sysMenu/add", api.CreateSysMenu)
	router.GET("/api/sysMenu/list", api.GetSysMenuList)
	router.GET("/api/sysMenu/info", api.GetSysMenu)
	router.PUT("/api/sysMenu/update", api.UpdateSysMenu)
	router.DELETE("/api/sysMenu/delete", api.DeleteSysMenu)

	router.POST("/api/sysRole/add", api.CreateSysRole)
	router.GET("/api/sysRole/list", api.GetSysRoleList)
	router.GET("/api/sysRole/info", api.GetSysRole)
	router.PUT("/api/sysRole/update", api.UpdateSysRole)
	router.DELETE("/api/sysRole/delete", api.DeleteSysRole)
}
