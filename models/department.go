//企业部门信息表

package models

import "api.openfileplatform.com/dao"

type EntDepartment struct {
	AutoID int64 `gorm:"AUTO_INCREMENT;column:AutoID;primary_key"`
	DepartmentID int64 `gorm:"column:DepartmentID"`
	DepartmentName string	`form:"name" gorm:"column:DepartmentName"`
	DepartmentCode int `gorm:"column:DepartmentCode"`
	HeadID int `gorm:"column:HeadID"`
	IsDeleted bool `gorm:"column:IsDeleted"`
}

type EntDepartmentModels interface {
	AddDepart()error
	ChangeHeader(id int)error
	DeleteDepart()error
	ReverseDepart()error
}

func (d *EntDepartment)AddDepart()error{
	return dao.DB.Create(d).Error
}

func (d *EntDepartment)DeleteDepart()error{
	return dao.DB.Model(&EntDepartment{}).Where("DepartmentID = ? ",d.DepartmentID).Update("IsDeleted",true).Error
}

func (d *EntDepartment)ReverseDepart()error{
	return dao.DB.Model(&EntDepartment{}).Where("DepartmentID = ? ",d.DepartmentID).Update("IsDeleted",false).Error
}

func (d *EntDepartment)ChangeHeader(id int)error{
	return  dao.DB.Model(&EntDepartment{}).Where("HeadID = ?",d.HeadID).Update("HeadID",id).Error
}