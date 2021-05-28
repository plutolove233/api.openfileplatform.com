//企业用户角色表

package ent

import (
	"DocumentSystem/dao"
	"time"
)

type EntUserRole struct{
	UserRoleID int `gorm:"AUTO_INCREMENT;column:UserRoleID"`
	EntUserID int `gorm:"column:EntUserID"`
	RoleID int `gorm:"column:RoleID"`
	Created time.Time `gorm:"column:Created"`
	Creator string `gorm:"column:Creator"`
	Edited time.Time `gorm:"column:Edited"`
	Deleted bool `gorm:"column:Deleted"`
}

type EntUserRoleModels interface {
	LinkUserRole(user *EntUser,role *EntRole)(error,bool)
	DeleteUserRole()error
	ReverseUserRole()error
}

func (ur *EntUserRole)LinkUserRole(u *EntUser,r *EntRole)(error,bool){
	ur.RoleID = r.RoleID
	ur.EntUserID = u.EntUserID
	ur.Created = time.Now()
	var p EntUserRole
	dao.DB.Where("role_id = ?", r.RoleID).Find(&p)
	if p.EntUserID == u.EntUserID && p.RoleID == r.RoleID{
		return nil,false
	}
	return dao.DB.Create(ur).Error,true
}

func (ur *EntUserRole)DeleteUserRole() error{
	return dao.DB.Model(EntUserRole{}).Where("id = ?",ur.UserRoleID).Update("deleted",true).Error
}

func (ur *EntUserRole)ReverseUserRole()error{
	return dao.DB.Model(EntUserRole{}).Where("id = ?",ur.UserRoleID).Update("deleted",false).Error
}