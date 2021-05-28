//企业角色权限表

package ent

import (
	"DocumentSystem/dao"
	"time"
)

type EntRoleAuthority struct{
	RoleAuthorityID int `gorm:"AUTO_INCREMENT;column:RoleAuthorityID"`
	RoleID int `form:"role_id" gorm:"column:RoleID"`//角色ID
	AuthorityID int `form:"permission_id" gorm:"column:AuthorityID"`//权限ID
	Created time.Time `gorm:"column:Created"`//创建时间
	Creator string `gorm:"column:Creator"`//创建人
	Edited time.Time `gorm:"column:Edited"`//修改时间
	Deleted bool `gorm:"column:Deleted"`
	FunctionCode string `gorm:"column:FunctionCode"`
}

type EntRoleAuthorityModels interface {
	LinkRoleAuth(r *EntRole,p *EntAuthority)(error,bool)
	DeleteRoleAuth()error
	ReverseRoleAuth()error
	ModifyRoleAuth()error
}

func (rp *EntRoleAuthority)LinkRoleAuth(role *EntRole,per *EntAuthority) (error,bool){
	rp.RoleID = role.RoleID
	rp.AuthorityID = per.AuthorityID
	rp.Created = time.Now()
	var p EntRoleAuthority
	dao.DB.Where("role_id = ?",role.RoleID).Find(&p)
	if p.RoleID==role.RoleID && p.AuthorityID==per.AuthorityID{
		return nil,false
	}
	return dao.DB.Create(rp).Error,true
}

func (rp *EntRoleAuthority)DeleteRoleAuth() error{
	return dao.DB.Model(EntRoleAuthority{}).Where("RoleAuthorityID = ?",rp.RoleAuthorityID).Update("deleted",true).Error
}

func (rp *EntRoleAuthority)ReverseRoleAuth()error{
	return dao.DB.Model(EntRoleAuthority{}).Where("RoleAuthorityID = ?",rp.RoleAuthorityID).Update("deleted",false).Error
}

func (rp *EntRoleAuthority)ModifyRoleAuth()error{
	var rp1 EntRoleAuthority
	dao.DB.Where("RoleAuthorityID = ?",rp.RoleAuthorityID).Find(&rp1)
	return dao.DB.Model(rp1).Update(rp).Error
}