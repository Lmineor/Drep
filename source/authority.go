package source

import (
	"github.com/drep/global"
	"github.com/drep/model"
	"time"

	"gorm.io/gorm"
)

var Authority = new(authority)

type authority struct{}

var authorities = []model.SysAuthority{
	{CreatedAt: time.Now(), UpdatedAt: time.Now(), AuthorityId: "888", AuthorityName: "普通用户", ParentId: "0", DefaultRouter: "dashboard"},
	{CreatedAt: time.Now(), UpdatedAt: time.Now(), AuthorityId: "100", AuthorityName: "超级管理员", ParentId: "0", DefaultRouter: "dashboard"},
	{CreatedAt: time.Now(), UpdatedAt: time.Now(), AuthorityId: "1001", AuthorityName: "普通管理员", ParentId: "0", DefaultRouter: "dashboard"},
}

func (a *authority) Init() error {
	return global.DB.Transaction(func(tx *gorm.DB) error {
		if tx.Where("authority_id IN ? ", []string{"888", "1001", "100"}).Find(&[]model.SysAuthority{}).RowsAffected == 3 {
			global.LOG.Warn("\n[Mysql] --> sys_authorities 表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&authorities).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		global.LOG.Info("\n[Mysql] --> sys_authorities 表初始数据成功!")
		return nil
	})
}
