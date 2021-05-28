package platform


type PlatFunction struct{
	AutoID int `gorm:"column:AutoID"`
	FunctionCode string `gorm:"column:FunctionCode"`
	FunctionName string `gorm:"column:FunctionName"`
	FunctionUrl string `gorm:"column:FunctionUrl"`
	ModuleCode string `gorm:"column:ModuleCode"`
	ClientType int `gorm:"column:ClientType"`
	ParentFunCode string `gorm:"column:ParentFunCode"`
	ParentFunName string `gorm:"column:ParentFunName"`
	IsFunctionPage int `gorm:"column:IsFunctionPage"`
}