//员工信息表

package models

import (
	"DocumentSystem/dao"
	"time"
)

type EntStaff struct {
	AutoID int64 `gorm:"AUTO_INCREMENT;column:AutoID;primary_key"`
	StaffID int64 `gorm:"column:StaffID"`
	StaffName string	`form:"name" gorm:"column:StaffName"`
	StaffPhone int	`form:"phone" gorm:"column:StaffPhone"`
	Salary int	`form:"salary" gorm:"column:Salary"`
	DepartmentID int64	`form:"depart_id" gorm:"column:DepartmentID"`
	Status int 	`form:"status" gorm:"column:Status"`//员工是否转正 0--正式员工 1--临时工
	EnterpriseID int64	`form:"ent_id" gorm:"column:EntID"`
	IsDeleted bool `gorm:"column:Deleted"`
	CreateTime time.Time `gorm:"column:CreateTime"`
}

type EntStaffModels interface {
	AddStaff()error
	DeleteStaff()error
	IncreaseSalary(much int)error
	DecreaseSalary(much int)error
	TransDepart(d int)error
	TransEmployee()error
}

func (s *EntStaff)AddStaff()error{
	return dao.DB.Create(s).Error
}

func (s *EntStaff)DeleteStaff()error{
	return dao.DB.Unscoped().Delete(s).Error
}

func (s *EntStaff)IncreaseSalary(much int)error{
	return dao.DB.Model(&EntStaff{}).Where("StaffID = ?",s.StaffID).Update("Salary",s.Salary+much).Error
}

func (s *EntStaff)DecreaseSalary(much int)error{
	return dao.DB.Model(&EntStaff{}).Where("StaffID = ?",s.StaffID).Update("Salary",s.Salary-much).Error
}

func (s *EntStaff)TransDepart(d int)error{
	return dao.DB.Model(&EntStaff{}).Where("StaffID = ?",s.StaffID).Update("DepartmentID",d).Error
}

func (s *EntStaff)TransEmployee()error{
	return dao.DB.Model(&EntStaff{}).Where("StaffID = ?",s.StaffID).Update("Status",1).Error
}