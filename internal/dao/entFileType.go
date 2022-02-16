/*
@Coding : utf-8
@Time : 2022/2/15 16:19
@Author : 刘浩宇
@Software: GoLand
*/
package dao

import (
	"api.openfileplatform.com/internal/globals/database"
	"api.openfileplatform.com/internal/models/mysqlModel"
)

type EntFileType struct {
	mysqlModel.EntFileType
}

func (m *EntFileType)Get() error {
	mysqlMamager := database.GetMysqlClient()
	return mysqlMamager.Where(map[string]interface{}{
		"FileTypeID":m.FileTypeID,
	}).Where(m).Take(m).Error
}


func (m *EntFileType)Add() error{
	mysqlMamager := database.GetMysqlClient()
	return mysqlMamager.Create(&m).Error
}

func (m *EntFileType) Update(args map[string]interface{}) error {
	mysqlManager := database.GetMysqlClient()
	err := m.Get()
	if err != nil {
		return err
	}
	return mysqlManager.Model(&m).Updates(args).Error
}

func (m *EntFileType) Delete() error {
	mysqlManager := database.GetMysqlClient()
	err := m.Get()
	if err != nil {
		return err
	}
	return mysqlManager.Model(&m).Where(map[string]interface{}{
		"FileTypeID":m.FileTypeID,
	}).Delete(EntFileType{}).Error
}

func (*EntFileType)GetAll(id string) ([]EntFileType,error){
	mysqlManager := database.GetMysqlClient()
	file_types := []EntFileType{}
	return file_types,mysqlManager.Model(&EntFileType{}).Where(map[string]interface{}{
		"EnterpriseID":id,
	}).Find(&file_types).Error
}