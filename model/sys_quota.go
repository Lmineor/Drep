package model

import (
	"github.com/drep/global"
)

type SysQuota struct {
	global.MODEL
	Resource    string       `json:"resource" gorm:"column:resource;comment:资源名称"`
	Quota       int          `json:"quota" gorm:"column:quota;comment:配额;default:1"`
	Authority   SysAuthority `json:"authority" gorm:"foreignKey:AuthorityId;comment:用户角色"`
	AuthorityId string       `json:"authorityId" gorm:"unique;default:888;comment:用户角色ID"` // 用户角色ID
}

func (SysQuota) TableName() string {
	return "sys_quotas"
}
