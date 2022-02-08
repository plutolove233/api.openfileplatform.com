package mysqlModel

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _PlatEnterprisesMgr struct {
	*_BaseMgr
}

// PlatEnterprisesMgr open func
func PlatEnterprisesMgr(db *gorm.DB) *_PlatEnterprisesMgr {
	if db == nil {
		panic(fmt.Errorf("PlatEnterprisesMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_PlatEnterprisesMgr{_BaseMgr: &_BaseMgr{DB: db.Table("plat_enterprises"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_PlatEnterprisesMgr) GetTableName() string {
	return "plat_enterprises"
}

// Reset 重置gorm会话
func (obj *_PlatEnterprisesMgr) Reset() *_PlatEnterprisesMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_PlatEnterprisesMgr) Get() (result PlatEnterprises, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PlatEnterprises{}).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_PlatEnterprisesMgr) Gets() (results []*PlatEnterprises, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PlatEnterprises{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_PlatEnterprisesMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(PlatEnterprises{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithAutoID AutoID获取
func (obj *_PlatEnterprisesMgr) WithAutoID(autoID int64) Option {
	return optionFunc(func(o *options) { o.query["AutoID"] = autoID })
}

// WithEnterpriseID EnterpriseID获取
func (obj *_PlatEnterprisesMgr) WithEnterpriseID(enterpriseID string) Option {
	return optionFunc(func(o *options) { o.query["EnterpriseID"] = enterpriseID })
}

// WithEnterpriseName EnterpriseName获取
func (obj *_PlatEnterprisesMgr) WithEnterpriseName(enterpriseName string) Option {
	return optionFunc(func(o *options) { o.query["EnterpriseName"] = enterpriseName })
}

// WithEnterprisePassword EnterprisePassword获取 企业登录密码
func (obj *_PlatEnterprisesMgr) WithEnterprisePassword(enterprisePassword string) Option {
	return optionFunc(func(o *options) { o.query["EnterprisePassword"] = enterprisePassword })
}

// WithAdminID AdminID获取 企业管理员ID
func (obj *_PlatEnterprisesMgr) WithAdminID(adminID string) Option {
	return optionFunc(func(o *options) { o.query["AdminID"] = adminID })
}

// WithLocation Location获取 企业地址
func (obj *_PlatEnterprisesMgr) WithLocation(location string) Option {
	return optionFunc(func(o *options) { o.query["Location"] = location })
}

// WithEnterprisePhone EnterprisePhone获取 企业电话
func (obj *_PlatEnterprisesMgr) WithEnterprisePhone(enterprisePhone string) Option {
	return optionFunc(func(o *options) { o.query["EnterprisePhone"] = enterprisePhone })
}

// WithEnterpriseURL EnterpriseUrl获取 企业URL
func (obj *_PlatEnterprisesMgr) WithEnterpriseURL(enterpriseURL string) Option {
	return optionFunc(func(o *options) { o.query["EnterpriseUrl"] = enterpriseURL })
}

// WithLogoPicURL LogoPicUrl获取 企业logo地址
func (obj *_PlatEnterprisesMgr) WithLogoPicURL(logoPicURL string) Option {
	return optionFunc(func(o *options) { o.query["LogoPicUrl"] = logoPicURL })
}

// GetByOption 功能选项模式获取
func (obj *_PlatEnterprisesMgr) GetByOption(opts ...Option) (result PlatEnterprises, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(PlatEnterprises{}).Where(options.query).Find(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_PlatEnterprisesMgr) GetByOptions(opts ...Option) (results []*PlatEnterprises, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(PlatEnterprises{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromAutoID 通过AutoID获取内容
func (obj *_PlatEnterprisesMgr) GetFromAutoID(autoID int64) (result PlatEnterprises, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PlatEnterprises{}).Where("`AutoID` = ?", autoID).Find(&result).Error

	return
}

// GetBatchFromAutoID 批量查找
func (obj *_PlatEnterprisesMgr) GetBatchFromAutoID(autoIDs []int64) (results []*PlatEnterprises, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PlatEnterprises{}).Where("`AutoID` IN (?)", autoIDs).Find(&results).Error

	return
}

// GetFromEnterpriseID 通过EnterpriseID获取内容
func (obj *_PlatEnterprisesMgr) GetFromEnterpriseID(enterpriseID string) (results []*PlatEnterprises, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PlatEnterprises{}).Where("`EnterpriseID` = ?", enterpriseID).Find(&results).Error

	return
}

// GetBatchFromEnterpriseID 批量查找
func (obj *_PlatEnterprisesMgr) GetBatchFromEnterpriseID(enterpriseIDs []string) (results []*PlatEnterprises, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PlatEnterprises{}).Where("`EnterpriseID` IN (?)", enterpriseIDs).Find(&results).Error

	return
}

// GetFromEnterpriseName 通过EnterpriseName获取内容
func (obj *_PlatEnterprisesMgr) GetFromEnterpriseName(enterpriseName string) (results []*PlatEnterprises, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PlatEnterprises{}).Where("`EnterpriseName` = ?", enterpriseName).Find(&results).Error

	return
}

// GetBatchFromEnterpriseName 批量查找
func (obj *_PlatEnterprisesMgr) GetBatchFromEnterpriseName(enterpriseNames []string) (results []*PlatEnterprises, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PlatEnterprises{}).Where("`EnterpriseName` IN (?)", enterpriseNames).Find(&results).Error

	return
}

// GetFromEnterprisePassword 通过EnterprisePassword获取内容 企业登录密码
func (obj *_PlatEnterprisesMgr) GetFromEnterprisePassword(enterprisePassword string) (results []*PlatEnterprises, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PlatEnterprises{}).Where("`EnterprisePassword` = ?", enterprisePassword).Find(&results).Error

	return
}

// GetBatchFromEnterprisePassword 批量查找 企业登录密码
func (obj *_PlatEnterprisesMgr) GetBatchFromEnterprisePassword(enterprisePasswords []string) (results []*PlatEnterprises, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PlatEnterprises{}).Where("`EnterprisePassword` IN (?)", enterprisePasswords).Find(&results).Error

	return
}

// GetFromAdminID 通过AdminID获取内容 企业管理员ID
func (obj *_PlatEnterprisesMgr) GetFromAdminID(adminID string) (results []*PlatEnterprises, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PlatEnterprises{}).Where("`AdminID` = ?", adminID).Find(&results).Error

	return
}

