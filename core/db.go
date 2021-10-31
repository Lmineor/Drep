package core

import (
	"github.com/drep/core/initial"
	"github.com/drep/global"
	"gorm.io/gorm"
)

func DB()*gorm.DB{
	switch global.CONFIG.System.DbType {
	case "mysql":
		return initial.GormMysql()
	default:
		return initial.GormMysql()
	}
}

