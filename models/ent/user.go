//企业用户表

package ent

import (
	"DocumentSystem/dao"
	"time"
)

type EntUser struct{
	EntUserID int `gorm:"AUTO_INCREMENT;column:EntUserID"`
	IsAdmin int `gorm:"column:IsAdmin"`
	EntID int `gorm:"column:EntID"`
	EntUserName string `form:"username" gorm:"column:EntUserName"`
	EntUserPwd string `form:"password" gorm:"column:EntUserPwd"`
	Company string `form:"company" gorm:"column:Company"`
	EntUserPhone string `form:"phone" gorm:"column:EntUserPhone"`
	EntUserEmail string `form:"email" gorm:"column:EntUserEmail"`
	LastLoginTime time.Time `form:"lastlogintime" gorm:"column:LastLoginTime"`
	LoginTime time.Time `form:"logintime" gorm:"column:LoginTime"`
	Times int64 `form:"times" gorm:"column:Times"`
	Face string `gorm:"column:Face"`//user avatar like the address
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
	dao.DB.Where("EntUserName = ?",u.EntUserName).Find(&user)
	ok := user.EntUserPwd==u.EntUserPwd
	if ok{
		u.LastLoginTime = u.LoginTime
		u.LoginTime = time.Now()
		u.Times++
		dao.DB.Model(user).Update(u)
	}
	return ok
}

func (u *EntUser)ChangeName(nick string) error{
	return dao.DB.Model(EntUser{}).Where("EntUserName = ?",u.EntUserName).Update("EntUserName",nick).Error
}

func (u *EntUser)ChangePwd(word string) error{
	return dao.DB.Model(EntUser{}).Where("EntUserName = ?",u.EntUserName).Update("EntUserPwd",word).Error
}

func (u *EntUser)ChangeFace(add string) error{
	return dao.DB.Model(EntUser{}).Where("EntUserName = ?",u.EntUserName).Update("Face",add).Error
}