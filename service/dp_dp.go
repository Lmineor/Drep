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
	err = db.Limit(limit).Offset(offset).Find(&dpList).Error
	return dpList, total, err
}

func ListDps(userId uint, pageNum, pageSize int) (list interface{}, total int64, err error) {
	var dpList []model.DpDp
	limit := pageSize
	offset := (pageNum - 1) * limit

	db := global.DB.Model(&model.DpDp{})
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Where("user_id = ?", userId).Find(&dpList).Error
	return dpList, total, err
}

func UpdateDailyReport(p *model.DpDp) (*model.DpDp, error) {
	err := global.DB.Updates(&p).Error
	return p, err
}

func DeleteDailyReport(id float64) error {
	var dp model.DpDp
	err := global.DB.Where("id = ?", id).Unscoped().Delete(&dp).Error
	return err
}
