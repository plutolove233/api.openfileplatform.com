package main

import (
	"api.openfileplatform.com/src/globals"
	"api.openfileplatform.com/src/models/ginModels"
	"api.openfileplatform.com/src/setting"
	"encoding/gob"
	"github.com/spf13/viper"
	"time"
)

func main() {
	//fmt.Println("program start")
	var log = globals.GetLogger()
	//fmt.Println("日志加载完成")
	gob.Register(ginModels.UserModel{})
	gob.Register(time.Time{})
	var err error
	//初始化viper
	err = setting.InitViper()
	//fmt.Println("viper加载完成")
	if err != nil {
		log.Errorln(err)
		return
	}
	//初始化RSA秘钥
	//setting.InitRSAKey()

	//初始化数据库（mysql、redis）
	err = setting.InitDatabase()
	//fmt.Println("数据库加载完成")
	if err != nil {
		log.Errorln(err)
		return
	}
	//初始化gin引擎
	engine, err := setting.InitGinEngine()
	//fmt.Println("gin引擎加载完成")
	if err != nil {
		log.Errorln(err)
		return
	}
	//监听端口
	port := viper.GetString("system.SysPort")
	err = engine.Run(":" + port)
	//fmt.Println("监听已开启加载完成")
	if err != nil {
		log.Errorln(err)
		return
	}
}
