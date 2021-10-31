package service

import (
	"github.com/drep/global"
	"github.com/drep/model"
)

func RegisterProject(p *model.DpProject) (*model.DpProject, error) {
	err := global.DB.Create(p).Error
	return p, err
}

func ListAllProjects(pageNum, pageSize int) (list interface{}, total int64, err error) {
	var projectList []model.DpProject
	limit := pageSize
	offset := (pageNum - 1) * limit

	db := global.DB.Model(&model.DpProject{})
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&projectList).Error
	return projectList, total, err
}

func UpdateProject(p *model.DpProject) (*model.DpProject, error) {
	err := global.DB.Updates(&p).Error
	return p, err
}

func DeleteProject(id float64) error {
	var dp model.DpProject
	err := global.DB.Where("id = ?", id).Unscoped().Delete(&dp).Error
	return err
}
