package mysqlModel

import (
	"time"
)

// EntDepartments [...]
type EntDepartments struct {
	AutoID         int64  `gorm:"primaryKey;column:AutoID;type:bigint(22);not null" json:"-"`
	EnterpriseID   string `gorm:"column:EnterpriseID;type:varchar(20)" json:"enterpriseId"`
	DepartmentID   string `gorm:"column:DepartmentID;type:varchar(20)" json:"departmentId"`
	DepartmentName string `gorm:"column:DepartmentName;type:varchar(255)" json:"departmentName"`
	DepartmentCode int    `gorm:"column:DepartmentCode;type:int(11)" json:"departmentCode"`
	HeadID         string `gorm:"column:HeadID;type:varchar(20)" json:"headId"` // 部门领导人ID
	IsDeleted      uint8  `gorm:"column:IsDeleted;type:tinyint(1) unsigned zerofill" json:"isDeleted"`
}

// TableName get sql table name.获取数据库表名
func (m *EntDepartments) TableName() string {
	return "ent_departments"
}

// EntDepartmentsColumns get sql column name.获取数据库列名
var EntDepartmentsColumns = struct {
	AutoID         string
	EnterpriseID   string
	DepartmentID   string
	DepartmentName string
	DepartmentCode string
	HeadID         string
	IsDeleted      string
}{
	AutoID:         "AutoID",
	EnterpriseID:   "EnterpriseID",
	DepartmentID:   "DepartmentID",
	DepartmentName: "DepartmentName",
	DepartmentCode: "DepartmentCode",
	HeadID:         "HeadID",
	IsDeleted:      "IsDeleted",
}

// EntFile    档案基本信息表
type EntFile struct {
	AutoID       int64     `gorm:"primaryKey;column:AutoID;type:bigint(22);not null" json:"-"`
	FileID       string    `gorm:"column:FileID;type:varchar(20)" json:"fileId"`             //  文件（档案）ID
	CategoryID   int       `gorm:"column:CategoryID;type:int(11)" json:"categoryId"`         //  档案类别
	ProjectID    string    `gorm:"column:ProjectID;type:varchar(20)" json:"projectId"`       // 文件所属项目
	EnterpriseID string    `gorm:"column:EnterpriseID;type:varchar(20)" json:"enterpriseId"` // 文件所属公司ID
	FileName     string    `gorm:"column:FileName;type:varchar(255)" json:"fileName"`        // 文件名
	FileURL      string    `gorm:"column:FileURL;type:varchar(255)" json:"fileUrl"`          // 文件存放地址
	FileTypeID   int       `gorm:"column:FileTypeID;type:int(11)" json:"fileTypeId"`         // 文件类型
	Status       int8      `gorm:"column:Status;type:tinyint(11)" json:"status"`             // 0表示没有被借出，1表示已经借出
	UserID       string    `gorm:"column:UserID;type:varchar(20)" json:"userId"`             // 文件上传人
	FileCabinet  string    `gorm:"column:FileCabinet;type:varchar(30)" json:"fileCabinet"`   //  存放档案柜编号
	IsDelete     bool      `gorm:"column:IsDelete;type:tinyint(1)" json:"isDelete"`          // 是否删除
	UpdateTime   time.Time `gorm:"column:UpdateTime;type:timestamp" json:"updateTime"`       // 文件上传时间
	CreatTime    time.Time `gorm:"column:CreatTime;type:timestamp" json:"creatTime"`         // 记录创建时间
}

// TableName get sql table name.获取数据库表名
func (m *EntFile) TableName() string {
	return "ent_file"
}

// EntFileColumns get sql column name.获取数据库列名
var EntFileColumns = struct {
	AutoID       string
	FileID       string
	CategoryID   string
	ProjectID    string
	EnterpriseID string
	FileName     string
	FileURL      string
	FileTypeID   string
	Status       string
	UserID       string
	FileCabinet  string
	IsDelete     string
	UpdateTime   string
	CreatTime    string
}{
	AutoID:       "AutoID",
	FileID:       "FileID",
	CategoryID:   "CategoryID",
	ProjectID:    "ProjectID",
	EnterpriseID: "EnterpriseID",
	FileName:     "FileName",
	FileURL:      "FileURL",
	FileTypeID:   "FileTypeID",
	Status:       "Status",
	UserID:       "UserID",
	FileCabinet:  "FileCabinet",
	IsDelete:     "IsDelete",
	UpdateTime:   "UpdateTime",
	CreatTime:    "CreatTime",
}

