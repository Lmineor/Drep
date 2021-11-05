package service

import (
	"github.com/drep/global"
	"github.com/drep/model"
)

func CreateInviteCode(invite model.SysInvite) (*model.SysInvite, error) {
	err := global.DB.Create(&invite).Error
	return &invite, err
}

func DeleteInviteCode(id uint) (err error) {
	var inv model.SysInvite
	err = global.DB.Where("id = ?", id).Unscoped().Delete(&inv).Error
	return err
}

func ListInviteCode(pageNum, pageSize int) (list interface{}, total int64, err error) {
	var inviteList []model.SysInvite
	limit := pageSize
	offset := (pageNum - 1) * limit

	db := global.DB.Model(&model.SysInvite{})
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&inviteList).Error
	return inviteList, total, err
}
