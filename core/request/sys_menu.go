package request

import (
	"github.com/drep/global"
	"github.com/drep/model"
)

// Add menu authority info structure
type AddMenuAuthorityInfo struct {
	Menus       []model.SysBaseMenu
	AuthorityId string // 角色ID
}

func DefaultMenu() []model.SysBaseMenu {
	return []model.SysBaseMenu{{
		MODEL:     global.MODEL{ID: 1},
		ParentId:  "0",
		Path:      "dashboard",
		Name:      "dashboard",
		Component: "view/dashboard/index.vue",
		Sort:      1,
		Meta: model.Meta{
			Title: "仪表盘",
			Icon:  "setting",
		},
	}}
}
