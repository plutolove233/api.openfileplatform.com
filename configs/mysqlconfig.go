package configs

import(
	"github.com/spf13/viper"
	"fmt"
	"sync"
)

type Config struct {
	Mysql MysqlConfig
}

type MysqlConfig struct{
	Host string
	Ports int
	Username string
	Password string
	Database string
	Charset string
}

var once sync.Once
var c *Config

func (p *Config)initConfig(){
	viper.SetConfigFile("configs/config.toml")
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