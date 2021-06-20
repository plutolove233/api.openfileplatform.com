package configs

import (
	"fmt"
	"github.com/spf13/viper"
	"sync"
)

type Config struct {
	Mysql MysqlConfig
	Reddis RedisConfig
}

type MysqlConfig struct{
	Host string
	Ports int
	Username string
	Password string
	Database string
	Charset string
}

type RedisConfig struct{
	Host string
	Port int
	Password string
}

var once sync.Once
var c *Config

func (p *Config)initConfig(){
	viper.SetConfigFile("configs/database.toml")
	if err:=viper.ReadInConfig(); err!=nil {
		fmt.Println("Failed to load toml file")
	}
	if err:=viper.Unmarshal(p); err!=nil {
		fmt.Println("Failed to unmarshal toml file")
	}
}

func GetConfig()(string,string){
	once.Do(func(){
		c = &Config{}
		c.initConfig()
	})
	order:=fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",c.Mysql.Username,
		c.Mysql.Password,c.Mysql.Host,c.Mysql.Ports,c.Mysql.Database)
	return "mysql",order
}