//企业文档信息

package ent

type EntFileinfo struct{
	Address string
	Name string	`form:"name"`
	TypeId int
}