package v1

import (
	"github.com/drep/core/response"
	"github.com/drep/global"
	"github.com/drep/model"
	"github.com/drep/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func JsonInBlacklist(c *gin.Context) {
	token := c.Request.Header.Get("x-token")
	jwt := model.JwtBlacklist{Jwt: token}
	if err := service.JsonInBlacklist(jwt); err != nil {
		global.LOG.Error("jwt作废失败!", zap.Any("err", err))
		response.FailWithMessage("jwt作废失败", c)
	} else {
		response.OkWithMessage("jwt作废成功", c)
	}
}
