//企业用户角色表

package ent

import (
	"DocumentSystem/dao"
	"time"
)

type EntUserRole struct{
	Id int `gorm:"AUTO_INCREMENT"`
	UserId int
	RoleId int
	Created time.Time
	Creator string
	Edited time.Time
	Deleted bool
}

type EntUserRoleModels interface {
	LinkUserRole(user *EntUser,role *EntRole)(error,bool)
	DeleteUserRole()error
	ReverseUserRole()error
}

func (ur *EntUserRole)LinkUserRole(u *EntUser,r *EntRole)(error,bool){
	ur.RoleId = r.Id
	ur.UserId = u.Id
	ur.Created = time.Now()
	var p EntUserRole
	dao.DB.Where("role_id = ?", r.Id).Find(&p)
	if p.UserId == u.Id && p.RoleId == r.Id{
		return nil,false
	}
	return dao.DB.Create(ur).Error,true
}

func (ur *EntUserRole)DeleteUserRole() error{
	return dao.DB.Model(EntUserRole{}).Where("id = ?",ur.Id).Update("deleted",true).Error
}

func (ur *EntUserRole)ReverseUserRole()error{
	return dao.DB.Model(EntUserRole{}).Where("id = ?",ur.Id).Update("deleted",false).Error
}