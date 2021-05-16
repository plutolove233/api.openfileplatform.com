//企业角色表

package ent

import (
	"DocumentSystem/dao"
	"time"
)

type EntRole struct{
	Id int//role的ID
	Code int//角色唯一code代码
	Info string//角色信息
	Name string
	Created time.Time
	Creator string
	Deleted bool
	AuthId int//权限id
}

type EntRoleModel interface {
	NewRole()error
}

func (r *EntRole)NewRole() error{
	return dao.DB.Create(r).Error
}