//平台超级管理员账号信息

package platform

import (
	"DocumentSystem/dao"
	"time"
)

type PlatUser struct {
	AutoID int64 `gorm:"AUTO_INCREMENT;column:AutoID;primary_key"`
	UserID int64 `gorm:"column:PlatUserID"`
	UserName string `gorm:"column:UserName" form:"username"`//用户真实姓名
	Account string `gorm:"column:Account" form:"account"`//用户账号
	Pwd string `gorm:"column:Pwd" form:"pwd"`
	Phone string `gorm:"column:Phone" form:"phone"`
	Email string `gorm:"column:Email" form:"email"`
	IsDeleted bool `gorm:"column:IsDeleted"`
	CreateTime time.Time `gorm:"column:CreateTime"`
}

type PlatUserModels interface {
	Add()error
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