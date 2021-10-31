package model

import (
	"github.com/drep/global"
)

type DpProject struct {
	global.MODEL
	Name        string `json:"name" gorm:"comment:项目名"`
	Description string `json:"description" gorm:"comment:"项目简介"`
}

func (DpProject) TableName() string {
	return "dp_projects"
}
