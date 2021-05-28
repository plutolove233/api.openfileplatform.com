//企业用户权限表

package ent

import "DocumentSystem/dao"

type EntAuthority struct {
	AuthorityID int `gorm:"column:AuthorityID"`
	AuthorityUrl string `gorm:"column:AuthorityUrl"`
	AuthorityInfo string `gorm:"column:AuthorityInfo"`//权限信息
	FunctionCode string `gorm:"column:FunctionCode"`
	FunctionUrl string `gorm:"column:FunctionUrl"`
}

type EntAuthorityModels interface {
	NewPer() error
}

func (per *EntAuthority)NewPer() error{
	return dao.DB.Create(per).Error
}
