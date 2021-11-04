package source

import (
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/drep/global"
	"gorm.io/gorm"
)

var Casbin = new(casbin)

type casbin struct{}

// 999: superAdmin
// 888: admin
// 777: user

var carbines = []gormadapter.CasbinRule{
	// superAdmin
	{Ptype: "p", V0: "999", V1: "/base/login", V2: "POST"},
	{Ptype: "p", V0: "999", V1: "/base/register", V2: "POST"},

	{Ptype: "p", V0: "999", V1: "/project/listAllProjects", V2: "GET"},
	{Ptype: "p", V0: "999", V1: "/project/createProject", V2: "POST"},
	{Ptype: "p", V0: "999", V1: "/project/getProjectDetail/:pathParam", V2: "GET"},
	{Ptype: "p", V0: "999", V1: "/project/listProjects", V2: "GET"},
	{Ptype: "p", V0: "999", V1: "/project/updateProject/:pathParam", V2: "PUT"},
	{Ptype: "p", V0: "999", V1: "/project/deleteProject/:pathParam", V2: "DELETE"},

	{Ptype: "p", V0: "999", V1: "/dp/listAllDps", V2: "GET"},
	{Ptype: "p", V0: "999", V1: "/dp/createDp", V2: "POST"},
	{Ptype: "p", V0: "999", V1: "/dp/listDps", V2: "GET"},
	{Ptype: "p", V0: "999", V1: "/dp/getDpDetail/:pathParam", V2: "GET"},
	{Ptype: "p", V0: "999", V1: "/dp/updateDpDetail/:pathParam", V2: "PUT"},
	{Ptype: "p", V0: "999", V1: "/dp/deleteDp/:pathParam", V2: "DELETE"},

	// admin
	{Ptype: "p", V0: "888", V1: "/base/login", V2: "POST"},
	{Ptype: "p", V0: "888", V1: "/base/register", V2: "POST"},

	{Ptype: "p", V0: "888", V1: "/project/createProject", V2: "POST"},
	{Ptype: "p", V0: "888", V1: "/project/getProjectDetail/:pathParam", V2: "GET"},
	{Ptype: "p", V0: "888", V1: "/project/listProjects", V2: "GET"},
	{Ptype: "p", V0: "888", V1: "/project/updateProject/:pathParam", V2: "PUT"},
	{Ptype: "p", V0: "888", V1: "/project/deleteProject/:pathParam", V2: "DELETE"},

	{Ptype: "p", V0: "888", V1: "/dp/listAllDps", V2: "GET"},
	{Ptype: "p", V0: "888", V1: "/dp/createDp", V2: "POST"},
	{Ptype: "p", V0: "888", V1: "/dp/listDps", V2: "GET"},
	{Ptype: "p", V0: "888", V1: "/dp/getDpDetail/:pathParam", V2: "GET"},
	{Ptype: "p", V0: "888", V1: "/dp/updateDpDetail/:pathParam", V2: "PUT"},
	{Ptype: "p", V0: "888", V1: "/dp/deleteDp/:pathParam", V2: "DELETE"},

	// user
	{Ptype: "p", V0: "777", V1: "/base/login", V2: "POST"},
	{Ptype: "p", V0: "777", V1: "/base/register", V2: "POST"},

	{Ptype: "p", V0: "777", V1: "/dp/createDp", V2: "POST"},
	{Ptype: "p", V0: "777", V1: "/dp/listDps", V2: "GET"},
	{Ptype: "p", V0: "777", V1: "/dp/getDpDetail/:pathParam", V2: "GET"},
	{Ptype: "p", V0: "777", V1: "/dp/updateDpDetail/:pathParam", V2: "PUT"},
	{Ptype: "p", V0: "777", V1: "/dp/deleteDp/:pathParam", V2: "DELETE"},
}

func (c *casbin) Init() error {
	global.DB.AutoMigrate(gormadapter.CasbinRule{})
	return global.DB.Transaction(func(tx *gorm.DB) error {
		if tx.Find(&[]gormadapter.CasbinRule{}).RowsAffected == 154 {
			global.LOG.Error("[Mysql] --> casbin_rule 表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&carbines).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		global.LOG.Info("\n[Mysql] --> casbin_rule 表初始数据成功!")
		return nil
	})
}
