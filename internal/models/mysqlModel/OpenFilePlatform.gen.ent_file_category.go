package mysqlModel

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _EntFileCategoryMgr struct {
	*_BaseMgr
}

// EntFileCategoryMgr open func
func EntFileCategoryMgr(db *gorm.DB) *_EntFileCategoryMgr {
	if db == nil {
		panic(fmt.Errorf("EntFileCategoryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EntFileCategoryMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ent_file_category"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EntFileCategoryMgr) GetTableName() string {
	return "ent_file_category"
}

// Reset 重置gorm会话
func (obj *_EntFileCategoryMgr) Reset() *_EntFileCategoryMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_EntFileCategoryMgr) Get() (result EntFileCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EntFileCategory{}).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_EntFileCategoryMgr) Gets() (results []*EntFileCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EntFileCategory{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_EntFileCategoryMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(EntFileCategory{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithAutoID AutoID获取
func (obj *_EntFileCategoryMgr) WithAutoID(autoID int64) Option {
	return optionFunc(func(o *options) { o.query["AutoID"] = autoID })
}

// WithCategoryID CategoryID获取 文件种类ID
func (obj *_EntFileCategoryMgr) WithCategoryID(categoryID int) Option {
	return optionFunc(func(o *options) { o.query["CategoryID"] = categoryID })
}

// WithCategoryParentID CategoryParentID获取  父级ID
func (obj *_EntFileCategoryMgr) WithCategoryParentID(categoryParentID int) Option {
	return optionFunc(func(o *options) { o.query["CategoryParentID"] = categoryParentID })
}

// WithProjectID ProjectID获取 所属项目ID
func (obj *_EntFileCategoryMgr) WithProjectID(projectID string) Option {
	return optionFunc(func(o *options) { o.query["ProjectID"] = projectID })
}

// WithEnterpriseID EnterpriseID获取 用户所属企业ID
func (obj *_EntFileCategoryMgr) WithEnterpriseID(enterpriseID string) Option {
	return optionFunc(func(o *options) { o.query["EnterpriseID"] = enterpriseID })
}

// WithCategoryName CategoryName获取 类别名称:  基建施工图纸，地建施工图纸
func (obj *_EntFileCategoryMgr) WithCategoryName(categoryName string) Option {
	return optionFunc(func(o *options) { o.query["CategoryName"] = categoryName })
}

// WithIsDelete IsDelete获取 是否删除
func (obj *_EntFileCategoryMgr) WithIsDelete(isDelete bool) Option {
	return optionFunc(func(o *options) { o.query["IsDelete"] = isDelete })
}

// WithCreatTime CreatTime获取 记录创建时间
func (obj *_EntFileCategoryMgr) WithCreatTime(creatTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["CreatTime"] = creatTime })
}

// GetByOption 功能选项模式获取
func (obj *_EntFileCategoryMgr) GetByOption(opts ...Option) (result EntFileCategory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(EntFileCategory{}).Where(options.query).Find(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_EntFileCategoryMgr) GetByOptions(opts ...Option) (results []*EntFileCategory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(EntFileCategory{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromAutoID 通过AutoID获取内容
func (obj *_EntFileCategoryMgr) GetFromAutoID(autoID int64) (result EntFileCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EntFileCategory{}).Where("`AutoID` = ?", autoID).Find(&result).Error

	return
}

// GetBatchFromAutoID 批量查找
func (obj *_EntFileCategoryMgr) GetBatchFromAutoID(autoIDs []int64) (results []*EntFileCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EntFileCategory{}).Where("`AutoID` IN (?)", autoIDs).Find(&results).Error

	return
}

// GetFromCategoryID 通过CategoryID获取内容 文件种类ID
func (obj *_EntFileCategoryMgr) GetFromCategoryID(categoryID int) (results []*EntFileCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EntFileCategory{}).Where("`CategoryID` = ?", categoryID).Find(&results).Error

	return
}

// GetBatchFromCategoryID 批量查找 文件种类ID
func (obj *_EntFileCategoryMgr) GetBatchFromCategoryID(categoryIDs []int) (results []*EntFileCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EntFileCategory{}).Where("`CategoryID` IN (?)", categoryIDs).Find(&results).Error

	return
}

// GetFromCategoryParentID 通过CategoryParentID获取内容  父级ID
func (obj *_EntFileCategoryMgr) GetFromCategoryParentID(categoryParentID int) (results []*EntFileCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EntFileCategory{}).Where("`CategoryParentID` = ?", categoryParentID).Find(&results).Error

	return
}

// GetBatchFromCategoryParentID 批量查找  父级ID
func (obj *_EntFileCategoryMgr) GetBatchFromCategoryParentID(categoryParentIDs []int) (results []*EntFileCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EntFileCategory{}).Where("`CategoryParentID` IN (?)", categoryParentIDs).Find(&results).Error

	return
}

// GetFromProjectID 通过ProjectID获取内容 所属项目ID
func (obj *_EntFileCategoryMgr) GetFromProjectID(projectID string) (results []*EntFileCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EntFileCategory{}).Where("`ProjectID` = ?", projectID).Find(&results).Error

	return
}

