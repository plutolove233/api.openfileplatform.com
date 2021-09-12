//企业角色表

package models

import (
	"api.openfileplatform.com/dao"
	"time"
)

type EntRole struct{
	AutoID int64 `gorm:"AUTO_INCREMENT;column:AutoID;primary_key"`
	RoleID int64 `gorm:"column:RoleID"`//role的ID
	RoleInfo string `gorm:"column:RoleInfo"`//角色信息
	RoleName string `gorm:"column:RoleName"`
	CreatTime time.Time `gorm:"column:CreatTime"`
	UserID string `gorm:"column:UserID"`//创建人ID
	IsDeleted bool `gorm:"column:IsDeleted"`
}

type EntRoleModel interface {
	NewRole()error
}

func (r *EntRole)NewRole() error{
	return dao.DB.Model(&EntRole{}).Create(r).Error
}