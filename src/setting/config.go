// coding: utf-8
// @Author : lryself
// @Date : 2021/4/6 18:44
// @Software: GoLand

package setting

import (
	"api.openfileplatform.com/src/globals"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var log = globals.GetLogger()

func InitViper() (err error) {
	viper.SetConfigName("config")
	viper.AddConfigPath("./sys/configs") // 添加搜索路径
	viper.SetConfigType("yaml")

	err = viper.ReadInConfig()
	if err != nil {
		log.Errorf("Fatal error config file: %s\n", err)
		return
	}
	viper.WatchConfig()

	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Warnf("Config file:%s Op:%s\n", e.Name, e.Op)
	})
	return
}
