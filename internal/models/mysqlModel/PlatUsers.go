// coding: utf-8
// @Author : lryself
// @Date : 2022/2/9 19:22
// @Software: GoLand

package mysqlModel

import "time"

// PlatUsers [...]
type PlatUsers struct {
	AutoID     int64     `gorm:"primaryKey;column:AutoID;type:bigint(22);not null" json:"-"`
	UserID     string    `gorm:"column:UserID;type:varchar(20)" json:"userId" form:"userId"`
	UserName   string    `gorm:"column:UserName;type:varchar(255)" json:"userName" form:"userName"`
	Account    string    `gorm:"column:Account;type:varchar(255)" json:"account" form:"account"`
	Password   string    `gorm:"column:Password;type:varchar(255)" json:"password" form:"password"`
	Phone      string    `gorm:"column:Phone;type:varchar(255)" json:"phone" form:"phone"`
	Email      string    `gorm:"column:Email;type:varchar(255)" json:"email" form:"email"`
	IsDeleted  bool      `gorm:"column:IsDeleted;type:tinyint(1)" json:"isDeleted" form:"isDeleted"`
	CreateTime time.Time `gorm:"column:CreateTime;type:timestamp;default:CURRENT_TIMESTAMP" json:"createTime" form:"createTime"`
}

// TableName get sql table name.获取数据库表名
func (m *PlatUsers) TableName() string {
	return "plat_users"
}

// PlatUsersColumns get sql column name.获取数据库列名
var PlatUsersColumns = struct {
	AutoID     string
	UserID     string
	UserName   string
	Account    string
	Password   string
	Phone      string
	Email      string
	IsDeleted  string
	CreateTime string
}{
	AutoID:     "AutoID",
	UserID:     "UserID",
	UserName:   "UserName",
	Account:    "Account",
	Password:   "Password",
	Phone:      "Phone",
	Email:      "Email",
	IsDeleted:  "IsDeleted",
	CreateTime: "CreateTime",
}

func (m *PlatUsers) GetUserID() string {
	return m.UserID
}
func (m *PlatUsers) SetUserID(id string) {
	m.UserID = id
}
func (m *PlatUsers) GetIsAdmin() bool {
	return true
}
func (m *PlatUsers) SetAccount(account string) {
	m.Account = account
}
func (m *PlatUsers) GetPassword() string {
	return m.Password
}
func (m *PlatUsers) SetUserName(p string) {
	m.UserName = p
}
func (m *PlatUsers) SetPassword(p string) {
	m.Password = p
}
func (m *PlatUsers) SetPhone(p string) {
	m.Phone = p
}
func (m *PlatUsers) SetEmail(p string) {
	m.Email = p
}
