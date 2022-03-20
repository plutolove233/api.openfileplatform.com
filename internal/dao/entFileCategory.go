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
	category := []EntFileCategory{}
	return category,mysqlManager.Model(&EntFileCategory{}).Where(map[string]interface{}{
		"EnterpriseID":id,
		"IsDeleted":0,
	}).Find(&category).Error
}

func (m *EntFileCategory)GetPath()(string,string,error){
	var name_path,id_path string
	var err error
	err = nil
	mysqlManager := database.GetMysqlClient()
	category := EntFileCategory{}
	nowId := m.CategoryID
	for{
		//如果遍历到根节点 则退出
		if nowId == "" {
			break
		}
		//如果无法找到分类信息则退出
		err = mysqlManager.Model(&EntFileCategory{}).Where(map[string]interface{}{
			"CategoryID":nowId,
			"IsDeleted":0,
		}).Take(&category).Error
		if err != nil {
			break
		}

		name_path = category.CategoryName + "/" + name_path
		id_path = category.CategoryID + "/" + id_path
		nowId = category.CategoryParentID
		category = EntFileCategory{}
	}
	return name_path, id_path, err
}

func (m *EntFileCategory)GetRootPath()(string,string,error){
	var name_path,id_path string
	var err error
	err = nil
	mysqlManager := database.GetMysqlClient()
	category := EntFileCategory{}
	err = mysqlManager.Model(&EntFileCategory{}).Where(map[string]interface{}{
		"EnterpriseID":m.EnterpriseID,
		"CategoryParentID":"",
	}).Take(&category).Error
	name_path = category.CategoryName + "/"
	id_path = category.CategoryID + "/"
	return name_path, id_path, err
}