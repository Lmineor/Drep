package v1

import (
	"fmt"
	"github.com/drep/core/request"
	"github.com/drep/core/response"
	"github.com/drep/global"
	"github.com/drep/model"
	"github.com/drep/service"
	"github.com/drep/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// @Tags Project
// @Summary 项目详情
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"成功"}"
// @Router /project/getProjectDetail [get]
func GetProjectDetail(c *gin.Context) {
	var reqUuid request.GetByUUID
	_ = c.ShouldBindQuery(&reqUuid)
	uuid := reqUuid.UUID
	if uuid == "" {
		response.FailWithMessage("获取项目失败，请指定项目的uuid", c)
		return
	}
	dbPj, err := service.GetProjectDetail(uuid)
	if err != nil {
		global.LOG.Info(fmt.Sprintf("error from db is: %s", err))
		response.FailWithMessage("获取项目失败, 项目不存在", c)
		return
	}
	response.OkWithDetailed(response.ProjectResponse{Project: dbPj}, "成功", c)
}

func CreateProject(c *gin.Context) {
	var pj request.Project
	_ = c.ShouldBindJSON(&pj)
	pj.UUID = utils.GenerateUUID()
	userId := getUserID(c)
	mPj := model.DpProject{Name: pj.Name, Description: pj.Description, UUID: pj.UUID, UserID: userId}
	dbPj, err := service.CreateProject(&mPj)
	if err != nil {
		response.FailWithMessage("注册项目失败，请联系管理员", c)
	} else {
		response.OkWithDetailed(response.ProjectResponse{Project: dbPj}, "项目注册成功", c)
	}
}

func ListAllProjects(c *gin.Context) {
	var total int64
	pageNum, pageSize := utils.ParsePaginateParams(c)

	list, total, err := service.ListAllProjects(pageNum, pageSize)
	if err != nil {
		response.FailWithMessage("failed to get all projects", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			PageSize: pageSize,
			Page:     pageNum,
		}, "获取成功", c)
	}
}

func ListProjects(c *gin.Context) {
	var total int64
	pageNum, pageSize := utils.ParsePaginateParams(c)
	userId := getUserID(c)
	list, total, err := service.ListProjects(userId, pageNum, pageSize)
	if err != nil {
		response.FailWithMessage("failed to get all projects", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			PageSize: pageSize,
			Page:     pageNum,
		}, "获取成功", c)
	}
}

func UpdateProject(c *gin.Context) {
	var pj model.DpProject
	_ = c.ShouldBindJSON(&pj)
	if pj.UUID == "" {
		response.FailWithMessage("获取项目失败，请指定项目的uuid", c)
		return
	}
	updatedPj, err := service.UpdateProject(&pj)
	if err != nil {
		errMsg := fmt.Sprintf("更新项目失败，错误：%s", err)
		response.FailWithMessage(errMsg, c)
	} else {
		response.OkWithDetailed(updatedPj, "更新成功", c)
	}
}

func DeleteProject(c *gin.Context) {
	var reqUuid request.GetByUUID
	_ = c.ShouldBindQuery(&reqUuid)
	uuid := reqUuid.UUID
	if uuid == "" {
		response.FailWithMessage("获取项目失败，请指定项目的uuid", c)
		return
	}
	err := service.DeleteProject(uuid)
	if err != nil {
		errMsg := fmt.Sprintf("删除项目失败，错误：%s", err)
		global.LOG.Info(errMsg)
		response.FailWithMessage("删除项目失败", c)
	} else {
		response.Ok(c)
	}
}

func DeleteProjectByUUIDs(c *gin.Context) {
	var reqUuids request.UuidsReq
	_ = c.ShouldBindJSON(&reqUuids)
	uuids := reqUuids.Uuids
	if len(uuids) == 0 {
		response.FailWithMessage("删除项目失败，请指定项目的uuid", c)
		return
	}
	err := service.DeleteProjectByUUIDs(uuids)
	if err != nil {
		global.LOG.Info("删除项目失败，错误", zap.Any("err: ", err))
		response.FailWithMessage("删除项目失败", c)
	} else {
		response.Ok(c)
	}
}

func getStringUUIDFromPath(c *gin.Context) string {
	uuid := c.Query("uuid")
	return uuid
}

func ExportProjectXls(c *gin.Context) {
	userId := getUserID(c)
	list, _ := service.ListProjectsToExportExcel(userId)
	headers := []string{"项目创建日期", "项目名称", "项目简介", "项目创建人"}

	filename := fmt.Sprintf("%s%s.xlsx", utils.GetDateTimeStr(), "-export_project")
	utils.ToExcel(c, filename, headers, list)
}
