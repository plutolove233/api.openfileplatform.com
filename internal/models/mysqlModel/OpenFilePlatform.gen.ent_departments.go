package mysqlModel

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _EntDepartmentsMgr struct {
	*_BaseMgr
}

// EntDepartmentsMgr open func
func EntDepartmentsMgr(db *gorm.DB) *_EntDepartmentsMgr {
	if db == nil {
		panic(fmt.Errorf("EntDepartmentsMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EntDepartmentsMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ent_departments"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EntDepartmentsMgr) GetTableName() string {
	return "ent_departments"
}

// Reset 重置gorm会话
func (obj *_EntDepartmentsMgr) Reset() *_EntDepartmentsMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_EntDepartmentsMgr) Get() (result EntDepartments, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EntDepartments{}).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_EntDepartmentsMgr) Gets() (results []*EntDepartments, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EntDepartments{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_EntDepartmentsMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(EntDepartments{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithAutoID AutoID获取
func (obj *_EntDepartmentsMgr) WithAutoID(autoID int64) Option {
	return optionFunc(func(o *options) { o.query["AutoID"] = autoID })
}

// WithEnterpriseID EnterpriseID获取
func (obj *_EntDepartmentsMgr) WithEnterpriseID(enterpriseID string) Option {
	return optionFunc(func(o *options) { o.query["EnterpriseID"] = enterpriseID })
}

// WithDepartmentID DepartmentID获取
func (obj *_EntDepartmentsMgr) WithDepartmentID(departmentID string) Option {
	return optionFunc(func(o *options) { o.query["DepartmentID"] = departmentID })
}

// WithDepartmentName DepartmentName获取
func (obj *_EntDepartmentsMgr) WithDepartmentName(departmentName string) Option {
	return optionFunc(func(o *options) { o.query["DepartmentName"] = departmentName })
}

// WithDepartmentCode DepartmentCode获取
func (obj *_EntDepartmentsMgr) WithDepartmentCode(departmentCode int) Option {
	return optionFunc(func(o *options) { o.query["DepartmentCode"] = departmentCode })
}

// WithHeadID HeadID获取 部门领导人ID
func (obj *_EntDepartmentsMgr) WithHeadID(headID string) Option {
	return optionFunc(func(o *options) { o.query["HeadID"] = headID })
}

// WithIsDeleted IsDeleted获取
func (obj *_EntDepartmentsMgr) WithIsDeleted(isDeleted uint8) Option {
	return optionFunc(func(o *options) { o.query["IsDeleted"] = isDeleted })
}

// GetByOption 功能选项模式获取
func (obj *_EntDepartmentsMgr) GetByOption(opts ...Option) (result EntDepartments, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(EntDepartments{}).Where(options.query).Find(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_EntDepartmentsMgr) GetByOptions(opts ...Option) (results []*EntDepartments, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(EntDepartments{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromAutoID 通过AutoID获取内容
func (obj *_EntDepartmentsMgr) GetFromAutoID(autoID int64) (result EntDepartments, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EntDepartments{}).Where("`AutoID` = ?", autoID).Find(&result).Error

	return
}

// GetBatchFromAutoID 批量查找
func (obj *_EntDepartmentsMgr) GetBatchFromAutoID(autoIDs []int64) (results []*EntDepartments, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EntDepartments{}).Where("`AutoID` IN (?)", autoIDs).Find(&results).Error

	return
}

// GetFromEnterpriseID 通过EnterpriseID获取内容
func (obj *_EntDepartmentsMgr) GetFromEnterpriseID(enterpriseID string) (results []*EntDepartments, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EntDepartments{}).Where("`EnterpriseID` = ?", enterpriseID).Find(&results).Error

	return
}

// GetBatchFromEnterpriseID 批量查找
func (obj *_EntDepartmentsMgr) GetBatchFromEnterpriseID(enterpriseIDs []string) (results []*EntDepartments, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EntDepartments{}).Where("`EnterpriseID` IN (?)", enterpriseIDs).Find(&results).Error

	return
}

// GetFromDepartmentID 通过DepartmentID获取内容
func (obj *_EntDepartmentsMgr) GetFromDepartmentID(departmentID string) (results []*EntDepartments, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EntDepartments{}).Where("`DepartmentID` = ?", departmentID).Find(&results).Error

	return
}

// GetBatchFromDepartmentID 批量查找
func (obj *_EntDepartmentsMgr) GetBatchFromDepartmentID(departmentIDs []string) (results []*EntDepartments, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EntDepartments{}).Where("`DepartmentID` IN (?)", departmentIDs).Find(&results).Error

	return
}

// GetFromDepartmentName 通过DepartmentName获取内容
func (obj *_EntDepartmentsMgr) GetFromDepartmentName(departmentName string) (results []*EntDepartments, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EntDepartments{}).Where("`DepartmentName` = ?", departmentName).Find(&results).Error

	return
}

// GetBatchFromDepartmentName 批量查找
func (obj *_EntDepartmentsMgr) GetBatchFromDepartmentName(departmentNames []string) (results []*EntDepartments, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EntDepartments{}).Where("`DepartmentName` IN (?)", departmentNames).Find(&results).Error

	return
}

// GetFromDepartmentCode 通过DepartmentCode获取内容
func (obj *_EntDepartmentsMgr) GetFromDepartmentCode(departmentCode int) (results []*EntDepartments, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EntDepartments{}).Where("`DepartmentCode` = ?", departmentCode).Find(&results).Error

	return
}

// GetBatchFromDepartmentCode 批量查找
func (obj *_EntDepartmentsMgr) GetBatchFromDepartmentCode(departmentCodes []int) (results []*EntDepartments, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EntDepartments{}).Where("`DepartmentCode` IN (?)", departmentCodes).Find(&results).Error

	return
}

// GetFromHeadID 通过HeadID获取内容 部门领导人ID
func (obj *_EntDepartmentsMgr) GetFromHeadID(headID string) (results []*EntDepartments, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EntDepartments{}).Where("`HeadID` = ?", headID).Find(&results).Error

	return
}

// GetBatchFromHeadID 批量查找 部门领导人ID
func (obj *_EntDepartmentsMgr) GetBatchFromHeadID(headIDs []string) (results []*EntDepartments, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EntDepartments{}).Where("`HeadID` IN (?)", headIDs).Find(&results).Error

	return
}

// GetFromIsDeleted 通过IsDeleted获取内容
func (obj *_EntDepartmentsMgr) GetFromIsDeleted(isDeleted uint8) (results []*EntDepartments, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EntDepartments{}).Where("`IsDeleted` = ?", isDeleted).Find(&results).Error

	return
}

// GetBatchFromIsDeleted 批量查找
func (obj *_EntDepartmentsMgr) GetBatchFromIsDeleted(isDeleteds []uint8) (results []*EntDepartments, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EntDepartments{}).Where("`IsDeleted` IN (?)", isDeleteds).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_EntDepartmentsMgr) FetchByPrimaryKey(autoID int64) (result EntDepartments, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EntDepartments{}).Where("`AutoID` = ?", autoID).Find(&result).Error

	return
}
