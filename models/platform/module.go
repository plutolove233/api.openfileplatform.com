package platform

type PlatModule struct{
	AutoID int `gorm:"column:AutoID"`
	ModuleCode string `gorm:"column:ModuleCode"`
	ModuleName string `gorm:"column:ModuleName"`
	ClientType int `gorm:"column:ClientType"`
	Deleted int `gorm:"column:Deleted"`
}
