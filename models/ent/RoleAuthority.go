//企业角色权限表

package ent

import (
	"DocumentSystem/dao"
	"time"
)

type EntRoleAuthority struct{
	Id int `gorm:"AUTO_INCREMENT"`
	RoleId int `form:"role_id"`//角色ID
	AuthId int `form:"permission_id"`//权限ID
	Created time.Time //创建时间
	Creator string //创建人
	Edited time.Time//修改时间
	Deleted bool
}

type EntRoleAuthorityModels interface {
	LinkRoleAuth(r *EntRole,p *EntAuthority)(error,bool)
	DeleteRoleAuth()error
	ReverseRoleAuth()error
	ModifyRoleAuth()error
}

func (rp *EntRoleAuthority)LinkRoleAuth(role *EntRole,per *EntAuthority) (error,bool){
	rp.RoleId = role.Id
	rp.AuthId = per.Id
	rp.Created = time.Now()
	var p EntRoleAuthority
	dao.DB.Where("role_id = ?",role.Id).Find(&p)
	if p.RoleId==role.Id && p.AuthId==per.Id{
		return nil,false
	}
	return dao.DB.Create(rp).Error,true
}

func (rp *EntRoleAuthority)DeleteRoleAuth() error{
	return dao.DB.Model(EntRoleAuthority{}).Where("id = ?",rp.Id).Update("deleted",true).Error
}

func (rp *EntRoleAuthority)ReverseRoleAuth()error{
	return dao.DB.Model(EntRoleAuthority{}).Where("id = ?",rp.Id).Update("deleted",false).Error
}

func (rp *EntRoleAuthority)ModifyRoleAuth()error{
	var rp1 EntRoleAuthority
	dao.DB.Where("id = ?",rp.Id).Find(&rp1)
	return dao.DB.Model(rp1).Update(rp).Error
}