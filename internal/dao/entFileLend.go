/*
@Coding : utf-8
@Time : 2022/2/15 15:46
@Author : 刘浩宇
@Software: GoLand
*/
package dao

import (
	"api.openfileplatform.com/internal/globals/database"
	"api.openfileplatform.com/internal/models/mysqlModel"
)

type EntFileLend struct {
	mysqlModel.EntFileLend
}

func (m *EntFileLend)Get() error {
	mysqlMamager := database.GetMysqlClient()
	return mysqlMamager.Where(map[string]interface{}{
		"IsDeleted":0,
	}).Where(m).Take(m).Error
}


func (m *EntFileLend)Add() error{
	mysqlMamager := database.GetMysqlClient()
	return mysqlMamager.Create(&m).Error
}

func (m *EntFileLend) Update(args map[string]interface{}) error {
	mysqlManager := database.GetMysqlClient()
	err := m.Get()
	if err != nil {
		return err
	}
	return mysqlManager.Model(&m).Updates(args).Error
}

func (m *EntFileLend) Delete() error {
	mysqlManager := database.GetMysqlClient()
	err := m.Get()
	if err != nil {
		return err
	}
	return mysqlManager.Model(&m).Updates(map[string]interface{}{
		"IsDeleted": 1,
	}).Error
}

func (*EntFileLend)GetAll(id string) ([]EntFileLend,error){
	mysqlManager := database.GetMysqlClient()
	file_lend := []EntFileLend{}
	return file_lend,mysqlManager.Model(&EntFileLend{}).Where("EnterpriseID = ? AND isDelete = ?",id,false).
		Find(&file_lend).Error
}