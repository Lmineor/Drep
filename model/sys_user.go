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
	HeaderImg   string       `json:"headerImg" gorm:"default:http://101.200.193.47/images/avatar.png;comment:用户头像"` // 用户头像
	Authority   SysAuthority `json:"authority" gorm:"foreignKey:AuthorityId;comment:用户角色"`
	AuthorityId string       `json:"authorityId" gorm:"default:888;comment:用户角色ID"`     // 用户角色ID
	SideMode    string       `json:"sideMode" gorm:"default:dark;comment:用户角色ID"`       // 用户侧边主题
	ActiveColor string       `json:"activeColor" gorm:"default:#1890ff;comment:用户角色ID"` // 活跃颜色
	BaseColor   string       `json:"baseColor" gorm:"default:#fff;comment:用户角色ID"`      // 基础颜色
	ParentId    uint         `json:"parentId" gorm:"comment:父用户ID"`
}

func (SysUser) TableName() string {
	return "sys_users"
}
