package main

import (
	"api.openfileplatform.com/dao"
	"api.openfileplatform.com/settings"
	"fmt"
)

func main() {
	err:= dao.StartMysql()
	if err!=nil{
		fmt.Println("Failed to open Mysql")
		return
	}
	defer dao.DB.Close()

	//settings.InitRSAKey()

	engine,err := settings.InitEngine()
	err2 :=  engine.Run(":9000")
	if err2!=nil{
		fmt.Println("Engine start error")
		return
	}

}
