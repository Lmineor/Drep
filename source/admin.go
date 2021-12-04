package source

import (
	"github.com/drep/global"
	"github.com/drep/model"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"time"
)

var Admin = new(admin)

type admin struct{}

var admins = []model.SysUser{
	{MODEL: global.MODEL{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()}, UUID: uuid.NewV4().String(), Username: "admin", Password: "e10adc3949ba59abbe56e057f20f883e", NickName: "超级管理员", HeaderImg: "https://www.mineor.xyz/images/avatar.png", AuthorityId: "100"},
	{MODEL: global.MODEL{ID: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()}, UUID: uuid.NewV4().String(), Username: "superUser", Password: "e10adc3949ba59abbe56e057f20f883e", NickName: "普通管理员", HeaderImg: "https://www.mineor.xyz/images/avatar.png", AuthorityId: "1001"},
}

func (a *admin) Init() error {
	return global.DB.Transaction(func(tx *gorm.DB) error {
		if tx.Where("id IN ?", []int{1, 2}).Find(&[]model.SysUser{}).RowsAffected == 2 {
			global.LOG.Error("\n[Mysql] --> sys_users 表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&admins).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		global.LOG.Info("\n[Mysql] --> sys_users 表初始数据成功!")
		return nil
	})
}
