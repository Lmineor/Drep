package model

import (
	"github.com/drep/global"
)

type SysAdminDpMaps struct {
	global.MODEL
	Userid     string `json:"user_id" gorm:"column:user_id;comment:用户ID"`
	DpUsername string `json:"dp_username" gorm:"comment:日报创建人"`
	DpNickname string `json:"dp_nickname" gorm:"dp_nickname;comment:用户昵称"`
	PjName     string `json:"pj_name" gorm:"comment:项目名"`
	DpUuid     string `json:"dp_uuid" gorm:"comment:日报UUID"`
	Title      string `json:"title" gorm:"type:varchar(255);comment:日报标题"`
	Content    string `json:"content" gorm:"type:varchar(255);comment:日报内容"`
}

func (SysAdminDpMaps) TableName() string {
	return "sys_admin_dp_maps"
}
