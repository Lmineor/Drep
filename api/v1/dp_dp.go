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

func CreateDailyReport(c *gin.Context) {

	var dp request.Dp
	_ = c.ShouldBindJSON(&dp)

	if err := utils.Verify(dp, utils.DpVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	userId := getUserID(c)
	dbProject, err := service.GetProjectByUuid(dp.ProjectUUID)
	if err != nil {
		response.FailWithMessage("项目uuid不存在", c)
	}
	uuid := utils.GenerateUUID()
	mDp := model.DpDp{Title: dp.Title, Content: dp.Content, ProjectID: dbProject.ID, UserID: userId, UUID: uuid}
	dbDp, err := service.CreateDailyReport(&mDp)
	if err != nil {
		response.FailWithMessage("日报填写失败，请联系管理员", c)
	} else {
		response.OkWithDetailed(response.DpResponse{DailyReport: dbDp}, "日报填写成功", c)
	}
}

func GetDpDetail(c *gin.Context) {
	var reqUuid request.GetByUUID
	_ = c.ShouldBindQuery(&reqUuid)
	uuid := reqUuid.UUID
	if uuid == "" {
		response.FailWithMessage("未指定uuid", c)
		return
	}
	dbDp, err := service.GetDailyReport(uuid)
	if err != nil {
		response.FailWithMessage("该日报不存在", c)
	} else {
		response.OkWithDetailed(response.DpResponse{DailyReport: dbDp}, "成功", c)
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

func ListAllDpsWithSuperUser(c *gin.Context) {
	var total int64
	adminId := getUserID(c)
	pageNum, pageSize := utils.ParsePaginateParams(c)

	list, total, err := service.ListAllDpsWithSuperUser(adminId, pageNum, pageSize)
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

// ListDps list daily report according to the user id.
func ListDps(c *gin.Context) {
	var total int64
	pageNum, pageSize := utils.ParsePaginateParams(c)
	userID := getUserID(c)
	list, total, err := service.ListDps(userID, pageNum, pageSize)
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
	var dp model.DpDp
	_ = c.ShouldBindJSON(&dp)
	if err := utils.Verify(dp, utils.DpVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	updatedPj, err := service.UpdateDailyReport(&dp)
	if err != nil {
		errMsg := fmt.Sprintf("更新项目失败，错误：%s", err)
		response.FailWithMessage(errMsg, c)
	} else {
		response.OkWithDetailed(updatedPj, "更新成功", c)
	}

}

func DeleteDailyReport(c *gin.Context) {
	var reqUuid request.GetByUUID
	_ = c.ShouldBindQuery(&reqUuid)
	uuid := reqUuid.UUID
	if uuid == "" {
		response.FailWithMessage("请指定uuid", c)
		return
	}
	err := service.DeleteDailyReport(uuid)
	if err != nil {
		errMsg := fmt.Sprintf("删除项目失败，错误：%s", err)
		global.LOG.Info(errMsg)
		response.FailWithMessage("删除项目失败", c)
	} else {
		response.Ok(c)
	}
}

func DeleteDailyReportByUUIds(c *gin.Context) {
	var reqUuid request.UuidsReq
	_ = c.ShouldBindJSON(&reqUuid)
	if len(reqUuid.Uuids) == 0 {
		response.FailWithMessage("请指定uuid", c)
		return
	}
	err := service.DeleteDailyReportByUUIds(reqUuid.Uuids)
	if err != nil {
		global.LOG.Info("批量删除日报失败", zap.Any("error:", err))
		response.FailWithMessage("批量删除日报失败", c)
	} else {
		response.Ok(c)
	}
}

func ExportDpXls(c *gin.Context) {
	userId := getUserID(c)
	list, _ := service.ListDpsToExportExcel(userId)
	headers := []string{"填写日期", "标题", "内容", "项目名", "填写人"}

	filename := fmt.Sprintf("%s%s.xlsx", utils.GetDateTimeStr(), "-export_daily_report")
	utils.ToExcel(c, filename, headers, list)
}
