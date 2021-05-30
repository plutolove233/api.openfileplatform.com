//企业用户表

package ent

import (
	"DocumentSystem/dao"
	"time"
)

type EntUser struct{
	AutoID int64 `gorm:"AUTO_INCREMENT;column:AutoID;primary_key"`
	UserID int `gorm:"AUTO_INCREMENT;column:UserID"`
	EnterpriseID int64 `gorm:"column:EnterpriseID"`
	Account string `gorm:"column:Account"`//账号
	Pwd string `form:"Pwd" gorm:"column:Pwd"`
	UserName string `form:"UserName" gorm:"column:UserName"`//用户真实姓名
	UserRoleID string `gorm:"column:UserRoleID"`
	Phone string `gorm:"column:Phone"`
	Email string `gorm:"column:Email"`
	FacePicUrl string `gorm:"column:FacePicUrl"`//user avatar like the address
	IsAdmin int `gorm:"column:IsAdmin"`
	IsDeleted bool `gorm:"column:IsDeleted"`
	CreateTime time.Time `gorm:"column:CreateTime"`
}

type EntUserModels interface {
	Register()(error,bool)
	Login()bool
	ChangeName(u string)error
	ChangePwd(u string)error
	ChangeFace(u string) error
}

func (u *EntUser)Register() (error, bool){
	var user []EntUser
	ok := true
	dao.DB.Find(&user)
	for _,index := range user{
		if index==*u{
			ok = false
			return nil,ok
		}
	}
	return dao.DB.Create(u).Error,ok
}

func (u *EntUser)Login() bool {
	var user EntUser
	dao.DB.Where("UserName = ?",u.UserName).Find(&user)
	ok := user.Pwd==u.Pwd
	if ok{
		dao.DB.Model(user).Update(u)
	}
	return ok
}

func (u *EntUser)ChangeName(nick string) error{
	return dao.DB.Model(EntUser{}).Where("UserName = ?",u.UserName).Update("UserName",nick).Error
}

func (u *EntUser)ChangePwd(word string) error{
	return dao.DB.Model(EntUser{}).Where("UserName = ?",u.UserName).Update("UserPwd",word).Error
}

func (u *EntUser)ChangeFace(add string) error{
	return dao.DB.Model(EntUser{}).Where("UserName = ?",u.UserName).Update("FacePicUrl",add).Error
}