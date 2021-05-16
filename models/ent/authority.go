//企业用户权限表

package ent

import "DocumentSystem/dao"

type EntAuthority struct {
	Id int
	Url string
	Code int//权限唯一code代码
	Info string//权限信息
}

type EntAuthorityModels interface {
	NewPer() error
}

func (per *EntAuthority)NewPer() error{
	return dao.DB.Create(per).Error
}