// GetBatchFromProjectID 批量查找 所属项目ID
func (obj *_EntFileCategoryMgr) GetBatchFromProjectID(projectIDs []string) (results []*EntFileCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EntFileCategory{}).Where("`ProjectID` IN (?)", projectIDs).Find(&results).Error

	return
}

// GetFromEnterpriseID 通过EnterpriseID获取内容 用户所属企业ID
func (obj *_EntFileCategoryMgr) GetFromEnterpriseID(enterpriseID string) (results []*EntFileCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EntFileCategory{}).Where("`EnterpriseID` = ?", enterpriseID).Find(&results).Error

	return
}

// GetBatchFromEnterpriseID 批量查找 用户所属企业ID
func (obj *_EntFileCategoryMgr) GetBatchFromEnterpriseID(enterpriseIDs []string) (results []*EntFileCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EntFileCategory{}).Where("`EnterpriseID` IN (?)", enterpriseIDs).Find(&results).Error

	return
}

// GetFromCategoryName 通过CategoryName获取内容 类别名称:  基建施工图纸，地建施工图纸
func (obj *_EntFileCategoryMgr) GetFromCategoryName(categoryName string) (results []*EntFileCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EntFileCategory{}).Where("`CategoryName` = ?", categoryName).Find(&results).Error

	return
}

// GetBatchFromCategoryName 批量查找 类别名称:  基建施工图纸，地建施工图纸
func (obj *_EntFileCategoryMgr) GetBatchFromCategoryName(categoryNames []string) (results []*EntFileCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EntFileCategory{}).Where("`CategoryName` IN (?)", categoryNames).Find(&results).Error

	return
}

// GetFromIsDelete 通过IsDelete获取内容 是否删除
func (obj *_EntFileCategoryMgr) GetFromIsDelete(isDelete bool) (results []*EntFileCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EntFileCategory{}).Where("`IsDelete` = ?", isDelete).Find(&results).Error

	return
}

// GetBatchFromIsDelete 批量查找 是否删除
func (obj *_EntFileCategoryMgr) GetBatchFromIsDelete(isDeletes []bool) (results []*EntFileCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EntFileCategory{}).Where("`IsDelete` IN (?)", isDeletes).Find(&results).Error

	return
}

// GetFromCreatTime 通过CreatTime获取内容 记录创建时间
func (obj *_EntFileCategoryMgr) GetFromCreatTime(creatTime time.Time) (results []*EntFileCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EntFileCategory{}).Where("`CreatTime` = ?", creatTime).Find(&results).Error

	return
}

// GetBatchFromCreatTime 批量查找 记录创建时间
func (obj *_EntFileCategoryMgr) GetBatchFromCreatTime(creatTimes []time.Time) (results []*EntFileCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EntFileCategory{}).Where("`CreatTime` IN (?)", creatTimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_EntFileCategoryMgr) FetchByPrimaryKey(autoID int64) (result EntFileCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EntFileCategory{}).Where("`AutoID` = ?", autoID).Find(&result).Error

	return
}
