package models

type PlatFunctionNavigation struct{
	AutoID int64 `gorm:"AUTO_INCREMENT;column:AutoID;primary_key"`
	NavigationName string `gorm:"column:NavigationName"`
	NavigationCode string `gorm:"column:NavigationCode"`
	ModuleCode string `gorm:"column:ModuleCode"`
	IsDeleted int `gorm:"column:IsDeleted"`
}