//平台用户表

package platform

type PlatEnterprise struct {
	Id int `gorm:"AUTO_INCREMENT"`
	Name string	`form:"name"`
	Admin string	`form:"admin"`
	AdminId int
	Location string	`form:"location"`
	Phone string	`form:"phone"`
	Url string		`form:"url"`
}