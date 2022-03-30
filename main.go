package main

import (
	"api.openfileplatform.com/internal/globals"
	"api.openfileplatform.com/internal/settings"
	"encoding/gob"
	"fmt"
	"github.com/spf13/viper"
	"time"
)

func main() {
	gob.Register(time.Time{})
	var err error
	//初始化viper
	err = settings.InitViper()
	if err != nil {
		fmt.Println("配置文件加载出错！", err)
		return
	}
	var log = globals.GetLogger()

	//初始化数据库（mysql、redis）
	err = settings.InitDatabase()
	if err != nil {
		log.Errorln(err)
		return
	}

	//初始化gin引擎
	engine, err := settings.InitGinEngine()
	if err != nil {
		log.Errorln(err)
		return
	}
	//开始运行
	err = engine.Run(fmt.Sprintf("%s:%s", viper.GetString("system.SysIP"), viper.GetString("system.SysPort")))
	if err != nil {
		log.Errorln(err)
		return
	}
}
