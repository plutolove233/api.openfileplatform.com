//系统用户表

package models

import "DocumentSystem/dao"

type NormalUsers struct{
	Id int `gorm:"AUTO_INCREMENT"`
	Name string `form:"name"`
	Pwd string	`form:"pwd"`
	Phone string	`form:"phone"`
	Email string	`form:"email"`
}

func (c *NormalUsers)Login()bool{
	var n NormalUsers
	dao.DB.Where("name = ?",c.Name).Find(&n)
	ok := n.Pwd==c.Pwd
	return ok
}

func (c *NormalUsers)Register()(error,bool){
	var user []NormalUsers
	ok := true
	dao.DB.Find(&user)
	for _,index := range user{
		if index==*c{
			ok = false
			return nil,ok
		}
	}
	return dao.DB.Create(c).Error,ok
}