// EntFileCategory   档案类别表：合同，图纸,,,,,
type EntFileCategory struct {
	AutoID           int64     `gorm:"primaryKey;column:AutoID;type:bigint(22);not null" json:"-"`
	CategoryID       int       `gorm:"column:CategoryID;type:int(11)" json:"categoryId"`             // 文件种类ID
	CategoryParentID int       `gorm:"column:CategoryParentID;type:int(11)" json:"categoryParentId"` //  父级ID
	ProjectID        string    `gorm:"column:ProjectID;type:varchar(20)" json:"projectId"`           // 所属项目ID
	EnterpriseID     string    `gorm:"column:EnterpriseID;type:varchar(20)" json:"enterpriseId"`     // 用户所属企业ID
	CategoryName     string    `gorm:"column:CategoryName;type:varchar(50)" json:"categoryName"`     // 类别名称:  基建施工图纸，地建施工图纸
	IsDelete         bool      `gorm:"column:IsDelete;type:tinyint(1)" json:"isDelete"`              // 是否删除
	CreatTime        time.Time `gorm:"column:CreatTime;type:timestamp" json:"creatTime"`             // 记录创建时间
}

// TableName get sql table name.获取数据库表名
func (m *EntFileCategory) TableName() string {
	return "ent_file_category"
}

// EntFileCategoryColumns get sql column name.获取数据库列名
var EntFileCategoryColumns = struct {
	AutoID           string
	CategoryID       string
	CategoryParentID string
	ProjectID        string
	EnterpriseID     string
	CategoryName     string
	IsDelete         string
	CreatTime        string
}{
	AutoID:           "AutoID",
	CategoryID:       "CategoryID",
	CategoryParentID: "CategoryParentID",
	ProjectID:        "ProjectID",
	EnterpriseID:     "EnterpriseID",
	CategoryName:     "CategoryName",
	IsDelete:         "IsDelete",
	CreatTime:        "CreatTime",
}

// EntFileLend      档案借阅信息表
type EntFileLend struct {
	AutoID       int64     `gorm:"primaryKey;column:AutoID;type:bigint(22);not null" json:"-"`
	FileID       string    `gorm:"column:FileID;type:varchar(20)" json:"fileId"`             //  文件（档案）ID
	EnterpriseID string    `gorm:"column:EnterpriseID;type:varchar(20)" json:"enterpriseId"` // 文件所属公司ID
	BorrowerID   string    `gorm:"column:BorrowerID;type:varchar(20)" json:"borrowerId"`     // 文件借阅人
	BorrowTime   time.Time `gorm:"column:BorrowTime;type:timestamp" json:"borrowTime"`       // 借出时间
	BorrowTerm   int8      `gorm:"column:BorrowTerm;type:tinyint(2)" json:"borrowTerm"`      //  借阅周期
	ReturnTime   time.Time `gorm:"column:ReturnTime;type:timestamp" json:"returnTime"`       // 归还时间
	IsDelete     bool      `gorm:"column:IsDelete;type:tinyint(1)" json:"isDelete"`          // 是否删除
	CreatTime    time.Time `gorm:"column:CreatTime;type:timestamp" json:"creatTime"`         // 记录创建时间
}

// TableName get sql table name.获取数据库表名
func (m *EntFileLend) TableName() string {
	return "ent_file_lend"
}

// EntFileLendColumns get sql column name.获取数据库列名
var EntFileLendColumns = struct {
	AutoID       string
	FileID       string
	EnterpriseID string
	BorrowerID   string
	BorrowTime   string
	BorrowTerm   string
	ReturnTime   string
	IsDelete     string
	CreatTime    string
}{
	AutoID:       "AutoID",
	FileID:       "FileID",
	EnterpriseID: "EnterpriseID",
	BorrowerID:   "BorrowerID",
	BorrowTime:   "BorrowTime",
	BorrowTerm:   "BorrowTerm",
	ReturnTime:   "ReturnTime",
	IsDelete:     "IsDelete",
	CreatTime:    "CreatTime",
}

