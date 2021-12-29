// coding: utf-8
// @Author : lryself
// @Date : 2021/12/28 0:02
// @Software: GoLand

package vipers

import (
	"api.openfileplatform.com/src/utils/logs"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"sync"
)

var (
	databaseViper     *viper.Viper
	databaseViperOnce sync.Once
)

var log = logs.GetLogger()

func GetDatabaseViper() *viper.Viper {
	databaseViperOnce.Do(func() {
		databaseViper = viper.New()
		databaseViper.SetConfigName("database")
		databaseViper.AddConfigPath("./config") // 添加搜索路径
		databaseViper.SetConfigType("yaml")

		err := databaseViper.ReadInConfig()
		if err != nil {
			log.Errorf("Fatal error config file: %s \n", err)
			panic(err)
		}

		databaseViper.WatchConfig()

		databaseViper.OnConfigChange(func(e fsnotify.Event) {
			log.Warnf("Config file:%s Op:%s\n", e.Name, e.Op)
		})
	})
	return databaseViper
}
