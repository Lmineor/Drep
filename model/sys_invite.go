package model

import "github.com/drep/global"

type SysInvite struct {
	global.MODEL
	InviteCode  string       `json:"invite_code" gorm:"default:888;comment:邀请码"`
	Authority   SysAuthority `json:"-" gorm:"foreignKey:AuthorityId;comment:用户角色"`
	AuthorityId string       `json:"authority_id" gorm:"unique;default:888;comment:用户角色ID"` // 用户角色ID
	Description string       `json:"description" gorm:"comment:简介"`
}

func (SysInvite) TableName() string {
	return "sys_invites"
}
