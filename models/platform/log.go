package platform

import "time"

type PlatSystemLog struct{
	AutoID int64 `gorm:"AUTO_INCREMENT;column:AutoID;primary_key"`
	UserID int64 `gorm:"column:UserID"`
	UserName string `gorm:"column:UserName"`
	OperationIP string `gorm:"column:OperationIP"`//发出操作的IP地址
	OperationType string `gorm:"column:OperationType"`//操作类型：1:正常登录；2 登录异常（密码出错超出次数）；3 Ip地址异常；4-非法链接访问
	OperationContent string `gorm:"column:OperationContent"`
	OperationResult int `gorm:"column:OperationResult"`//0--failed 1--succeed
	OperationStatus int `gorm:"column:OperationStatus"`//0--error 1--regular
	CreateTime time.Time `gorm:"column:CreateTime"`
}