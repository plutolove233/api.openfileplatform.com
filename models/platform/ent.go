//平台用户表

package platform

type PlatEnterprise struct {
	EnterpriseID int `gorm:"AUTO_INCREMENT;column:EnterpriseID"`
	EnterpriseName string	`form:"name" gorm:"column:EnterpriseName"`
	Admin string	`form:"admin" gorm:"column:Admin"`
	AdminID int `gorm:"column:AdminID"`
	Location string	`form:"location" gorm:"column:Location"`
	EnterprisePhone string	`form:"phone" gorm:"column:EnterprisePhone"`
	EnterpriseUrl string		`form:"url" gorm:"column:EnterpriseUrl"`
}