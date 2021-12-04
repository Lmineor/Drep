package source

import (
	"github.com/drep/global"
	"github.com/drep/model"
)

var AdminDpMap = new(adminDpMap)

type adminDpMap struct{}

func (a *adminDpMap) Init() error {
	if global.DB.Find(&[]model.SysAdminDpMaps{}).RowsAffected > 0 {
		global.LOG.Warn("\n[Mysql] --> sys_admin_dp_maps 视图已存在!")
		return nil
	}
	if err := global.DB.Exec(
		"CREATE ALGORITHM = UNDEFINED SQL SECURITY DEFINER VIEW `sys_admin_dp_maps` AS SELECT `dp_dps`.`uuid` AS `dp_uuid`, `dp_dps`.`created_at` AS `created_at`, `dp_dps`.`updated_at` AS `updated_at`, `dp_dps`.`deleted_at` AS `deleted_at`, `dp_dps`.`title` AS `title`, `dp_dps`.`content` AS `content`, `sys_users`.`username` AS `dp_username`, `sys_users`.`nick_name` AS `dp_nickname`, `dp_projects`.`user_id` AS `user_id`, `dp_projects`.`name` AS `pj_name` from `dp_projects` join sys_users on `dp_projects`.`user_id` = `sys_users`.`id` join `dp_dps` on `dp_projects`.`id` = `dp_dps`.`project_id`;").Error; err != nil {
		return err
	}
	global.LOG.Info("\n[Mysql] --> sys_admin_dp_maps 视图创建成功!")
	return nil
}
