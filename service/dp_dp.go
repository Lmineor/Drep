package service

import (
	"github.com/drep/global"
	"github.com/drep/model"
)

func CreateDailyReport(mdp *model.DpDp) (*model.DpDp, error) {
	err := global.DB.Create(&mdp).Error
	return mdp, err
}
