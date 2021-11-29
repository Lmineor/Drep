package router

import (
	"github.com/drep/api/v1"
	"github.com/drep/middleware"
	"github.com/gin-gonic/gin"
)

func InitSystemRouter(Router *gin.RouterGroup) {
	SystemRouter := Router.Group("system").Use(middleware.OperationRecord())
	{
		SystemRouter.GET("getSystemConfig", v1.GetSystemConfig) // 获取配置文件内容
		SystemRouter.PUT("setSystemConfig", v1.SetSystemConfig) // 设置配置文件内容
		SystemRouter.GET("getServerInfo", v1.GetServerInfo)     // 获取服务器信息
		SystemRouter.POST("reloadSystem", v1.ReloadSystem)      // 重启服务
	}
}
