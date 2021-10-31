package service

import (
	"errors"
	"github.com/drep/global"
	"github.com/drep/model"
)

func SaveDailyReport(uuid string) (err error, user *model.SysUser) {
	var u model.SysUser
	if err = global.DB.Where("`uuid` = ?", uuid).First(&u).Error; err != nil {
		return errors.New("用户不存在"), &u
	}
	return nil, &u
}
