//平台超级管理员账号信息

package platform

import (
	"DocumentSystem/dao"
	"time"
)

type PlatUser struct {
	AutoID int64 `gorm:"AUTO_INCREMENT;column:AutoID;primary_key"`
	UserID int64 `gorm:"column:PlatUserID"`
	UserName string `gorm:"column:UserName"`//用户真实姓名
	Account string `gorm:"column:Account"`//用户账号
	Pwd string `gorm:"column:Pwd"`
	Phone string `gorm:"column:Phone"`
	Email string `gorm:"column:Email"`
	IsDeleted bool `gorm:"column:IsDeleted"`
	CreateTime time.Time `gorm:"column:CreateTime"`
}

type PlatUserModels interface {
	Add()(error,bool)
	Login()bool
	Delete()error
	Reverse()error
	ChangePwd(p string)error
	ChangePhone(p string)error
	ChangeEmail(e string)error
}

func (u *PlatUser)Add()(error,bool){
	var user PlatUser
	user.AutoID = 0
	dao.DB.Model(&PlatUser{}).Where("Account = ?",u.Account).Find(&user)
	if user.AutoID==0{
		return dao.DB.Create(u).Error,true
	}
	return nil,false
}

func (u *PlatUser)Delete()error{
	return dao.DB.Model(&PlatUser{}).Where("UserID = ?",u.UserID).Update("IsDeleted",true).Error
}

func (u *PlatUser)Reverse()error{
	return dao.DB.Model(&PlatUser{}).Where("UserID = ?",u.UserID).Update("IsDeleted",false).Error
}

func (u *PlatUser)ChangePwd(p string)error{
	return dao.DB.Model(&PlatUser{}).Where("UserID = ?",u.UserID).Update("Pwd",p).Error
}

func (u *PlatUser)ChangePhone(p string)error{
	return dao.DB.Model(&PlatUser{}).Where("UserID = ?",u.UserID).Update("Phone",p).Error
}

func (u *PlatUser)ChangeEmail(e string)error{
	return dao.DB.Model(&PlatUser{}).Where("UserID = ?",u.UserID).Update("Email",e).Error
}

func (u *PlatUser)Login()bool{
	var user PlatUser
	dao.DB.Where("Account = ?",u.Account).Find(&user)
	ok := user.Pwd==u.Pwd
	return ok
}