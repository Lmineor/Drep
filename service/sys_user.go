package service

import (
	"errors"
	"github.com/drep/core/request"
	"github.com/drep/global"
	"github.com/drep/model"
	"github.com/drep/utils"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

func Register(u model.SysUser) (err error, userInter model.SysUser) {
	var user model.SysUser
	if !errors.Is(global.DB.Where("username = ?", u.Username).First(&user).Error, gorm.ErrRecordNotFound) { // 判断用户名是否注册
		return errors.New("用户名已注册"), userInter
	}
	// 否则 附加uuid 密码md5简单加密 注册
	u.Password = utils.MD5V([]byte(u.Password))
	u.UUID = utils.GenerateUUID()
	err = global.DB.Create(&u).Error
	return err, u
}

func Login(u *model.SysUser) (err error, userInter *model.SysUser) {
	var user model.SysUser
	u.Password = utils.MD5V([]byte(u.Password))
	err = global.DB.Where("username = ? AND password = ?", u.Username, u.Password).Preload("Authority").First(&user).Error
	return err, &user
}

func ChangePassword(u *model.SysUser, newPassword string) (err error, userInter *model.SysUser) {
	var user model.SysUser
	u.Password = utils.MD5V([]byte(u.Password))
	err = global.DB.Where("username = ? AND password = ?", u.Username, u.Password).First(&user).Update("password", utils.MD5V([]byte(newPassword))).Error
	return err, u
}

func GetUserInfoList(currentUserId uint, info request.PageInfo) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.DB.Model(&model.SysUser{}).Where("parent_id = ?", currentUserId)
	var userList []model.SysUser
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Preload("Authority").Find(&userList).Error
	return err, userList, total
}

func GetAllUserInfoList(info request.PageInfo) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.DB.Model(&model.SysUser{})
	var userList []model.SysUser
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Preload("Authority").Find(&userList).Error
	return err, userList, total
}

func SetUserAuthority(uuid uuid.UUID, authorityId string) (err error) {
	err = global.DB.Where("uuid = ?", uuid).First(&model.SysUser{}).Update("authority_id", authorityId).Error
	return err
}

func DeleteUser(uuid string) (err error) {
	var user model.SysUser
	err = global.DB.Where("uuid = ?", uuid).Delete(&user).Error
	return err
}

func SetUserInfo(reqUser *model.SysUser) (err error, user *model.SysUser) {
	err = global.DB.Updates(&reqUser).Error
	return err, reqUser
}

// FindUserById is find user by id
func FindUserById(id int) (err error, user *model.SysUser) {
	var u model.SysUser
	err = global.DB.Where("`id` = ?", id).First(&u).Error
	return err, &u
}

// FindUserByUuid is find user by uuid
func FindUserByUuid(uuid string) (err error, user *model.SysUser) {
	var u model.SysUser
	if err = global.DB.Where("`uuid` = ?", uuid).First(&u).Error; err != nil {
		return errors.New("用户不存在"), &u
	}
	return nil, &u
}

func VerifyInviteCode(code string, authorityId string) (err error) {
	var ia model.SysInvite
	if err = global.DB.Where("`invite_code` = ? AND `authority_id` = ? ", code, authorityId).Preload("Authority").First(&ia).Error; err != nil {
		return errors.New("邀请码或角色id不存在")
	}
	return nil
}

func VerifyAuthorityIdExist(authorityId string) bool {
	var au model.SysAuthority
	if err := global.DB.Where("authority_id = ?", authorityId).First(&au).Error; err != nil {
		return false
	}
	return true
}

func ResetPassword(userId float64) (err error) {
	var dbUser model.SysUser
	if err = global.DB.Where("id = ?", userId).First(&dbUser).Error; err != nil {
		return
	}
	dbUser.Password = utils.MD5V([]byte("123456"))
	err = global.DB.Save(&dbUser).Error
	return err
}
