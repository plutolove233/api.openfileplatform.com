//企业部门信息表

package ent

import "DocumentSystem/dao"

type EntDepartment struct {
	Id int `gorm:"AUTO_INCREMENT"`
	Name string	`form:"name"`
	Code int
	HeadId int
	Deleted bool
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
	return dao.DB.Model(&EntDepartment{}).Where("id = ? ",d.Id).Update("deleted",true).Error
}

func (d *EntDepartment)ReverseDepart()error{
	return dao.DB.Model(&EntDepartment{}).Where("id = ? ",d.Id).Update("deleted",false).Error
}

func (d *EntDepartment)ChangeHeader(id int)error{
	return  dao.DB.Model(&EntDepartment{}).Where("head_id = ?",d.HeadId).Update("head_id",id).Error
}