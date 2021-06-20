//平台用户表

package models

type PlatEnterprise struct {
	AutoID int64 `gorm:"AUTO_INCREMENT;column:AutoID;primary_key"`
	EnterpriseID int64 `gorm:"column:EnterpriseID"`
	EnterpriseName string	`form:"name" gorm:"column:EnterpriseName"`
	AdminID int `gorm:"column:AdminID"`
	Location string	`form:"location" gorm:"column:Location"`
	EnterprisePhone string	`form:"phone" gorm:"column:EnterprisePhone"`
	EnterpriseUrl string		`form:"url" gorm:"column:EnterpriseUrl"`
}