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
	{"100", 19},
	{"100", 20},
	{"100", 21},
	{"100", 22},
	{"100", 23},
	{"100", 24},
	{"100", 25},
	{"100", 26},
	{"1001", 1},
	{"1001", 2},
	{"1001", 3},
	{"1001", 4},
	{"1001", 5},
	{"1001", 6},
	{"1001", 7},
	{"1001", 8},
	{"1001", 9},
	{"1001", 10},
	{"1001", 11},
	{"1001", 12},
	{"1001", 13},
	{"1001", 14},
	{"1001", 15},
	{"1001", 16},
	{"1001", 17},
	{"1001", 18},
	{"1001", 19},
	{"1001", 20},
	{"1001", 21},
	{"1001", 22},
	{"1001", 23},
	{"1001", 25},
	{"1001", 26},
	{"888", 1},
	{"888", 2},
	{"888", 3},
	{"888", 4},
	{"888", 5},
	{"888", 6},
	{"888", 7},
	{"888", 8},
	{"888", 9},
	{"888", 10},
	{"888", 11},
	{"888", 12},
	{"888", 13},
	{"888", 14},
	{"888", 15},
	{"888", 16},
	{"888", 17},
	{"888", 18},
	{"888", 19},
	{"888", 20},
	{"888", 21},
	{"888", 22},
	{"888", 23},
	{"888", 24},
	{"888", 25},
	{"8881", 1},
	{"8881", 2},
	{"8881", 8},
	{"9528", 1},
	{"9528", 2},
	{"9528", 3},
	{"9528", 4},
	{"9528", 5},
	{"9528", 6},
	{"9528", 7},
	{"9528", 8},
	{"9528", 9},
	{"9528", 10},
	{"9528", 11},
	{"9528", 12},
	{"9528", 14},
	{"9528", 15},
	{"9528", 16},
	{"9528", 17},
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: sys_authority_menus 表数据初始化
func (a *authoritiesMenus) Init() error {
	return global.DB.Table("sys_authority_menus").Transaction(func(tx *gorm.DB) error {
		if tx.Where("sys_authority_authority_id IN ('888', '8881', '9528')").Find(&[]AuthorityMenus{}).RowsAffected == 48 {
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
