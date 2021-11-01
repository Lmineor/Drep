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
		//project.PUT("updateProject", v1.UpdateProject)
		//project.GET("listAllProjects", v1.ListAllProjects) // 管理员获取所有项目
		//project.DELETE("deleteProject", v1.DeleteProject)  // 管理员获取所有项目
		//UserRouter.POST("changePassword", v1.ChangePassword)     // 用户修改密码
		//UserRouter.POST("getUserList", v1.GetUserList)           // 分页获取用户列表
		//UserRouter.POST("setUserAuthority", v1.SetUserAuthority) // 设置用户权限
		//UserRouter.DELETE("deleteUser", v1.DeleteUser)           // 删除用户
		//UserRouter.PUT("setUserInfo", v1.SetUserInfo)            // 设置用户信息
	}
}
