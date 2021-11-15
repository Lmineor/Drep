package source

import (
	"github.com/drep/global"
	"gorm.io/gorm"
)

var AuthoritiesMenus = new(authoritiesMenus)

type authoritiesMenus struct{}

type AuthorityMenus struct {
	AuthorityId string `gorm:"column:sys_authority_authority_id"`
	BaseMenuId  uint   `gorm:"column:sys_base_menu_id"`
}

var authorityMenus = []AuthorityMenus{
	{"100", 1},
	{"100", 2},
	{"100", 3},
	{"100", 4},
	{"100", 5},
	{"100", 6},
	{"100", 7},
	{"100", 8},
	{"100", 9},
	{"100", 10},
	{"100", 11},
	{"100", 12},
	{"100", 13},
	{"100", 14},
	{"100", 15},
	{"100", 16},
	{"100", 17},
	{"100", 18},
	{"1001", 1},
	{"1001", 2},
	{"1001", 8},
	{"1001", 9},
	{"1001", 10},
	{"1001", 11},
}

func (a *authoritiesMenus) Init() error {
	return global.DB.Table("sys_authority_menus").Transaction(func(tx *gorm.DB) error {
		if tx.Where("sys_authority_authority_id IN ('888', '100', '1001')").Find(&[]AuthorityMenus{}).RowsAffected == 48 {
			global.LOG.Warn("\n[Mysql] --> sys_authority_menus 表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&authorityMenus).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		global.LOG.Info("\n[Mysql] --> sys_authority_menus 表初始数据成功!")
		return nil
	})
}
