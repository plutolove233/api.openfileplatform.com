//员工信息表

package ent

import (
	"DocumentSystem/dao"
	"time"
)

type EntStaff struct {
	StaffID int `gorm:"AUTO_INCREMENT;column:StaffID"`
	StaffName string	`form:"name" gorm:"column:StaffName"`
	StaffPhone int	`form:"phone" gorm:"column:StaffPhone"`
	Salary int	`form:"salary" gorm:"column:Salary"`
	DepartID int	`form:"depart_id" gorm:"column:DepartID"`
	College string	`form:"college" gorm:"column:College"`
	Status int 	`form:"status" gorm:"column:Status"`//员工是否转正 1--正式员工 0--实习生
	EntID int	`form:"ent_id" gorm:"column:EntID"`
	Deleted bool `gorm:"column:Deleted"`
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
	return dao.DB.Model(&EntStaff{}).Where("StaffID = ?",s.StaffID).Update("DepartID",d).Error
}

func (s *EntStaff)TransEmployee()error{
	return dao.DB.Model(&EntStaff{}).Where("StaffID = ?",s.StaffID).Update("Status",1).Error
}