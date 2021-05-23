//企业文档信息

package ent

type EntFileinfo struct{
	Address string
	Name string	`form:"name"`
	TypeId int
}

type EntFileModels interface {
	Add()error
	Find()error
}