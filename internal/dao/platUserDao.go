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

func (m *PlatUsers)GetUserID()string{
	return m.UserID
}
func (m *PlatUsers)SetUserID(id string){
	m.UserID  = id
}
func (m *PlatUsers)GetIsAdmin()bool {
	return true
}
func (m *PlatUsers)SetAccount(account string){
	m.Account = account
}
func (m *PlatUsers) GetPassword() string {
	return m.Password
}
func (m *PlatUsers) SetUserName(p string) {
	m.UserName = p
}
func (m *PlatUsers) SetPassword(p string) {
	m.Password = p
}
func (m *PlatUsers) SetPhone(p string) {
	m.Phone = p
}
func (m *PlatUsers) SetEmail(p string) {
	m.Email = p
}

func (m *PlatUsers) Add() error {
	mysqlManager := database.GetMysqlClient()
	return mysqlManager.Create(&m).Error
}

func (m *PlatUsers) Update(args map[string]interface{}) error {
	mysqlManager := database.GetMysqlClient()
	err := m.Get()
	if err != nil {
		return err
	}
	return mysqlManager.Model(&m).Updates(args).Error
}

func (m *PlatUsers) Delete() error {
	mysqlManager := database.GetMysqlClient()
	err := m.Get()
	if err != nil {
		return err
	}
	return mysqlManager.Model(&m).Updates(map[string]interface{}{
		"IsDeleted": 1,
	}).Error
}

func (*PlatUsers)GetAll() ([]PlatUsers,error){
	mysqlManager := database.GetMysqlClient()
	users := []PlatUsers{}
	return users,mysqlManager.Model(&PlatUsers{}).Find(&users).Error
}