// EntFileType    档案基本信息表
type EntFileType struct {
	AutoID       int64  `gorm:"primaryKey;column:AutoID;type:bigint(22);not null" json:"-"`
	FileTypeID   int    `gorm:"column:FileTypeID;type:int(11)" json:"fileTypeId"`         // 文件(格式)类型
	FileTypeName string `gorm:"column:FileTypeName;type:varchar(50)" json:"fileTypeName"` // 文件格式名称：DOC，PDF
}

// TableName get sql table name.获取数据库表名
func (m *EntFileType) TableName() string {
	return "ent_file_type"
}

// EntFileTypeColumns get sql column name.获取数据库列名
var EntFileTypeColumns = struct {
	AutoID       string
	FileTypeID   string
	FileTypeName string
}{
	AutoID:       "AutoID",
	FileTypeID:   "FileTypeID",
	FileTypeName: "FileTypeName",
}

// EntProject    档案基本信息表
type EntProject struct {
	AutoID       int64     `gorm:"primaryKey;column:AutoID;type:bigint(22);not null" json:"-"`
	ProjectID    string    `gorm:"column:ProjectID;type:varchar(20)" json:"projectId"`       //  文件（档案）ID
	EnterpriseID string    `gorm:"column:EnterpriseID;type:varchar(20)" json:"enterpriseId"` // 文件所属公司ID
	ProjectName  string    `gorm:"column:ProjectName;type:varchar(255)" json:"projectName"`  // 文件名
	IsDelete     bool      `gorm:"column:IsDelete;type:tinyint(1)" json:"isDelete"`          // 是否删除
	UpdateTime   time.Time `gorm:"column:UpdateTime;type:timestamp" json:"updateTime"`       // 文件上传时间
	CreatTime    time.Time `gorm:"column:CreatTime;type:timestamp" json:"creatTime"`         // 记录创建时间
}

// TableName get sql table name.获取数据库表名
func (m *EntProject) TableName() string {
	return "ent_project"
}

// EntProjectColumns get sql column name.获取数据库列名
var EntProjectColumns = struct {
	AutoID       string
	ProjectID    string
	EnterpriseID string
	ProjectName  string
	IsDelete     string
	UpdateTime   string
	CreatTime    string
}{
	AutoID:       "AutoID",
	ProjectID:    "ProjectID",
	EnterpriseID: "EnterpriseID",
	ProjectName:  "ProjectName",
	IsDelete:     "IsDelete",
	UpdateTime:   "UpdateTime",
	CreatTime:    "CreatTime",
}

// EntUsers [...]
type EntUsers struct {
	AutoID       int64     `gorm:"primaryKey;column:AutoID;type:bigint(22);not null" json:"-"`
	UserID       string    `gorm:"column:UserID;type:varchar(20)" json:"userId"`
	EnterpriseID string    `gorm:"column:EnterpriseID;type:varchar(20)" json:"enterpriseId"` // 用户所属企业ID
	Account      string    `gorm:"column:Account;type:varchar(255)" json:"account"`          //  账号（默认手机号）
	Password     string    `gorm:"column:Password;type:varchar(255)" json:"password"`
	UserName     string    `gorm:"column:UserName;type:varchar(255)" json:"userName"` // 用户姓名
	Phone        string    `gorm:"column:Phone;type:varchar(255)" json:"phone"`
	Email        string    `gorm:"column:Email;type:varchar(255)" json:"email"`
	FacePicURL   string    `gorm:"column:FacePicUrl;type:varchar(255)" json:"facePicUrl"` // 头像
	IsAdmin      bool      `gorm:"column:IsAdmin;type:tinyint(1);default:0" json:"isAdmin"`
	IsDeleted    bool      `gorm:"column:IsDeleted;type:tinyint(1);default:0" json:"isDeleted"` // 是否已删除：0--未删除；1--已经删除
	CreateTime   time.Time `gorm:"column:CreateTime;type:timestamp;default:CURRENT_TIMESTAMP" json:"createTime"`
	Token        string    `gorm:"column:Token;type:varchar(255)" json:"token"`
}

// TableName get sql table name.获取数据库表名
func (m *EntUsers) TableName() string {
	return "ent_users"
}

