package models

import "api.openfileplatform.com/dao"

type PlatModule struct{
	AutoID int64 `gorm:"AUTO_INCREMENT;column:AutoID;primary_key"`
	ModuleCode string `gorm:"column:ModuleCode"`
	ModuleName string `gorm:"column:ModuleName"`
	ClientType int `gorm:"column:ClientType"`
	IsDeleted bool `gorm:"column:IsDeleted"`
}

type PlatModuleModels interface{
	NewModule()error
	DeleteModule()error
}

func (module *PlatModule)NewModule()error{
	return dao.DB.Model(&PlatModule{}).Create(module).Error
}

func (module *PlatModule) DeleteModule()error{
	return dao.DB.Model(&PlatModule{}).Unscoped().Delete(module).Error
}