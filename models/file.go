//企业文档信息

package models

import "api.openfileplatform.com/dao"

type EntFileinfo struct{
	AutoID int64 `gorm:"AUTO_INCREMENT;column:AutoID;primary_key"`
	FileAddress string `gorm:"column:FileAddress"`
	FileName string	`form:"name" gorm:"column:FileName"`
	TypeID int `gorm:"column:TypeID"`

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