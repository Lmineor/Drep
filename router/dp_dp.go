package router

import (
	"github.com/drep/api/v1"
	"github.com/drep/middleware"
	"github.com/gin-gonic/gin"
)

func InitDpRouter(Router *gin.RouterGroup) {
	dp := Router.Group("dp").Use(middleware.OperationRecord())
	{
		dp.GET("getDpDetail/:uuid", v1.GetDpDetail)
		dp.POST("createDp", v1.CreateDailyReport)
		dp.PUT("updateDpDetail/:uuid", v1.UpdateDailyReport)
		dp.GET("listAllDps", v1.ListAllDailyReports)
		dp.GET("listDps", v1.ListDps)
		dp.DELETE("deleteDp/:uuid", v1.DeleteDailyReport)
	}
}
