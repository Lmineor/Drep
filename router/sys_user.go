package router

import (
	"github.com/drep/api/v1"
	"github.com/drep/middleware"
	"github.com/gin-gonic/gin"
)

func InitUserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("user").Use(middleware.OperationRecord())
	{
		UserRouter.POST("register", v1.Register)                 // 用户注册账号
		UserRouter.POST("changePassword", v1.ChangePassword)     // 用户修改密码
		UserRouter.GET("getUserList", v1.GetUserList)            // 分页获取用户列表
		UserRouter.GET("getAllUserList", v1.GetAllUserList)      // 超级管理员分页获取用户列表GetAllUserList
		UserRouter.POST("setUserAuthority", v1.SetUserAuthority) // 设置用户权限
		UserRouter.DELETE("deleteUser", v1.DeleteUser)           // 删除用户
		UserRouter.PUT("setUserInfo", v1.SetUserInfo)            // 设置用户信息
		UserRouter.POST("resetUserPassword", v1.ResetPassword)   //重置用户密码
	}
}
