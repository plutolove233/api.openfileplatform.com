package mysqlModel

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _EntProjectMgr struct {
	*_BaseMgr
}

// EntProjectMgr open func
func EntProjectMgr(db *gorm.DB) *_EntProjectMgr {
	if db == nil {
		panic(fmt.Errorf("EntProjectMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EntProjectMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ent_project"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EntProjectMgr) GetTableName() string {
	return "ent_project"
}

// Reset 重置gorm会话
func (obj *_EntProjectMgr) Reset() *_EntProjectMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_EntProjectMgr) Get() (result EntProject, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EntProject{}).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_EntProjectMgr) Gets() (results []*EntProject, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EntProject{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_EntProjectMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(EntProject{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithAutoID AutoID获取
func (obj *_EntProjectMgr) WithAutoID(autoID int64) Option {
	return optionFunc(func(o *options) { o.query["AutoID"] = autoID })
}

// WithProjectID ProjectID获取  文件（档案）ID
func (obj *_EntProjectMgr) WithProjectID(projectID string) Option {
	return optionFunc(func(o *options) { o.query["ProjectID"] = projectID })
}

// WithEnterpriseID EnterpriseID获取 文件所属公司ID
func (obj *_EntProjectMgr) WithEnterpriseID(enterpriseID string) Option {
	return optionFunc(func(o *options) { o.query["EnterpriseID"] = enterpriseID })
}

// WithProjectName ProjectName获取 文件名
func (obj *_EntProjectMgr) WithProjectName(projectName string) Option {
	return optionFunc(func(o *options) { o.query["ProjectName"] = projectName })
}

// WithIsDelete IsDelete获取 是否删除
func (obj *_EntProjectMgr) WithIsDelete(isDelete bool) Option {
	return optionFunc(func(o *options) { o.query["IsDelete"] = isDelete })
}

// WithUpdateTime UpdateTime获取 文件上传时间
func (obj *_EntProjectMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["UpdateTime"] = updateTime })
}

// WithCreatTime CreatTime获取 记录创建时间
func (obj *_EntProjectMgr) WithCreatTime(creatTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["CreatTime"] = creatTime })
}

// GetByOption 功能选项模式获取
func (obj *_EntProjectMgr) GetByOption(opts ...Option) (result EntProject, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(EntProject{}).Where(options.query).Find(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_EntProjectMgr) GetByOptions(opts ...Option) (results []*EntProject, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(EntProject{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromAutoID 通过AutoID获取内容
func (obj *_EntProjectMgr) GetFromAutoID(autoID int64) (result EntProject, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EntProject{}).Where("`AutoID` = ?", autoID).Find(&result).Error

	return
}

// GetBatchFromAutoID 批量查找
func (obj *_EntProjectMgr) GetBatchFromAutoID(autoIDs []int64) (results []*EntProject, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EntProject{}).Where("`AutoID` IN (?)", autoIDs).Find(&results).Error

	return
}

// GetFromProjectID 通过ProjectID获取内容  文件（档案）ID
func (obj *_EntProjectMgr) GetFromProjectID(projectID string) (results []*EntProject, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EntProject{}).Where("`ProjectID` = ?", projectID).Find(&results).Error

	return
}

// GetBatchFromProjectID 批量查找  文件（档案）ID
func (obj *_EntProjectMgr) GetBatchFromProjectID(projectIDs []string) (results []*EntProject, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EntProject{}).Where("`ProjectID` IN (?)", projectIDs).Find(&results).Error

	return
}

// GetFromEnterpriseID 通过EnterpriseID获取内容 文件所属公司ID
func (obj *_EntProjectMgr) GetFromEnterpriseID(enterpriseID string) (results []*EntProject, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EntProject{}).Where("`EnterpriseID` = ?", enterpriseID).Find(&results).Error

	return
}

// GetBatchFromEnterpriseID 批量查找 文件所属公司ID
func (obj *_EntProjectMgr) GetBatchFromEnterpriseID(enterpriseIDs []string) (results []*EntProject, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EntProject{}).Where("`EnterpriseID` IN (?)", enterpriseIDs).Find(&results).Error

	return
}

// GetFromProjectName 通过ProjectName获取内容 文件名
func (obj *_EntProjectMgr) GetFromProjectName(projectName string) (results []*EntProject, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EntProject{}).Where("`ProjectName` = ?", projectName).Find(&results).Error

	return
}

// GetBatchFromProjectName 批量查找 文件名
func (obj *_EntProjectMgr) GetBatchFromProjectName(projectNames []string) (results []*EntProject, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EntProject{}).Where("`ProjectName` IN (?)", projectNames).Find(&results).Error

	return
}

// GetFromIsDelete 通过IsDelete获取内容 是否删除
func (obj *_EntProjectMgr) GetFromIsDelete(isDelete bool) (results []*EntProject, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EntProject{}).Where("`IsDelete` = ?", isDelete).Find(&results).Error

	return
}

// GetBatchFromIsDelete 批量查找 是否删除
func (obj *_EntProjectMgr) GetBatchFromIsDelete(isDeletes []bool) (results []*EntProject, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EntProject{}).Where("`IsDelete` IN (?)", isDeletes).Find(&results).Error

	return
}

// GetFromUpdateTime 通过UpdateTime获取内容 文件上传时间
func (obj *_EntProjectMgr) GetFromUpdateTime(updateTime time.Time) (results []*EntProject, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EntProject{}).Where("`UpdateTime` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找 文件上传时间
func (obj *_EntProjectMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*EntProject, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EntProject{}).Where("`UpdateTime` IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromCreatTime 通过CreatTime获取内容 记录创建时间
func (obj *_EntProjectMgr) GetFromCreatTime(creatTime time.Time) (results []*EntProject, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EntProject{}).Where("`CreatTime` = ?", creatTime).Find(&results).Error

	return
}

// GetBatchFromCreatTime 批量查找 记录创建时间
func (obj *_EntProjectMgr) GetBatchFromCreatTime(creatTimes []time.Time) (results []*EntProject, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EntProject{}).Where("`CreatTime` IN (?)", creatTimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_EntProjectMgr) FetchByPrimaryKey(autoID int64) (result EntProject, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EntProject{}).Where("`AutoID` = ?", autoID).Find(&result).Error

	return
}
