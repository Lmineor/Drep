package service

import (
	"github.com/drep/global"
	"github.com/drep/model"
)

func CreateDailyReport(mdp *model.DpDp) (*model.DpDp, error) {
	err := global.DB.Create(&mdp).Error
	return mdp, err
}

func ListAllDailyReports(pageNum, pageSize int) (list interface{}, total int64, err error) {
	var dpList []model.DpDp
	limit := pageSize
	offset := (pageNum - 1) * limit

	db := global.DB.Model(&model.DpDp{})
	err = db.Count(&total).Error
	err = db.Preload("User").Preload("Project").Limit(limit).Offset(offset).Find(&dpList).Error
	return dpList, total, err
}

func ListAllDpsWithSuperUser(adminId uint, pageNum, pageSize int) (list interface{}, total int64, err error) {
	var dpList []model.SysAdminDpMaps
	limit := pageSize
	offset := (pageNum - 1) * limit

	db := global.DB.Model(&model.SysAdminDpMaps{}).Where("user_id = ?", adminId)
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&dpList).Error
	return dpList, total, err
}

func ListDpsToExportExcel(userId uint) (list []map[string]interface{}, err error) {
	var dpList []model.SysAdminDpMaps
	db := global.DB.Model(&model.SysAdminDpMaps{}).Where("user_id = ?", userId)
	err = db.Find(&dpList).Error

	for _, v := range dpList {
		singleDp := make(map[string]interface{})
		singleDp["填写日期"] = v.CreatedAt
		singleDp["标题"] = v.Title
		singleDp["内容"] = v.Content
		singleDp["项目名"] = v.PjName
		singleDp["填写人"] = v.DpUsername + "(" + v.DpNickname + ")"
		list = append(list, singleDp)
	}
	return
}

func ListDps(userID uint, pageNum, pageSize int) (list interface{}, total int64, err error) {
	var dpList []model.DpDp
	limit := pageSize
	offset := (pageNum - 1) * limit

	db := global.DB.Model(&model.DpDp{})
	err = db.Count(&total).Error
	err = db.Preload("User").Preload("Project").Limit(limit).Offset(offset).Where("user_id = ?", userID).Find(&dpList).Error
	return dpList, total, err
}

func UpdateDailyReport(p *model.DpDp) (*model.DpDp, error) {
	var oldDp model.DpDp
	err := global.DB.Where("uuid = ?", p.UUID).First(&oldDp).Error
	if err != nil {
		return nil, err
	}
	oldDp.Title = p.Title
	oldDp.Content = p.Content
	global.DB.Save(&oldDp)
	return &oldDp, err
}

func DeleteDailyReport(uuid string) error {
	var dp model.DpDp
	err := global.DB.Where("uuid = ?", uuid).Unscoped().Delete(&dp).Error
	return err
}

func DeleteDailyReportByUUIds(uuids []string) error {
	err := global.DB.Delete(&[]model.DpDp{}, "uuid in ?", uuids).Error
	return err
}

func GetDailyReport(uuid string) (*model.DpDp, error) {
	var dp model.DpDp
	err := global.DB.Preload("User").Preload("Project").Where("uuid = ?", uuid).First(&dp).Error
	return &dp, err
}
