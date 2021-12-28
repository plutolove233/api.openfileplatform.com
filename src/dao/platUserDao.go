// coding: utf-8
// @Author : lryself
// @Date : 2021/12/28 20:31
// @Software: GoLand

package dao

import (
	"api.openfileplatform.com/src/globals/database"
	"api.openfileplatform.com/src/models/mysqlModel"
)

type PlatUsers struct {
	mysqlModel.PlatUsers
}

func (m *PlatUsers) Get() error {
	m.IsDeleted = false
	mysqlManager := database.GetMysqlClient()
	return mysqlManager.Where(m).Take(m).Error
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
