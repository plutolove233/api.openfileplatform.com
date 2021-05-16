//平台超级管理员账号信息

package platform

import (
	"DocumentSystem/dao"
	"time"
)

type PlatUser struct {
	Id int `gorm:"AUTO_INCREMENT"`
	Name string
	Pwd string
	Phone string
	Email string
	Deleted bool
	CreateTime time.Time
}

type PlatUserModels interface {
	Add()error
	Login()bool
	Delete()error
	Reverse()error
	ChangePwd(p string)error
	ChangePhone(p string)error
	ChangeEmail(e string)error
}

func (u *PlatUser)Add()error{
	return dao.DB.Create(u).Error
}

func (u *PlatUser)Delete()error{
	return dao.DB.Model(&PlatUser{}).Where("id = ?",u.Id).Update("deleted",true).Error
}

func (u *PlatUser)Reverse()error{
	return dao.DB.Model(&PlatUser{}).Where("id = ?",u.Id).Update("deleted",false).Error
}

func (u *PlatUser)ChangePwd(p string)error{
	return dao.DB.Model(&PlatUser{}).Where("id = ?",u.Id).Update("pwd",p).Error
}

func (u *PlatUser)ChangePhone(p string)error{
	return dao.DB.Model(&PlatUser{}).Where("id = ?",u.Id).Update("phone",p).Error
}

func (u *PlatUser)ChangeEmail(e string)error{
	return dao.DB.Model(&PlatUser{}).Where("id = ?",u.Id).Update("email",e).Error
}

func (u *PlatUser)Login()bool{
	var user PlatUser
	dao.DB.Where("name = ?",u.Name).Find(&user)
	ok := user.Pwd==u.Pwd
	return ok
}