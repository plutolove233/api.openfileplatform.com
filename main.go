package main

import (
	"DocumentSystem/dao"
	"DocumentSystem/models"
	"DocumentSystem/models/ent"
	"DocumentSystem/models/platform"
	"DocumentSystem/router"
	"fmt"
)

func main() {
	err:= dao.StartMysql()
	if err!=nil{
		fmt.Println("Failed to open Mysql")
		return
	}
	defer dao.DB.Close()

	dao.DB.AutoMigrate(&ent.EntAuthority{})
	dao.DB.AutoMigrate(&ent.EntDepartment{})
	dao.DB.AutoMigrate(&ent.EntFileinfo{})
	dao.DB.AutoMigrate(&ent.EntRole{})
	dao.DB.AutoMigrate(&ent.EntRoleAuthority{})
	dao.DB.AutoMigrate(&ent.EntStaff{})
	dao.DB.AutoMigrate(&ent.EntUser{})
	dao.DB.AutoMigrate(&ent.EntUserRole{})
	dao.DB.AutoMigrate(&platform.PlatEnterprise{})
	dao.DB.AutoMigrate(&platform.PlatUser{})
	dao.DB.AutoMigrate(&models.NormalUsers{})

	engine := router.StartEngine()
	err2 :=  engine.Run(":9090")
	if err2!=nil{
		fmt.Println("Engine start error")
		return
	}

}
