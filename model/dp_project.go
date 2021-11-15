package model

import (
	"github.com/drep/global"
)

type DpProject struct {
	global.MODEL
	UUID        string  `json:"uuid" gorm:"comment:项目UUID"`
	Name        string  `json:"name" gorm:"unique;comment:项目名"`
	Description string  `json:"description" gorm:"comment:项目简介"`
	UserID      uint    `json:"user_id" form:"user_id" gorm:"column:user_id;comment:用户id"` // 管理员id
	User        SysUser `json:"user"`                                                      //创建该项目的管理员
}

func (DpProject) TableName() string {
	return "dp_projects"
}
