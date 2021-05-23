//员工信息表

package ent

import (
	"DocumentSystem/dao"
	"time"
)

type EntStaff struct {
	Id int `gorm:"AUTO_INCREMENT"`
	Name string	`form:"name"`
	Phone int	`form:"phone"`
	Salary int	`form:"salary"`
	DepartId int	`form:"depart_id"`
	College string	`form:"college"`
	Status int 	`form:"status"`//员工是否转正 1--正式员工 0--实习生
	EntId int	`form:"ent_id"`
	Deleted bool
	CreateTime time.Time
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
	return dao.DB.Model(&EntStaff{}).Where("id = ?",s.Id).Update("salary",s.Salary+much).Error
}

func (s *EntStaff)DecreaseSalary(much int)error{
	return dao.DB.Model(&EntStaff{}).Where("id = ?",s.Id).Update("salary",s.Salary-much).Error
}

func (s *EntStaff)TransDepart(d int)error{
	return dao.DB.Model(&EntStaff{}).Where("id = ?",s.Id).Update("depart_id",d).Error
}

func (s *EntStaff)TransEmployee()error{
	return dao.DB.Model(&EntStaff{}).Where("id = ?",s.Id).Update("status",1).Error
}