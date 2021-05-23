//企业文档信息

package ent

import "DocumentSystem/dao"

type EntFileinfo struct{
	Address string
	Name string	`form:"name"`
	TypeId int

}

type EntFileModels interface {
	Add()error
	Find()[]EntFileinfo
}

func (f *EntFileinfo)Add()error{
	return dao.DB.Create(f).Error
}

func (f* EntFileinfo)Find() []EntFileinfo{
	var f1 []EntFileinfo
	dao.DB.Model(&EntFileinfo{}).Find(&f1)
	return f1
}