// GetBatchFromAdminID 批量查找 企业管理员ID
func (obj *_PlatEnterprisesMgr) GetBatchFromAdminID(adminIDs []string) (results []*PlatEnterprises, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PlatEnterprises{}).Where("`AdminID` IN (?)", adminIDs).Find(&results).Error

	return
}

// GetFromLocation 通过Location获取内容 企业地址
func (obj *_PlatEnterprisesMgr) GetFromLocation(location string) (results []*PlatEnterprises, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PlatEnterprises{}).Where("`Location` = ?", location).Find(&results).Error

	return
}

// GetBatchFromLocation 批量查找 企业地址
func (obj *_PlatEnterprisesMgr) GetBatchFromLocation(locations []string) (results []*PlatEnterprises, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PlatEnterprises{}).Where("`Location` IN (?)", locations).Find(&results).Error

	return
}

// GetFromEnterprisePhone 通过EnterprisePhone获取内容 企业电话
func (obj *_PlatEnterprisesMgr) GetFromEnterprisePhone(enterprisePhone string) (results []*PlatEnterprises, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PlatEnterprises{}).Where("`EnterprisePhone` = ?", enterprisePhone).Find(&results).Error

	return
}

// GetBatchFromEnterprisePhone 批量查找 企业电话
func (obj *_PlatEnterprisesMgr) GetBatchFromEnterprisePhone(enterprisePhones []string) (results []*PlatEnterprises, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PlatEnterprises{}).Where("`EnterprisePhone` IN (?)", enterprisePhones).Find(&results).Error

	return
}

// GetFromEnterpriseURL 通过EnterpriseUrl获取内容 企业URL
func (obj *_PlatEnterprisesMgr) GetFromEnterpriseURL(enterpriseURL string) (results []*PlatEnterprises, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PlatEnterprises{}).Where("`EnterpriseUrl` = ?", enterpriseURL).Find(&results).Error

	return
}

// GetBatchFromEnterpriseURL 批量查找 企业URL
func (obj *_PlatEnterprisesMgr) GetBatchFromEnterpriseURL(enterpriseURLs []string) (results []*PlatEnterprises, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PlatEnterprises{}).Where("`EnterpriseUrl` IN (?)", enterpriseURLs).Find(&results).Error

	return
}

// GetFromLogoPicURL 通过LogoPicUrl获取内容 企业logo地址
func (obj *_PlatEnterprisesMgr) GetFromLogoPicURL(logoPicURL string) (results []*PlatEnterprises, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PlatEnterprises{}).Where("`LogoPicUrl` = ?", logoPicURL).Find(&results).Error

	return
}

// GetBatchFromLogoPicURL 批量查找 企业logo地址
func (obj *_PlatEnterprisesMgr) GetBatchFromLogoPicURL(logoPicURLs []string) (results []*PlatEnterprises, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PlatEnterprises{}).Where("`LogoPicUrl` IN (?)", logoPicURLs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_PlatEnterprisesMgr) FetchByPrimaryKey(autoID int64) (result PlatEnterprises, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PlatEnterprises{}).Where("`AutoID` = ?", autoID).Find(&result).Error

	return
}
