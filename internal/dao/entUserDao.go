package dao

import (
	"api.openfileplatform.com/internal/globals/database"
	"api.openfileplatform.com/internal/models/mysqlModel"
)

type EntUsers struct {
	mysqlModel.EntUsers
}

func (m *EntUsers)Get() error {
	mysqlMamager := database.GetMysqlClient()
	return mysqlMamager.Where(map[string]interface{}{
		"IsDeleted":0,
	}).Where(m).Take(m).Error
}

func (m *EntUsers)GetUserID()string{
	return m.UserID
}
func (m *EntUsers)SetUserID(id string){
	m.UserID  = id
}
func (m *EntUsers)GetIsAdmin()bool {
	return m.IsAdmin
}
func (m *EntUsers)SetAccount(account string){
	m.Account = account
}
func (m *EntUsers) GetPassword() string {
	return m.Password
}
func (m *EntUsers) SetUserName(p string) {
	m.UserName = p
}
func (m *EntUsers) SetPassword(p string) {
	m.Password = p
}
func (m *EntUsers) SetPhone(p string) {
	m.Phone = p
}
func (m *EntUsers) SetEmail(p string) {
	m.Email = p
}

func (m *EntUsers)Add() error{
	mysqlMamager := database.GetMysqlClient()
	return mysqlMamager.Create(&m).Error
}

func (m *EntUsers) Update(args map[string]interface{}) error {
	mysqlManager := database.GetMysqlClient()
	return mysqlManager.Model(&m).Updates(args).Error
}

func (m *EntUsers) Delete(updateUser int64) error {
	mysqlManager := database.GetMysqlClient()

	return mysqlManager.Model(&m).Updates(map[string]interface{}{
		"IsDeleted": 1,
	}).Error
}

func (*EntUsers)GetAll() ([]EntUsers,error){
	mysqlManager := database.GetMysqlClient()
	users := []EntUsers{}
	return users,mysqlManager.Model(&PlatUsers{}).Find(&users).Error
}