package router

import (
	"github.com/drep/api/v1"
	"github.com/drep/middleware"
	"github.com/gin-gonic/gin"
)

func InitProjectRouter(Router *gin.RouterGroup) {
	project := Router.Group("project").Use(middleware.OperationRecord())
	{
		project.POST("createProject", v1.CreateProject) // 管理员创建项目
		project.GET("getProjectDetail", v1.GetProjectDetail)
		project.PUT("updateProject", v1.UpdateProject)
		project.GET("listAllProjects", v1.ListAllProjects) // 管理员获取所有项目
		project.GET("listProjects", v1.ListProjects)
		project.DELETE("deleteProject", v1.DeleteProject)
		project.DELETE("deleteProjectByUUIDs", v1.DeleteProjectByUUIDs)
		project.GET("exportProjectXls", v1.ExportProjectXls)
	}
}
