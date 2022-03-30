// coding: utf-8
// @Author : lryself
// @Date : 2021/4/6 20:40
// @Software: GoLand

package settings

import (
	"api.openfileplatform.com/internal/globals"
	"api.openfileplatform.com/internal/globals/database"
	"github.com/spf13/viper"
)

func InitDatabase() (err error) {
	var log = globals.GetLogger()
	if viper.GetBool("system.UseMysql") {
		err = database.InitMysqlClient()
		if err != nil {
			log.Errorln("mysql初始化出错:", err)
			return
		}
	}
	if viper.GetBool("system.UseRedis") {
		err = database.InitRedisClient()
		if err != nil {
			log.Errorln("redis初始化出错:", err)
			return
		}
	}
	return
}
