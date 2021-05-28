//企业部门信息表

package ent

import "DocumentSystem/dao"

type EntDepartment struct {
	DepartmentID int `gorm:"AUTO_INCREMENT;column:DepartmentID"`
	DepartmentName string	`form:"name" gorm:"column:DepartmentName"`
	DepartmentCode int `gorm:"column:DepartmentCode"`
	HeadID int `gorm:"column:HeadID"`
	Deleted bool `gorm:"column:Deleted"`
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
	return dao.DB.Model(&EntDepartment{}).Where("DepartmentID = ? ",d.DepartmentID).Update("deleted",true).Error
}

func (d *EntDepartment)ReverseDepart()error{
	return dao.DB.Model(&EntDepartment{}).Where("DepartmentID = ? ",d.DepartmentID).Update("deleted",false).Error
}

func (d *EntDepartment)ChangeHeader(id int)error{
	return  dao.DB.Model(&EntDepartment{}).Where("HeadID = ?",d.HeadID).Update("HeadID",id).Error
}