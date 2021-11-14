package router

import (
	"github.com/drep/api/v1"
	"github.com/drep/middleware"
	"github.com/gin-gonic/gin"
)

func InitDpRouter(Router *gin.RouterGroup) {
	dp := Router.Group("dp").Use(middleware.OperationRecord())
	{
		dp.GET("getDpDetail", v1.GetDpDetail)
		dp.POST("createDp", v1.CreateDailyReport)
		dp.PUT("updateDp", v1.UpdateDailyReport)
		dp.GET("listAllDps", v1.ListAllDailyReports)
		dp.GET("listAllDpsSub", v1.ListAllDailyReportsUnderAdmin)
		dp.GET("listDps", v1.ListDps)
		dp.DELETE("deleteDp", v1.DeleteDailyReport)
	}
}
