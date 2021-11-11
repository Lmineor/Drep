package model

import (
	"github.com/drep/global"
)

type DpProject struct {
	global.MODEL
	UUID        string `json:"uuid" gorm:"comment:项目UUID"`
	Name        string `json:"name" gorm:"unique;comment:项目名"`
	Description string `json:"description" gorm:"comment:项目简介"`
}

func (DpProject) TableName() string {
	return "dp_projects"
}
