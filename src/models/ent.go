//平台用户表(企业端)

package models

import (
	"api.openfileplatform.com/dao"
)

type PlatEnterprise struct {
	AutoID int64 `gorm:"AUTO_INCREMENT;column:AutoID;primary_key"`
	EnterpriseID int64 `gorm:"column:EnterpriseID"`
	EnterpriseName string	`form:"name" gorm:"column:EnterpriseName"`
	EnterprisePwd string `form:"pwd" gorm:"column:EnterprisePwd"`
	AdminID int `gorm:"column:AdminID"`
	Location string	`form:"location" gorm:"column:Location"`
	EnterprisePhone string	`form:"phone" gorm:"column:EnterprisePhone"`
	EnterpriseUrl string		`form:"url" gorm:"column:EnterpriseUrl"`
	LogoPicUrl string `gorm:"column:LogoPicUrl"`
}

type PlatEnterpriseModels interface{
	Add()error
	Delete()error
	ChangeAddress(address string)error
	ChangePassword(pwd string)error
	ChangeName(name string )error
	ChangeLeader(id int)error
	ChangeLogo()error
}

func (ent *PlatEnterprise)Add()error{
	return dao.DB.Model(&PlatEnterprise{}).Create(ent).Error
}

func (ent *PlatEnterprise)Delete()error{
	return dao.DB.Model(&PlatEnterprise{}).Unscoped().Delete(ent).Error
}

func (ent *PlatEnterprise)ChangeAddress(address string) error{
	return dao.DB.Model(PlatEnterprise{}).Where("EnterpriseID = ?",ent.EnterpriseID).Update("Location",address).Error
}

func (ent *PlatEnterprise)ChangePassword(pwd string)error{
	return dao.DB.Model(&PlatEnterprise{}).Where("EnterpriseID = ?",ent.EnterpriseID).Update("EnterprisePwd",pwd).Error
}

func (ent *PlatEnterprise)ChangeName(name string)error{
	return dao.DB.Model(&PlatEnterprise{}).Where("EnterpriseID = ?",ent.EnterpriseID).Update("EnterpriseName",name).Error
}

func (ent *PlatEnterprise)ChangeLeader(id int)error{
	return dao.DB.Model(&PlatEnterprise{}).Where("EnterpriseID = ?",ent.EnterpriseID).Update("AdminID",id).Error
}

func (ent *PlatEnterprise)ChangeLogo(dst string)error{
	return dao.DB.Model(&PlatEnterprise{}).Where("EnterpriseID = ?",ent.EnterpriseID).Update("logoPicUrl",dst).Error
}