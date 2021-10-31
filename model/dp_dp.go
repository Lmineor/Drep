package model

import (
	"github.com/drep/global"
)

type DpDp struct {
	global.MODEL
	Content   string    `json:"context" gorm:"type:varchar(255);comment:日报内容"`
	UserID    int       `json:"user_id" form:"user_id" gorm:"column:user_id;comment:用户id"`
	User      SysUser   `json:"user"`
	ProjectID int       `json:"project_id" form:"project_id" gorm:"column:project_id;comment:项目id"`
	Project   DpProject `json:"project"`
}
