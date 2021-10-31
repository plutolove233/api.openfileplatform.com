package models

import "time"

type EntUserLog struct{
	AutoID int64 `gorm:"AUTO_INCREMENT;column:AutoID;primary_key"`
	UserID int64 `gorm:"column:UserID"`
	UserName string `gorm:"column:UserName"`
	Account string `gorm:"column:Account"`//登录账号
	EnterpriseID int64 `gorm:"column:EnterpriseID"`
	OperationIP string `gorm:"column:OperationIP"`//发出操作的IP地址
	OperationType string `gorm:"column:OperationType"`//操作类型:1创建、2修改、3删除 4审核等操作
	OperationContent string `gorm:"column:OperationContent"`
	OperationResult int `gorm:"column:OperationResult"`//0--failed 1--succeed
	OperationStatus int `gorm:"column:OperationStatus"`//0--error 1--regular
	CreateTime time.Time `gorm:"column:CreateTime"`
}
