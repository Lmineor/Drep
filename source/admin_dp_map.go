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
		"CREATE ALGORITHM = UNDEFINED SQL SECURITY DEFINER VIEW `sys_admin_dp_maps` AS select `dp_user_table`.`created_at` AS `created_at`, `dp_user_table`.`updated_at` AS `updated_at`, `dp_user_table`.`deleted_at` AS `deleted_at`, `dp_user_table`.`title` AS `title`, `dp_user_table`.`content` AS `content`, `dp_user_table`.`dp_uuid` AS `dp_uuid`, `dp_user_table`.`username` AS `dp_username`, `dp_user_table`.`nick_name` AS `dp_nickname`, `project_user_table`.`user_id` AS `user_id`, `project_user_table`.`name` AS `pj_name` from (select `dp_projects`.`user_id`, `dp_projects`.`name`, `dp_projects`.`id` AS `pj_id` from `dp_projects` join `sys_users` on (`dp_projects`.`user_id` = `sys_users`.`id`)) `project_user_table` join (select `dp_dps`.`project_id`, `dp_dps`.`created_at`, `dp_dps`.`updated_at`, `dp_dps`.`deleted_at`, `dp_dps`.`title`, `dp_dps`.`content`, `sys_users`.`username`, `sys_users`.`nick_name`, `dp_dps`.`uuid` AS `dp_uuid` from `dp_dps` join `sys_users` on (`dp_dps`.`user_id` = `sys_users`.`id`)) `dp_user_table` on (`project_user_table`.`pj_id` = `dp_user_table`.`project_id`)").Error; err != nil {
		return err
	}
	global.LOG.Info("\n[Mysql] --> sys_admin_dp_maps 视图创建成功!")
	return nil
}
