package model

import (
	"github.com/drep/global"
)

type DpDp struct {
	global.MODEL
	UUID      string    `json:"uuid" gorm:"comment:日报UUID"`
	Title     string    `json:"title" gorm:"type:varchar(255);comment:日报标题"`
	Content   string    `json:"content" gorm:"type:varchar(255);comment:日报内容"`
	UserID    uint      `json:"user_id" form:"user_id" gorm:"column:user_id;comment:用户id"`
	User      SysUser   `json:"user"`
	ProjectID uint      `json:"project_id" form:"project_id" gorm:"column:project_id;comment:项目id"`
	Project   DpProject `json:"project"`
}

func (DpDp) TableName() string {
	return "dp_dps"
}
