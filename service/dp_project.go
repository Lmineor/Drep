package service

import (
	"github.com/drep/global"
	"github.com/drep/model"
	"gorm.io/gorm"
)

func CreateProject(p *model.DpProject) (*model.DpProject, error) {
	err := global.DB.Create(p).Error
	return p, err
}

func GetProjectDetail(uuid string) (*model.DpProject, error) {
	return GetProjectByUuid(uuid)
}

func GetProjectByUuid(uuid string) (*model.DpProject, error) {
	var dp model.DpProject
	err := global.DB.Preload("User").Where("uuid = ?", uuid).Find(&dp).Error
	return &dp, err
}

func GetProjectById(id uint) (*model.DpProject, error) {
	var dp model.DpProject
	err := global.DB.Where("id = ?", id).Find(&dp).Error
	return &dp, err
}

func ListProjects(userId uint, pageNum, pageSize int) (list interface{}, total int64, err error) {
	var projectList []model.DpProject
	limit := pageSize
	offset := (pageNum - 1) * limit

	db := global.DB.Model(&model.DpProject{}).Where("user_id = ?", userId)
	err = db.Count(&total).Error
	// TODO: solve this omit func.
	err = db.Limit(limit).Offset(offset).Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Omit("sys_users.header_img")
	}).Find(&projectList).Error
	return projectList, total, err
}

func ListAllProjects(pageNum, pageSize int) (list interface{}, total int64, err error) {
	var projectList []model.DpProject
	limit := pageSize
	offset := (pageNum - 1) * limit

	db := global.DB.Model(&model.DpProject{})
	err = db.Count(&total).Error
	// TODO: solve this omit func.
	err = db.Limit(limit).Offset(offset).Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Omit("sys_users.header_img")
	}).Find(&projectList).Error
	return projectList, total, err
}

func UpdateProject(p *model.DpProject) (*model.DpProject, error) {
	var oldPj model.DpProject
	err := global.DB.Where("uuid = ?", p.UUID).First(&oldPj).Error
	if err != nil {
		return nil, err
	}
	oldPj.Name = p.Name
	oldPj.Description = p.Description
	global.DB.Save(&oldPj)
	return &oldPj, err
}

func DeleteProject(uuid string) error {
	var dp model.DpProject
	err := global.DB.Where("uuid = ?", uuid).Unscoped().Delete(&dp).Error
	return err
}

func DeleteProjectByUUIDs(uuids []string) error {
	err := global.DB.Delete(&[]model.DpProject{}, "uuid in ?", uuids).Error
	return err
}
