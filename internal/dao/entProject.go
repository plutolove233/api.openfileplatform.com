/*
@Coding : utf-8
@Time : 2022/2/15 16:21
@Author : 刘浩宇
@Software: GoLand
*/
package dao

import (
	"api.openfileplatform.com/internal/globals/database"
	"api.openfileplatform.com/internal/models/mysqlModel"
)

type EntProject struct {
	mysqlModel.EntProject
}

func (m *EntProject)Get() error {
	mysqlMamager := database.GetMysqlClient()
	return mysqlMamager.Where(map[string]interface{}{
		"IsDeleted":0,
	}).Where(m).Take(m).Error
}


func (m *EntProject)Add() error{
	mysqlMamager := database.GetMysqlClient()
	return mysqlMamager.Create(&m).Error
}

func (m *EntProject) Update(args map[string]interface{}) error {
	mysqlManager := database.GetMysqlClient()
	err := m.Get()
	if err != nil {
		return err
	}
	return mysqlManager.Model(&m).Updates(args).Error
}

func (m *EntProject) Delete() error {
	mysqlManager := database.GetMysqlClient()
	err := m.Get()
	if err != nil {
		return err
	}
	return mysqlManager.Model(&m).Updates(map[string]interface{}{
		"IsDeleted": 1,
	}).Error
}

func (*EntProject)GetAll(id string) ([]EntProject,error){
	mysqlManager := database.GetMysqlClient()
	department := []EntProject{}
	return department,mysqlManager.Model(&EntProject{}).Where(map[string]interface{}{
		"EnterpriseID":id,
		"IsDeleted":0,
	}).Find(&department).Error
}