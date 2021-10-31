package router

import (
	"github.com/drep/api/v1"
	"github.com/drep/middleware"
	"github.com/gin-gonic/gin"
)

func InitJwtRouter(Router *gin.RouterGroup) {
	JwtRouter := Router.Group("jwt").Use(middleware.OperationRecord())
	{
		JwtRouter.POST("jsonInBlacklist", v1.JsonInBlacklist) // jwt加入黑名单
	}
}
