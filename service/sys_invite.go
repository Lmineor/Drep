package service

import (
	"fmt"
	"github.com/drep/global"
	"github.com/drep/model"
)

func CreateInviteCode(invite *model.SysInvite) (*model.SysInvite, error) {
	err := global.DB.Create(&invite).Error
	return invite, err
}

func UpdateInviteCodeById(id float64, invite *model.SysInvite) (*model.SysInvite, error) {
	var oldInviteCode model.SysInvite
	err := global.DB.Where("id = ?", id).First(&oldInviteCode).Error
	if err != nil {
		return nil, err
	}
	if invite.Description != "" {
		err = global.DB.Model(&oldInviteCode).Update("description", invite.Description).Error
	}
	return &oldInviteCode, err
}
func GetInviteCodeById(id float64) (*model.SysInvite, error) {
	var inv model.SysInvite
	err := global.DB.Where("id = ?", id).First(&inv).Error
	return &inv, err
}

func DeleteInviteCode(id float64) (err error) {
	var inv model.SysInvite
	str := fmt.Sprintf("%f", id)
	global.LOG.Info(str)
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
