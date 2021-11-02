package initial

import (
	"github.com/drep/global"
	"github.com/drep/middleware"
	"github.com/drep/router"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// 初始化总路由

func Routers() *gin.Engine {
	var Router = gin.Default()
	//Router.StaticFS(global.CONFIG.Local.Path, http.Dir(global.CONFIG.Local.Path)) // 为用户头像和文件提供静态地址
	// Router.Use(middleware.LoadTls())  // 打开就能玩https了
	global.LOG.Info("use middleware logger")
	// 跨域
	Router.Use(middleware.Cors()) // 如需跨域可以打开
	global.LOG.Info("use middleware cors")
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	global.LOG.Info("register swagger handler")
	// 方便统一添加路由组前缀 多服务器上线使用
	PublicGroup := Router.Group("")
	{
		router.InitBaseRouter(PublicGroup) // 注册基础功能路由 不做鉴权
		if global.CONFIG.System.Env == "develop" {
			router.InitPopRouter(PublicGroup)
		}
	}
	PrivateGroup := Router.Group("")
	PrivateGroup.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		router.InitJwtRouter(PrivateGroup)
		router.InitProjectRouter(PrivateGroup)
		router.InitDpRouter(PrivateGroup)
		router.InitUserRouter(PrivateGroup)               // 注册用户路由
		router.InitSysOperationRecordRouter(PrivateGroup) // 操作记录
	}
	global.LOG.Info("router register success")
	return Router
}
