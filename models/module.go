package models

type PlatModule struct{
	AutoID int64 `gorm:"AUTO_INCREMENT;column:AutoID;primary_key"`
	ModuleCode string `gorm:"column:ModuleCode"`
	ModuleName string `gorm:"column:ModuleName"`
	ClientType int `gorm:"column:ClientType"`
	IsDeleted bool `gorm:"column:IsDeleted"`
}
