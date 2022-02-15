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

type EntFileProject struct {
	mysqlModel.EntProject
}

func (m *EntFileProject)Get() error {
	mysqlMamager := database.GetMysqlClient()
	return mysqlMamager.Where(map[string]interface{}{
		"IsDeleted":0,
	}).Where(m).Take(m).Error
}


func (m *EntFileProject)Add() error{
	mysqlMamager := database.GetMysqlClient()
	return mysqlMamager.Create(&m).Error
}

func (m *EntFileProject) Update(args map[string]interface{}) error {
	mysqlManager := database.GetMysqlClient()
	err := m.Get()
	if err != nil {
		return err
	}
	return mysqlManager.Model(&m).Updates(args).Error
}

func (m *EntFileProject) Delete() error {
	mysqlManager := database.GetMysqlClient()
	err := m.Get()
	if err != nil {
		return err
	}
	return mysqlManager.Model(&m).Updates(map[string]interface{}{
		"IsDeleted": 1,
	}).Error
}

func (*EntFileProject)GetAll(id string) ([]EntFileProject,error){
	mysqlManager := database.GetMysqlClient()
	department := []EntFileProject{}
	return department,mysqlManager.Model(&EntFileProject{}).Where("EnterpriseID = ? AND isDelete = ?",id,false).
		Find(&department).Error
}