// EntUsersColumns get sql column name.获取数据库列名
var EntUsersColumns = struct {
	AutoID       string
	UserID       string
	EnterpriseID string
	Account      string
	Password     string
	UserName     string
	Phone        string
	Email        string
	FacePicURL   string
	IsAdmin      string
	IsDeleted    string
	CreateTime   string
	Token        string
}{
	AutoID:       "AutoID",
	UserID:       "UserID",
	EnterpriseID: "EnterpriseID",
	Account:      "Account",
	Password:     "Password",
	UserName:     "UserName",
	Phone:        "Phone",
	Email:        "Email",
	FacePicURL:   "FacePicUrl",
	IsAdmin:      "IsAdmin",
	IsDeleted:    "IsDeleted",
	CreateTime:   "CreateTime",
	Token:        "Token",
}

// PlatEnterprises [...]
type PlatEnterprises struct {
	AutoID             int64  `gorm:"primaryKey;column:AutoID;type:bigint(22);not null" json:"-"`
	EnterpriseID       string `gorm:"column:EnterpriseID;type:varchar(20);not null" json:"enterpriseId"`
	EnterpriseName     string `gorm:"column:EnterpriseName;type:varchar(255)" json:"enterpriseName"`
	EnterprisePassword string `gorm:"column:EnterprisePassword;type:varchar(255)" json:"enterprisePassword"` // 企业登录密码
	AdminID            string `gorm:"column:AdminID;type:varchar(20)" json:"adminId"`                        // 企业管理员ID
	Location           string `gorm:"column:Location;type:varchar(255)" json:"location"`                     // 企业地址
	EnterprisePhone    string `gorm:"column:EnterprisePhone;type:varchar(255)" json:"enterprisePhone"`       // 企业电话
	EnterpriseURL      string `gorm:"column:EnterpriseUrl;type:varchar(255)" json:"enterpriseUrl"`           // 企业URL
	LogoPicURL         string `gorm:"column:LogoPicUrl;type:varchar(255)" json:"logoPicUrl"`                 // 企业logo地址
}

// TableName get sql table name.获取数据库表名
func (m *PlatEnterprises) TableName() string {
	return "plat_enterprises"
}

// PlatEnterprisesColumns get sql column name.获取数据库列名
var PlatEnterprisesColumns = struct {
	AutoID             string
	EnterpriseID       string
	EnterpriseName     string
	EnterprisePassword string
	AdminID            string
	Location           string
	EnterprisePhone    string
	EnterpriseURL      string
	LogoPicURL         string
}{
	AutoID:             "AutoID",
	EnterpriseID:       "EnterpriseID",
	EnterpriseName:     "EnterpriseName",
	EnterprisePassword: "EnterprisePassword",
	AdminID:            "AdminID",
	Location:           "Location",
	EnterprisePhone:    "EnterprisePhone",
	EnterpriseURL:      "EnterpriseUrl",
	LogoPicURL:         "LogoPicUrl",
}

// PlatUsers [...]
type PlatUsers struct {
	AutoID     int64     `gorm:"primaryKey;column:AutoID;type:bigint(22);not null" json:"-"`
	UserID     string    `gorm:"column:UserID;type:varchar(20)" json:"userId" form:"userId"`
	UserName   string    `gorm:"column:UserName;type:varchar(255)" json:"userName" form:"userName"`
	Account    string    `gorm:"column:Account;type:varchar(255)" json:"account" form:"account"`
	Password   string    `gorm:"column:Password;type:varchar(255)" json:"password" form:"password"`
	Phone      string    `gorm:"column:Phone;type:varchar(255)" json:"phone" form:"phone"`
	Email      string    `gorm:"column:Email;type:varchar(255)" json:"email" form:"email"`
	IsDeleted  bool      `gorm:"column:IsDeleted;type:tinyint(1)" json:"isDeleted" form:"isDeleted"`
	CreateTime time.Time `gorm:"column:CreateTime;type:timestamp;default:CURRENT_TIMESTAMP" json:"createTime" form:"createTime"`
}

// TableName get sql table name.获取数据库表名
func (m *PlatUsers) TableName() string {
	return "plat_users"
}

// PlatUsersColumns get sql column name.获取数据库列名
var PlatUsersColumns = struct {
	AutoID     string
	UserID     string
	UserName   string
	Account    string
	Password   string
	Phone      string
	Email      string
	IsDeleted  string
	CreateTime string
}{
	AutoID:     "AutoID",
	UserID:     "UserID",
	UserName:   "UserName",
	Account:    "Account",
	Password:   "Password",
	Phone:      "Phone",
	Email:      "Email",
	IsDeleted:  "IsDeleted",
	CreateTime: "CreateTime",
}
