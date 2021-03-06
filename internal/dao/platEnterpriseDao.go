package dao

import (
	"api.openfileplatform.com/internal/globals/database"
	"api.openfileplatform.com/internal/models/mysqlModel"
)

type PlatEnterprise struct {
	mysqlModel.PlatEnterprises
}

func (m *PlatEnterprise) Get() error {
	mysqlMamager := database.GetMysqlClient()
	return mysqlMamager.Where(map[string]interface{}{
		"IsDeleted": 0,
	}).Where(m).Take(m).Error
}

func (m *PlatEnterprise) Add() error {
	mysqlMamager := database.GetMysqlClient()
	return mysqlMamager.Create(&m).Error
}

func (m *PlatEnterprise) Update(args map[string]interface{}) error {
	mysqlManager := database.GetMysqlClient()
	err := m.Get()
	if err != nil {
		return err
	}
	return mysqlManager.Model(&m).Updates(args).Error
}

func (m *PlatEnterprise) Delete() error {
	mysqlManager := database.GetMysqlClient()
	err := m.Get()
	if err != nil {
		return err
	}
	return mysqlManager.Model(&m).Updates(map[string]interface{}{
		"IsDeleted": 1,
	}).Error

}

func (*PlatEnterprise) GetAll() (error, []PlatEnterprise) {
	mysqlManager := database.GetMysqlClient()
	enterprise := []PlatEnterprise{}
	return mysqlManager.Model(&PlatEnterprise{}).Where(map[string]interface{}{
		"IsDeleted":false,
	}).Find(&enterprise).Error, enterprise
}
