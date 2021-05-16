//企业用户表

package ent

import (
	"DocumentSystem/dao"
	"time"
)

type EntUser struct{
	Id int `gorm:"AUTO_INCREMENT"`
	IsAdmin int
	EntId int
	Name string `form:"username"`
	Pwd string `form:"password"`
	Company string `form:"company"`
	Phone string `form:"phone"`
	Email string `form:"email"`
	LastLoginTime time.Time `form:"lastlogintime"`
	LoginTime time.Time `form:"logintime"`
	Times int64 `form:"times"`
	Face string //user avatar like the address
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
	dao.DB.Where("name = ?",u.Name).Find(&user)
	ok := user.Pwd==u.Pwd
	if ok{
		u.LastLoginTime = u.LoginTime
		u.LoginTime = time.Now()
		u.Times++
		dao.DB.Model(user).Update(u)
	}
	return ok
}

func (u *EntUser)ChangeName(nick string) error{
	return dao.DB.Model(EntUser{}).Where("name = ?",u.Name).Update("name",nick).Error
}

func (u *EntUser)ChangePwd(word string) error{
	return dao.DB.Model(EntUser{}).Where("name = ?",u.Name).Update("pwd",word).Error
}

func (u *EntUser)ChangeFace(add string) error{
	return dao.DB.Model(EntUser{}).Where("name = ?",u.Name).Update("face",add).Error
}