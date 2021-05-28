//平台超级管理员账号信息

package platform

import (
	"DocumentSystem/dao"
	"time"
)

type PlatUser struct {
	PlatUserID int `gorm:"AUTO_INCREMENT" gorm:"column:PlatUserID"`
	PlatUserName string `gorm:"column:PlatUserName"`
	Pwd string `gorm:"column:Pwd"`
	PlatUserPhone string `gorm:"column:PlatUserPhone"`
	PlatUserEmail string `gorm:"column:PlatUserEmail"`
	Deleted bool `gorm:"column:Deleted"`
	CreateTime time.Time `gorm:"column:CreateTime"`
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
	return dao.DB.Model(&PlatUser{}).Where("PlatUserID = ?",u.PlatUserID).Update("Deleted",true).Error
}

func (u *PlatUser)Reverse()error{
	return dao.DB.Model(&PlatUser{}).Where("PlatUserID = ?",u.PlatUserID).Update("Deleted",false).Error
}

func (u *PlatUser)ChangePwd(p string)error{
	return dao.DB.Model(&PlatUser{}).Where("PlatUserID = ?",u.PlatUserID).Update("Pwd",p).Error
}

func (u *PlatUser)ChangePhone(p string)error{
	return dao.DB.Model(&PlatUser{}).Where("PlatUserID = ?",u.PlatUserID).Update("PlatUserPhone",p).Error
}

func (u *PlatUser)ChangeEmail(e string)error{
	return dao.DB.Model(&PlatUser{}).Where("PlatUserID = ?",u.PlatUserID).Update("PlatUserEmail",e).Error
}

func (u *PlatUser)Login()bool{
	var user PlatUser
	dao.DB.Where("PlatUserName = ?",u.PlatUserName).Find(&user)
	ok := user.Pwd==u.Pwd
	return ok
}