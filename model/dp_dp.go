package model

import (
	"github.com/drep/global"
)

type DpDp struct {
	global.MODEL
	Title     string    `json:"title" gorm:"type:varchar(255);comment:日报标题"`
	Content   string    `json:"context" gorm:"type:varchar(255);comment:日报内容"`
	UserID    uint      `json:"user_id" form:"user_id" gorm:"column:user_id;comment:用户id"`
	User      SysUser   `json:"-"`
	ProjectID uint      `json:"project_id" form:"project_id" gorm:"column:project_id;comment:项目id"`
	Project   DpProject `json:"-"`
}

func (DpDp) TableName() string {
	return "dp_dps"
}
