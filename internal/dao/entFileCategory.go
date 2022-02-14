/*
@Coding : utf-8
@Time : 2022/2/14 15:52
@Author : 刘浩宇
@Software: GoLand
*/
package dao

import (
	"api.openfileplatform.com/internal/globals/database"
	"api.openfileplatform.com/internal/models/mysqlModel"
)

type EntFileCategory struct {
	mysqlModel.EntFileCategory
}

func (m *EntFileCategory)Get() error {
	mysqlMamager := database.GetMysqlClient()
	return mysqlMamager.Where(map[string]interface{}{
		"IsDeleted":0,
	}).Where(m).Take(m).Error
}


func (m *EntFileCategory)Add() error{
	mysqlMamager := database.GetMysqlClient()
	return mysqlMamager.Create(&m).Error
}

func (m *EntFileCategory) Update(args map[string]interface{}) error {
	mysqlManager := database.GetMysqlClient()
	err := m.Get()
	if err != nil {
		return err
	}
	return mysqlManager.Model(&m).Updates(args).Error
}

func (m *EntFileCategory) Delete() error {
	mysqlManager := database.GetMysqlClient()
	err := m.Get()
	if err != nil {
		return err
	}
	return mysqlManager.Model(&m).Updates(map[string]interface{}{
		"IsDeleted": 1,
	}).Error
}

func (*EntFileCategory)GetAll(id string) ([]EntFileCategory,error){
	mysqlManager := database.GetMysqlClient()
	users := []EntFileCategory{}
	return users,mysqlManager.Model(&EntFileCategory{}).Where("EnterpriseID = ?",id).Find(&users).Error
}