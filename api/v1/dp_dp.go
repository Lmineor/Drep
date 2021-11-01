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

func CreateDailyReport(c *gin.Context) {

	var dp request.Dp
	c.ShouldBindJSON(&dp)
	userId := getUserID(c)
	mDp := model.DpDp{Title: dp.Title, Content: dp.Content, ProjectID: dp.ProjectID, UserID: userId}
	dbDp, err := service.CreateDailyReport(&mDp)
	if err != nil {
		response.FailWithMessage("日报填写失败，请联系管理员", c)
	} else {
		response.OkWithDetailed(response.DpResponse{DailyReport: dbDp}, "日报填写成功", c)
	}
}

func ListAllDailyReports(c *gin.Context) {
	var total int64
	pageNum, pageSize := utils.ParsePaginateParams(c)

	list, total, err := service.ListAllDailyReports(pageNum, pageSize)
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

// ListDailyReports list daily report according to the user id.
func ListDps(c *gin.Context) {
	var total int64
	pageNum, pageSize := utils.ParsePaginateParams(c)
	userId := getUserID(c)
	list, total, err := service.ListDps(userId, pageNum, pageSize)
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

func UpdateDailyReport(c *gin.Context) {
	var pj model.DpDp
	c.ShouldBindJSON(&pj)
	updatedPj, err := service.UpdateDailyReport(&pj)
	if err != nil {
		errMsg := fmt.Sprintf("更新项目失败，错误：%s", err)
		response.FailWithMessage(errMsg, c)
	} else {
		response.OkWithDetailed(updatedPj, "更新成功", c)
	}

}

func DeleteDailyReport(c *gin.Context) {
	var reqId request.GetById
	_ = c.ShouldBindJSON(&reqId)
	if err := utils.Verify(reqId, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err := service.DeleteDailyReport(reqId.ID)
	if err != nil {
		errMsg := fmt.Sprintf("删除项目失败，错误：%s", err)
		global.LOG.Info(errMsg)
		response.FailWithMessage("删除项目失败", c)
	} else {
		response.Ok(c)
	}
}
