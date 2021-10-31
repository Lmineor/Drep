package model

import (
	"github.com/drep/global"
)

type DpLoc struct {
	global.MODEL
	Province  string    `json:"province" gorm:"comment:省"`
	City      string    `json:"province" gorm:"comment:市"`
	ProjectID int       `json:"project_id" form:"project_id" gorm:"column:project_id;comment:项目id"`
	Project   DpProject `json:"project"`
}
