//企业角色权限表

package ent

import (
	"DocumentSystem/dao"
	"time"
)

type EntRoleAuthority struct{
	AutoID int64 `gorm:"AUTO_INCREMENT;column:AutoID;primary_key"`
	RoleID int64 `form:"role_id" gorm:"column:RoleID"`//角色ID
	AuthorityID int `form:"permission_id" gorm:"column:AuthorityID"`//权限ID
	CreatTime time.Time `gorm:"column:CreatTime"`//创建时间
	Creator string `gorm:"column:Creator"`//创建人
	IsDeleted bool `gorm:"column:IsDeleted"`
	FunctionCode string `gorm:"column:FunctionCode"`
	FunctionUrl string `gorm:"column:FunctionUrl"`
}

type EntRoleAuthorityModels interface {
	LinkRoleAuth(r *EntRole)error
	DeleteRoleAuth()error
	ReverseRoleAuth()error
	ModifyRoleAuth()error
}

func (rp *EntRoleAuthority)LinkRoleAuth() error{
	return dao.DB.Create(rp).Error
}

func (rp *EntRoleAuthority)DeleteRoleAuth() error{
	return dao.DB.Model(EntRoleAuthority{}).Where("AutoID = ?",rp.AutoID).Update("IsDeleted",true).Error
}

func (rp *EntRoleAuthority)ReverseRoleAuth()error{
	return dao.DB.Model(EntRoleAuthority{}).Where("AutoID = ?",rp.AutoID).Update("IsDeleted",false).Error
}

func (rp *EntRoleAuthority)ModifyRoleAuth()error{
	var rp1 EntRoleAuthority
	dao.DB.Where("AutoID = ?",rp.AutoID).Find(&rp1)
	return dao.DB.Model(rp1).Update(rp).Error
}