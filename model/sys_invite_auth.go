package model

import "github.com/drep/global"

type SysInvite struct {
	global.MODEL
	InviteCode  string       `json:"-" gorm:"default:888;comment:邀请码"`
	Authority   SysAuthority `json:"authority" gorm:"foreignKey:AuthorityId;references:AuthorityId;comment:用户角色"`
	AuthorityId string       `json:"authorityId" gorm:"default:888;comment:用户角色ID"` // 用户角色ID
}
