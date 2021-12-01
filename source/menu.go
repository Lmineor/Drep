package source

import (
	"github.com/drep/global"
	"github.com/drep/model"
	"time"

	"gorm.io/gorm"
)

var BaseMenu = new(menu)

type menu struct{}

var menus = []model.SysBaseMenu{
	{MODEL: global.MODEL{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, ParentId: "0", Path: "dashboard", Name: "dashboard", Hidden: false, Component: "view/dashboard/index.vue", Sort: 1, Meta: model.Meta{Title: "仪表盘", Icon: "setting"}},
	{MODEL: global.MODEL{ID: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "0", Path: "about", Name: "about", Component: "view/about/index.vue", Sort: 7, Meta: model.Meta{Title: "关于我们", Icon: "info"}},
	{MODEL: global.MODEL{ID: 3, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "0", Path: "admin", Name: "superAdmin", Component: "view/superAdmin/index.vue", Sort: 2, Meta: model.Meta{Title: "超级管理员", Icon: "user-solid"}},
	{MODEL: global.MODEL{ID: 4, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "3", Path: "authority", Name: "authority", Component: "view/superAdmin/authority/authority.vue", Sort: 0, Meta: model.Meta{Title: "角色管理", Icon: "s-custom"}},
	{MODEL: global.MODEL{ID: 5, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: true, ParentId: "3", Path: "menu", Name: "menu", Component: "view/superAdmin/menu/menu.vue", Sort: 1, Meta: model.Meta{Title: "菜单管理", Icon: "s-order", KeepAlive: true}},
	{MODEL: global.MODEL{ID: 6, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "3", Path: "user", Name: "user", Component: "view/superAdmin/user/user.vue", Sort: 2, Meta: model.Meta{Title: "用户管理", Icon: "coordinate"}},
	{MODEL: global.MODEL{ID: 7, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: true, ParentId: "0", Path: "person", Name: "person", Component: "view/person/person.vue", Sort: 3, Meta: model.Meta{Title: "个人信息", Icon: "message-solid"}},
	{MODEL: global.MODEL{ID: 8, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "0", Path: "superuser", Name: "superUser", Component: "view/superUser/index.vue", Sort: 4, Meta: model.Meta{Title: "管理员", Icon: "user-solid"}},
	{MODEL: global.MODEL{ID: 9, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "8", Path: "worker", Name: "worker", Component: "view/superUser/worker/worker.vue", Sort: 0, Meta: model.Meta{Title: "工人管理", Icon: "coordinate"}},
	{MODEL: global.MODEL{ID: 10, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "8", Path: "workerproject", Name: "workerproject", Component: "view/superUser/project/project.vue", Sort: 1, Meta: model.Meta{Title: "项目管理", Icon: "s-open"}},
	{MODEL: global.MODEL{ID: 11, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "8", Path: "workerdailyreport", Name: "workerdailyreport", Component: "view/superUser/dailyreport/dailyreport.vue", Sort: 2, Meta: model.Meta{Title: "日报管理", Icon: "notebook-1"}},
	{MODEL: global.MODEL{ID: 12, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "0", Path: "systemTools", Name: "systemTools", Component: "view/systemTools/index.vue", Sort: 5, Meta: model.Meta{Title: "系统工具", Icon: "s-cooperation"}},
	{MODEL: global.MODEL{ID: 13, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "12", Path: "system", Name: "system", Component: "view/systemTools/system/system.vue", Sort: 1, Meta: model.Meta{Title: "系统配置", Icon: "s-operation"}},
	{MODEL: global.MODEL{ID: 14, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "3", Path: "operation", Name: "operation", Component: "view/superAdmin/operation/sysOperationRecord.vue", Sort: 3, Meta: model.Meta{Title: "操作历史", Icon: "time"}},
	{MODEL: global.MODEL{ID: 15, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, ParentId: "0", Path: "http://www.mineor.xyz", Name: "https://www.mineor.xyz", Hidden: false, Component: "/", Sort: 0, Meta: model.Meta{Title: "官方网站", Icon: "s-home"}},
	{MODEL: global.MODEL{ID: 16, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, ParentId: "0", Path: "state", Name: "state", Hidden: false, Component: "view/system/state.vue", Sort: 8, Meta: model.Meta{Title: "服务器状态", Icon: "cloudy"}},
	{MODEL: global.MODEL{ID: 17, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "3", Path: "project", Name: "project", Component: "view/superAdmin/project/project.vue", Sort: 5, Meta: model.Meta{Title: "项目管理", Icon: "s-open"}},
	{MODEL: global.MODEL{ID: 18, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "3", Path: "dailyreport", Name: "dailyreport", Component: "view/superAdmin/dailyreport/dailyreport.vue", Sort: 6, Meta: model.Meta{Title: "日报管理", Icon: "notebook-1"}},
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: sys_base_menus 表数据初始化
func (m *menu) Init() error {
	return global.DB.Transaction(func(tx *gorm.DB) error {
		if tx.Where("id IN ?", []int{1, 18}).Find(&[]model.SysBaseMenu{}).RowsAffected == 2 {
			global.LOG.Info("\n[Mysql] --> sys_base_menus 表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&menus).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		global.LOG.Info("\n[Mysql] --> sys_base_menus 表初始数据成功!")
		return nil
	})
}
