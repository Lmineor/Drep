package router

import (
	"github.com/drep/api/v1"
	"github.com/drep/middleware"
	"github.com/gin-gonic/gin"
)

func InitDpRouter(Router *gin.RouterGroup) {
	dp := Router.Group("dp").Use(middleware.OperationRecord())
	{
		dp.POST("createDp", v1.CreateDailyReport) // 管理员创建项目
		dp.PUT("updateDpDetail", v1.UpdateDailyReport)
		dp.GET("listAllDps", v1.ListAllDailyReports) // 管理员获取所有项目
		dp.GET("listDps", v1.ListDps)                // 用户获取自己的所有日报
		dp.DELETE("deleteDp", v1.DeleteDailyReport)  // 管理员获取所有项目
	}
}
