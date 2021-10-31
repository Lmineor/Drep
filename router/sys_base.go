package router

import (
	"github.com/drep/api/dev"
	"github.com/drep/api/v1"
	"github.com/gin-gonic/gin"
)

func InitBaseRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	BaseRouter := Router.Group("base")
	{
		BaseRouter.POST("login", v1.Login)
		BaseRouter.POST("register", v1.Register)
	}
	return BaseRouter
}

func InitPopRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	BaseRouter := Router.Group("dev")
	{
		BaseRouter.GET("pop", dev.Pop)
	}
	return BaseRouter
}
