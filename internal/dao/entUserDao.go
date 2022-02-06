package dao

import (
	"api.openfileplatform.com/internal/globals/database"
	"api.openfileplatform.com/internal/models/mysqlModel"
)

type EntUsers struct {
	mysqlModel.EntUsers
}

func (m *EntUsers)Get() error {
	mysqlMamager := database.GetMysqlClient()
	return mysqlMamager.Where(map[string]interface{}{
		"IsDeleted":0,
	}).Where(m).Take(m).Error
}

func (m *EntUsers)Add() error{
	mysqlMamager := database.GetMysqlClient()
	return mysqlMamager.Create(&m).Error
}

func (m *EntUsers) Update(args map[string]interface{}) error {
	mysqlManager := database.GetMysqlClient()
	return mysqlManager.Model(&m).Updates(args).Error
}

func (m *EntUsers) Delete(updateUser int64) error {
	mysqlManager := database.GetMysqlClient()

	return mysqlManager.Model(&m).Updates(map[string]interface{}{
		"IsDeleted": 1,
	}).Error
}

func (*EntUsers)GetAll() (error,[]EntUsers){
	mysqlManager := database.GetMysqlClient()
	users := []EntUsers{}
	return mysqlManager.Model(&PlatUsers{}).Find(&users).Error,users
}