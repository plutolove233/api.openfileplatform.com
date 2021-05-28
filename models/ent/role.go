//企业角色表

package ent

import (
	"DocumentSystem/dao"
	"time"
)

type EntRole struct{
	RoleID int `gorm:"column:RoleID"`//role的ID
	RoleInfo string `gorm:"column:RoleInfo"`//角色信息
	RoleName string `gorm:"column:RoleName"`
	Created time.Time `gorm:"column:Created"`
	Creator string `gorm:"column:Creator"`
	Deleted bool `gorm:"column:Deleted"`
	AuthID int `gorm:"column:AuthID"`//权限id
}

type EntRoleModel interface {
	NewRole()error
}

func (r *EntRole)NewRole() error{
	return dao.DB.Create(r).Error
}