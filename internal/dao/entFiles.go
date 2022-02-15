/*
@Coding : utf-8
@Time : 2022/2/14 14:59
@Author : 刘浩宇
@Software: GoLand
*/
package dao

import (
	"api.openfileplatform.com/internal/globals/database"
	"api.openfileplatform.com/internal/models/mysqlModel"
)

type EntFiles struct {
	mysqlModel.EntFile
}

func (m *EntFiles)Get() error {
	mysqlMamager := database.GetMysqlClient()
	return mysqlMamager.Where(map[string]interface{}{
		"IsDeleted":0,
	}).Where(m).Take(m).Error
}


func (m *EntFiles)Add() error{
	mysqlMamager := database.GetMysqlClient()
	return mysqlMamager.Create(&m).Error
}

func (m *EntFiles) Update(args map[string]interface{}) error {
	mysqlManager := database.GetMysqlClient()
	err := m.Get()
	if err != nil {
		return err
	}
	return mysqlManager.Model(&m).Updates(args).Error
}

func (m *EntFiles) Delete() error {
	mysqlManager := database.GetMysqlClient()
	err := m.Get()
	if err != nil {
		return err
	}
	return mysqlManager.Model(&m).Updates(map[string]interface{}{
		"IsDeleted": 1,
	}).Error
}

func (*EntFiles)GetAll(id string) ([]EntFiles,error){
	mysqlManager := database.GetMysqlClient()
	files := []EntFiles{}
	return files,mysqlManager.Model(&EntFiles{}).Where("EnterpriseID = ? AND isDelete = ?",id,false).
		Find(&files).Error
}