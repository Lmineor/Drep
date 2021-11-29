package source

import (
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/drep/global"
	"gorm.io/gorm"
)

var Casbin = new(casbin)

type casbin struct{}

// 100: superAdmin
// 1001: admin
// 888: user

var carbines = []gormadapter.CasbinRule{
	// superAdmin
	{Ptype: "p", V0: "100", V1: "/system/getSystemConfig", V2: "GET"},
	{Ptype: "p", V0: "100", V1: "/system/setSystemConfig", V2: "PUT"},
	{Ptype: "p", V0: "100", V1: "/system/getServerInfo", V2: "GET"},
	{Ptype: "p", V0: "100", V1: "/system/reloadSystem", V2: "POST"},

	{Ptype: "p", V0: "100", V1: "/fileUploadAndDownload/getFileList", V2: "GET"},
	{Ptype: "p", V0: "100", V1: "/fileUploadAndDownload/deleteFile", V2: "POST"},
	{Ptype: "p", V0: "100", V1: "/fileUploadAndDownload/upload", V2: "POST"},

	{Ptype: "p", V0: "100", V1: "/admin/createInviteCode", V2: "POST"},
	{Ptype: "p", V0: "100", V1: "/admin/listAllInviteCode", V2: "GET"},
	{Ptype: "p", V0: "100", V1: "/admin/deleteInviteCode", V2: "DELETE"},
	{Ptype: "p", V0: "100", V1: "/admin/getInviteCodeById", V2: "GET"},
	{Ptype: "p", V0: "100", V1: "/admin/updateInviteCodeById", V2: "PUT"},

	{Ptype: "p", V0: "100", V1: "/api/createApi", V2: "POST"},
	{Ptype: "p", V0: "100", V1: "/api/deleteApi", V2: "POST"},
	{Ptype: "p", V0: "100", V1: "/api/getApiById", V2: "POST"},
	{Ptype: "p", V0: "100", V1: "/api/updateApi", V2: "POST"},
	{Ptype: "p", V0: "100", V1: "/api/getAllApis", V2: "POST"},
	{Ptype: "p", V0: "100", V1: "/api/deleteApisByIds", V2: "DELETE"},

	{Ptype: "p", V0: "100", V1: "/menu/getMenu", V2: "GET"},
	{Ptype: "p", V0: "100", V1: "/menu/getMenuList", V2: "GET"},
	{Ptype: "p", V0: "100", V1: "/menu/addBaseMenu", V2: "POST"},
	{Ptype: "p", V0: "100", V1: "/menu/getBaseMenuTree", V2: "GET"},
	{Ptype: "p", V0: "100", V1: "/menu/addMenuAuthority", V2: "POST"},
	{Ptype: "p", V0: "100", V1: "/menu/getMenuAuthority", V2: "GET"},
	{Ptype: "p", V0: "100", V1: "/menu/deleteBaseMenu", V2: "POST"},
	{Ptype: "p", V0: "100", V1: "/menu/updateBaseMenu", V2: "POST"},
	{Ptype: "p", V0: "100", V1: "/menu/getBaseMenuById", V2: "GET"},

	{Ptype: "p", V0: "100", V1: "/casbin/updateCasbin", V2: "POST"},
	{Ptype: "p", V0: "100", V1: "/casbin/getPolicyPathByAuthorityId", V2: "POST"},

	{Ptype: "p", V0: "100", V1: "/authority/createAuthority", V2: "POST"},
	{Ptype: "p", V0: "100", V1: "/authority/deleteAuthority", V2: "POST"},
	{Ptype: "p", V0: "100", V1: "/authority/updateAuthority", V2: "PUT"},
	{Ptype: "p", V0: "100", V1: "/authority/copyAuthority", V2: "POST"},
	{Ptype: "p", V0: "100", V1: "/authority/getAuthorityList", V2: "GET"},
	{Ptype: "p", V0: "100", V1: "/authority/setDataAuthority", V2: "POST"},

	{Ptype: "p", V0: "100", V1: "/sysOperationRecord/createSysOperationRecord", V2: "POST"},
	{Ptype: "p", V0: "100", V1: "/sysOperationRecord/deleteSysOperationRecord", V2: "DELETE"},
	{Ptype: "p", V0: "100", V1: "/sysOperationRecord/deleteSysOperationRecordByIds", V2: "DELETE"},
	{Ptype: "p", V0: "100", V1: "/sysOperationRecord/findSysOperationRecord", V2: "GET"},
	{Ptype: "p", V0: "100", V1: "/sysOperationRecord/getSysOperationRecordList", V2: "GET"},

	{Ptype: "p", V0: "100", V1: "/base/login", V2: "POST"},
	{Ptype: "p", V0: "100", V1: "/user/register", V2: "POST"},
	{Ptype: "p", V0: "100", V1: "/user/getUserList", V2: "GET"},
	{Ptype: "p", V0: "100", V1: "/user/getAllUserList", V2: "GET"},
	{Ptype: "p", V0: "100", V1: "/user/deleteUser", V2: "GET"},
	{Ptype: "p", V0: "100", V1: "/user/setUserInfo", V2: "PUT"},
	{Ptype: "p", V0: "100", V1: "/user/changePassword", V2: "POST"},
	{Ptype: "p", V0: "100", V1: "/user/setUserAuthority", V2: "POST"},
	{Ptype: "p", V0: "100", V1: "/user/resetUserPassword", V2: "POST"},

	{Ptype: "p", V0: "100", V1: "/jwt/jsonInBlacklist", V2: "POST"},

	{Ptype: "p", V0: "100", V1: "/project/listAllProjects", V2: "GET"},
	{Ptype: "p", V0: "100", V1: "/project/createProject", V2: "POST"},
	{Ptype: "p", V0: "100", V1: "/project/getProjectDetail", V2: "GET"},
	{Ptype: "p", V0: "100", V1: "/project/listProjects", V2: "GET"},
	{Ptype: "p", V0: "100", V1: "/project/updateProject", V2: "PUT"},
	{Ptype: "p", V0: "100", V1: "/project/deleteProject", V2: "DELETE"},
	{Ptype: "p", V0: "100", V1: "/project/deleteProjectByUUIDs", V2: "DELETE"},
	{Ptype: "p", V0: "100", V1: "/project/exportProjectXls", V2: "GET"},

	{Ptype: "p", V0: "100", V1: "/dp/listAllDps", V2: "GET"},
	{Ptype: "p", V0: "100", V1: "/dp/createDp", V2: "POST"},
	{Ptype: "p", V0: "100", V1: "/dp/listDps", V2: "GET"},
	{Ptype: "p", V0: "100", V1: "/dp/getDpDetail", V2: "GET"},
	{Ptype: "p", V0: "100", V1: "/dp/updateDpDetail", V2: "PUT"},
	{Ptype: "p", V0: "100", V1: "/dp/deleteDp", V2: "DELETE"},
	{Ptype: "p", V0: "100", V1: "/dp/deleteDpByUUIDs", V2: "DELETE"},
	{Ptype: "p", V0: "100", V1: "/dp/exportDpXls", V2: "GET"},
	// admin
	{Ptype: "p", V0: "1001", V1: "/base/login", V2: "POST"},

	{Ptype: "p", V0: "1001", V1: "/menu/getMenu", V2: "GET"},
	{Ptype: "p", V0: "1001", V1: "/menu/getMenuList", V2: "GET"},
	{Ptype: "p", V0: "1001", V1: "/menu/addBaseMenu", V2: "POST"},
	{Ptype: "p", V0: "1001", V1: "/menu/getBaseMenuTree", V2: "GET"},
	{Ptype: "p", V0: "1001", V1: "/menu/addMenuAuthority", V2: "POST"},
	{Ptype: "p", V0: "1001", V1: "/menu/getMenuAuthority", V2: "GET"},
	{Ptype: "p", V0: "1001", V1: "/menu/deleteBaseMenu", V2: "POST"},
	{Ptype: "p", V0: "1001", V1: "/menu/updateBaseMenu", V2: "POST"},
	{Ptype: "p", V0: "1001", V1: "/menu/getBaseMenuById", V2: "GET"},

	{Ptype: "p", V0: "1001", V1: "/user/register", V2: "POST"},
	{Ptype: "p", V0: "1001", V1: "/user/getUserList", V2: "GET"},
	{Ptype: "p", V0: "1001", V1: "/user/deleteUser", V2: "GET"},
	{Ptype: "p", V0: "1001", V1: "/user/setUserInfo", V2: "PUT"},
	{Ptype: "p", V0: "1001", V1: "/user/changePassword", V2: "POST"},
	{Ptype: "p", V0: "1001", V1: "/user/resetUserPassword", V2: "POST"},

	{Ptype: "p", V0: "1001", V1: "/jwt/jsonInBlacklist", V2: "POST"},

	{Ptype: "p", V0: "1001", V1: "/project/createProject", V2: "POST"},
	{Ptype: "p", V0: "1001", V1: "/project/getProjectDetail", V2: "GET"},
	{Ptype: "p", V0: "1001", V1: "/project/listProjects", V2: "GET"},
	{Ptype: "p", V0: "1001", V1: "/project/updateProject", V2: "PUT"},
	{Ptype: "p", V0: "1001", V1: "/project/deleteProject", V2: "DELETE"},
	{Ptype: "p", V0: "1001", V1: "/project/deleteProjectByUUIDs", V2: "DELETE"},
	{Ptype: "p", V0: "1001", V1: "/project/exportProjectXls", V2: "GET"},

	{Ptype: "p", V0: "1001", V1: "/dp/listAllDpsWithSuperUser", V2: "GET"},
	{Ptype: "p", V0: "1001", V1: "/dp/createDp", V2: "POST"},
	{Ptype: "p", V0: "1001", V1: "/dp/listDps", V2: "GET"},
	{Ptype: "p", V0: "1001", V1: "/dp/getDpDetail", V2: "GET"},
	{Ptype: "p", V0: "1001", V1: "/dp/updateDpDetail", V2: "PUT"},
	{Ptype: "p", V0: "1001", V1: "/dp/deleteDp", V2: "DELETE"},
	{Ptype: "p", V0: "1001", V1: "/dp/deleteDpByUUIDs", V2: "DELETE"},
	{Ptype: "p", V0: "1001", V1: "/dp/exportDpXls", V2: "GET"},

	// user
	{Ptype: "p", V0: "888", V1: "/base/login", V2: "POST"},
	{Ptype: "p", V0: "888", V1: "/user/setUserInfo", V2: "PUT"},
	{Ptype: "p", V0: "888", V1: "/user/changePassword", V2: "POST"},

	{Ptype: "p", V0: "888", V1: "/jwt/jsonInBlacklist", V2: "POST"},

	{Ptype: "p", V0: "888", V1: "/dp/createDp", V2: "POST"},
	{Ptype: "p", V0: "888", V1: "/dp/listDps", V2: "GET"},
	{Ptype: "p", V0: "888", V1: "/dp/getDpDetail", V2: "GET"},
	{Ptype: "p", V0: "888", V1: "/dp/updateDpDetail", V2: "PUT"},
	{Ptype: "p", V0: "888", V1: "/dp/deleteDp", V2: "DELETE"},
}

func (c *casbin) Init() error {
	_ = global.DB.AutoMigrate(gormadapter.CasbinRule{})
	return global.DB.Transaction(func(tx *gorm.DB) error {
		if tx.Find(&[]gormadapter.CasbinRule{}).RowsAffected == 67 {
			global.LOG.Error("[Mysql] --> casbin_rule 表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&carbines).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		global.LOG.Info("\n[Mysql] --> casbin_rule 表初始数据成功!")
		return nil
	})
}
