// coding: utf-8
// @Author : lryself
// @Date : 2021/12/28 20:31
// @Software: GoLand

package dao

import (
	"api.openfileplatform.com/internal/globals/database"
	"api.openfileplatform.com/internal/models/mysqlModel"
)

type PlatUsers struct {
	mysqlModel.PlatUsers
}

func (m *PlatUsers) Get() error {
	mysqlManager := database.GetMysqlClient()
	return mysqlManager.Where(map[string]interface{}{
		"IsDeleted": 0,
	}).Where(m).Take(m).Error
}

func (m *PlatUsers) Add() error {
	mysqlManager := database.GetMysqlClient()
	return mysqlManager.Create(&m).Error
}

func (m *PlatUsers) Update(args map[string]interface{}) error {
	mysqlManager := database.GetMysqlClient()
	return mysqlManager.Model(&m).Updates(args).Error
}

func (m *PlatUsers) Delete(updateUser int64) error {
	mysqlManager := database.GetMysqlClient()

	return mysqlManager.Model(&m).Updates(map[string]interface{}{
		"IsDeleted": 1,
	}).Error
}

func (*PlatUsers)GetAll() (error,[]PlatUsers){
	mysqlManager := database.GetMysqlClient()
	users := []PlatUsers{}
	return mysqlManager.Model(&PlatUsers{}).Find(&users).Error,users
}