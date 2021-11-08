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
	"strconv"
)

func GetInviteCode(c *gin.Context) {
	var reqId request.GetById
	_ = c.ShouldBindQuery(&reqId)
	if err := utils.Verify(reqId, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	invCode, err := service.GetInviteCodeById(reqId.ID)
	if err != nil {
		global.LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithData(response.InviteCodeResponse{InviteCode: invCode}, c)
	}
}

func UpdateInviteCodeById(c *gin.Context) {
	var req request.InviteCode
	_ = c.ShouldBindJSON(&req)
	if err := utils.Verify(req, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	updateInviteCode := model.SysInvite{
		InviteCode:  req.InviteCode,
		AuthorityId: req.AuthorityId,
		Description: req.Description,
	}
	invCode, err := service.UpdateInviteCodeById(req.ID, &updateInviteCode)
	if err != nil {
		global.LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithData(response.InviteCodeResponse{InviteCode: invCode}, c)
	}
}

func CreateInviteCode(c *gin.Context) {
	var reqInvite request.InviteCode
	_ = c.ShouldBindJSON(&reqInvite)
	if err := utils.Verify(reqInvite, utils.InviteCodeVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	mInviteCode := model.SysInvite{
		InviteCode:  reqInvite.InviteCode,
		AuthorityId: reqInvite.AuthorityId,
		Description: reqInvite.Description,
	}
	dbInviteCode, err := service.CreateInviteCode(&mInviteCode)
	if err != nil {
		errMsg := fmt.Sprintf("邀请码创建失败: %s", err)
		response.FailWithMessage(errMsg, c)
		return
	}
	response.OkWithDetailed(response.InviteCodeResponse{InviteCode: dbInviteCode}, "邀请码创建成功", c)
}

func DeleteInviteCode(c *gin.Context) {
	var reqId request.GetById
	_ = c.ShouldBindJSON(&reqId)
	err := service.DeleteInviteCode(reqId.ID)
	if err != nil {
		errMsg := fmt.Sprintf("删除项目失败，错误：%s", err)
		global.LOG.Info(errMsg)
		response.FailWithMessage("删除项目失败", c)
	} else {
		response.Ok(c)
	}
}
func ListInviteCode(c *gin.Context) {
	var total int64
	pageNum, pageSize := utils.ParsePaginateParams(c)
	list, total, err := service.ListInviteCode(pageNum, pageSize)
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

func StrToUInt(str string) uint {
	i, e := strconv.Atoi(str)
	if e != nil {
		return 0
	}
	return uint(i)
}
