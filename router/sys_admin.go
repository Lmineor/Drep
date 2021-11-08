package router

import (
	"github.com/drep/api/v1"
	"github.com/drep/middleware"
	"github.com/gin-gonic/gin"
)

func InitAdminRouter(Router *gin.RouterGroup) {
	adminRouter := Router.Group("admin").Use(middleware.OperationRecord())
	{
		adminRouter.POST("createInviteCode", v1.CreateInviteCode)
		adminRouter.GET("listAllInviteCode", v1.ListInviteCode)
		adminRouter.DELETE("deleteInviteCode", v1.DeleteInviteCode)
		adminRouter.GET("getInviteCodeById", v1.GetInviteCode)
		adminRouter.PUT("updateInviteCodeById", v1.UpdateInviteCodeById)
	}
}
