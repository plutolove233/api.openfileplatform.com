package dao

import (
	"DocumentSystem/commons/configs"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func StartMysql()(err error){
	DB,err = gorm.Open(configs.GetConfig())
	return err
}