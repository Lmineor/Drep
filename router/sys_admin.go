package router

import (
	"github.com/drep/api/v1"
	"github.com/drep/middleware"
	"github.com/gin-gonic/gin"
)

func InitAdminRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("admin").Use(middleware.OperationRecord())
	{
		UserRouter.POST("createInviteCode", v1.CreateInviteCode)
		UserRouter.GET("listAllInviteCode", v1.ListInviteCode)
		UserRouter.DELETE("deleteInviteCode/:id", v1.DeleteInviteCode)

	}
}
