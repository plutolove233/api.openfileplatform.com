package platform

type PlatFunctionNavigation struct{
	AutoID int64 `gorm:"column:AutoID"`
	NavigationName string `gorm:"column:NavigationName"`
	NavigationCode string `gorm:"column:NavigationCode"`
	ModuleCode string `gorm:"column:ModuleCode"`
	Deleted int `gorm:"column:Deleted"`
}