/*
@Coding : utf-8
@Time : 2022/2/12 9:36
@Author : 刘浩宇
@Software: GoLand
*/
package dao

import (
	"api.openfileplatform.com/internal/globals/database"
	"api.openfileplatform.com/internal/models/mysqlModel"
)

type EntDepartment struct {
	mysqlModel.EntDepartments
}

func (m *EntDepartment)Get() error {
	mysqlMamager := database.GetMysqlClient()
	return mysqlMamager.Where(map[string]interface{}{
		"IsDeleted":0,
	}).Where(m).Take(m).Error
}


func (m *EntDepartment)Add() error{
	mysqlMamager := database.GetMysqlClient()
	return mysqlMamager.Create(&m).Error
}

func (m *EntDepartment) Update(args map[string]interface{}) error {
	mysqlManager := database.GetMysqlClient()
	return mysqlManager.Model(&m).Updates(args).Error
}

func (m *EntDepartment) Delete() error {
	mysqlManager := database.GetMysqlClient()

	return mysqlManager.Model(&m).Updates(map[string]interface{}{
		"IsDeleted": 1,
	}).Error
}

func (*EntDepartment)GetAll(id string) ([]EntDepartment,error){
	mysqlManager := database.GetMysqlClient()
	users := []EntDepartment{}
	return users,mysqlManager.Model(&PlatUsers{}).Where("EnterpriseID = ?",id).Find(&users).Error
}