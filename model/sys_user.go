package model

import (
	"github.com/drep/global"
)

type SysUser struct {
	global.MODEL
	UUID        string       `json:"uuid" gorm:"comment:用户UUID"`                                                    // 用户UUID
	Username    string       `json:"userName" gorm:"comment:用户登录名"`                                                 // 用户登录名
	Password    string       `json:"-" gorm:"comment:用户登录密码"`                                                       // 用户登录密码
	NickName    string       `json:"nickName" gorm:"default:系统用户;comment:用户昵称"`                                     // 用户昵称
	HeaderImg   string       `json:"headerImg" gorm:"default:http://qmplusimg.henrongyi.top/head.png;comment:用户头像"` // 用户头像
	Authority   SysAuthority `json:"authority" gorm:"foreignKey:AuthorityId;comment:用户角色"`
	AuthorityId string       `json:"authorityId" gorm:"unique;default:888;comment:用户角色ID"` // 用户角色ID
}

func (SysUser) TableName() string {
	return "sys_users"
}
