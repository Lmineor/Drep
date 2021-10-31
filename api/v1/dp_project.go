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
)

func RegistryProject(c *gin.Context) {
	var pj request.Project
	c.ShouldBindJSON(&pj)
	mPj := model.DpProject{Name: pj.Name, Description: pj.Description}
	dbPj, err := service.RegisterProject(&mPj)
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

func UpdateProject(c *gin.Context) {
	var pj model.DpProject
	c.ShouldBindJSON(&pj)
	updatedPj, err := service.UpdateProject(&pj)
	if err != nil {
		errMsg := fmt.Sprintf("更新项目失败，错误：%s", err)
		response.FailWithMessage(errMsg, c)
	} else {
		response.OkWithDetailed(updatedPj, "更新成功", c)
	}
}

func DeleteProject(c *gin.Context) {
	var reqId request.GetById
	_ = c.ShouldBindJSON(&reqId)
	if err := utils.Verify(reqId, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err := service.DeleteProject(reqId.ID)
	if err != nil {
		errMsg := fmt.Sprintf("删除项目失败，错误：%s", err)
		global.LOG.Info(errMsg)
		response.FailWithMessage("删除项目失败", c)
	} else {
		response.Ok(c)
	}
}
