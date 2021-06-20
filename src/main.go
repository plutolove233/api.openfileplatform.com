package src

import (
	"api.openfileplatform.com/dao"
	"api.openfileplatform.com/models"
	"api.openfileplatform.com/settings"
	"fmt"
)

func ReadyDataBase(){
	dao.DB.AutoMigrate(&models.EntDepartment{})
	dao.DB.AutoMigrate(&models.EntFileinfo{})
	dao.DB.AutoMigrate(&models.EntRole{})
	dao.DB.AutoMigrate(&models.EntRoleAuthority{})
	dao.DB.AutoMigrate(&models.EntStaff{})
	dao.DB.AutoMigrate(&models.EntUser{})
	dao.DB.AutoMigrate(&models.EntUserLog{})
	dao.DB.AutoMigrate(&models.PlatEnterprise{})
	dao.DB.AutoMigrate(&models.PlatUser{})
	dao.DB.AutoMigrate(&models.PlatFunctionNavigation{})
	dao.DB.AutoMigrate(&models.PlatFunction{})
	dao.DB.AutoMigrate(&models.PlatModule{})
	dao.DB.AutoMigrate(&models.PlatSystemLog{})
}

func main() {
	err:= dao.StartMysql()
	if err!=nil{
		fmt.Println("Failed to open Mysql")
		return
	}
	defer dao.DB.Close()

	ReadyDataBase()

	//settings.InitRSAKey()

	engine,err := settings.InitEngine()
	err2 :=  engine.Run(":80")
	if err2!=nil{
		fmt.Println("Engine start error")
		return
	}

}
