//企业文档信息

package models

import (
	"api.openfileplatform.com/dao"
	"time"
)

type EntFileinfo struct{
	AutoID int64 `gorm:"AUTO_INCREMENT;column:AutoID;primary_key"`
	FileAddress string `gorm:"column:FileAddress"`
	FileName string	`form:"name" gorm:"column:FileName"`
	TypeID int `gorm:"column:TypeID"`
	UpTime time.Time `gorm:"column:UpTime"`
	BorrowTimes int64	`gorm:"column:BorrowTimes"`
	Status int `gorm:"column:Status"`//0表示没有被借出，1表示已经借出
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