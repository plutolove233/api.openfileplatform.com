package mysqlModel

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _EntFileMgr struct {
	*_BaseMgr
}

// EntFileMgr open func
func EntFileMgr(db *gorm.DB) *_EntFileMgr {
	if db == nil {
		panic(fmt.Errorf("EntFileMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EntFileMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ent_file"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EntFileMgr) GetTableName() string {
	return "ent_file"
}

// Reset 重置gorm会话
func (obj *_EntFileMgr) Reset() *_EntFileMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_EntFileMgr) Get() (result EntFile, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EntFile{}).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_EntFileMgr) Gets() (results []*EntFile, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EntFile{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_EntFileMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(EntFile{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithAutoID AutoID获取
func (obj *_EntFileMgr) WithAutoID(autoID int64) Option {
	return optionFunc(func(o *options) { o.query["AutoID"] = autoID })
}

// WithFileID FileID获取  文件（档案）ID
func (obj *_EntFileMgr) WithFileID(fileID string) Option {
	return optionFunc(func(o *options) { o.query["FileID"] = fileID })
}

// WithCategoryID CategoryID获取  档案类别
func (obj *_EntFileMgr) WithCategoryID(categoryID int) Option {
	return optionFunc(func(o *options) { o.query["CategoryID"] = categoryID })
}

// WithProjectID ProjectID获取 文件所属项目
func (obj *_EntFileMgr) WithProjectID(projectID string) Option {
	return optionFunc(func(o *options) { o.query["ProjectID"] = projectID })
}

// WithEnterpriseID EnterpriseID获取 文件所属公司ID
func (obj *_EntFileMgr) WithEnterpriseID(enterpriseID string) Option {
	return optionFunc(func(o *options) { o.query["EnterpriseID"] = enterpriseID })
}

// WithFileName FileName获取 文件名
func (obj *_EntFileMgr) WithFileName(fileName string) Option {
	return optionFunc(func(o *options) { o.query["FileName"] = fileName })
}

// WithFileURL FileURL获取 文件存放地址
func (obj *_EntFileMgr) WithFileURL(fileURL string) Option {
	return optionFunc(func(o *options) { o.query["FileURL"] = fileURL })
}

// WithFileTypeID FileTypeID获取 文件类型
func (obj *_EntFileMgr) WithFileTypeID(fileTypeID int) Option {
	return optionFunc(func(o *options) { o.query["FileTypeID"] = fileTypeID })
}

// WithStatus Status获取 0表示没有被借出，1表示已经借出
func (obj *_EntFileMgr) WithStatus(status int8) Option {
	return optionFunc(func(o *options) { o.query["Status"] = status })
}

// WithUserID UserID获取 文件上传人
func (obj *_EntFileMgr) WithUserID(userID string) Option {
	return optionFunc(func(o *options) { o.query["UserID"] = userID })
}

// WithFileCabinet FileCabinet获取  存放档案柜编号
func (obj *_EntFileMgr) WithFileCabinet(fileCabinet string) Option {
	return optionFunc(func(o *options) { o.query["FileCabinet"] = fileCabinet })
}

// WithIsDelete IsDelete获取 是否删除
func (obj *_EntFileMgr) WithIsDelete(isDelete bool) Option {
	return optionFunc(func(o *options) { o.query["IsDelete"] = isDelete })
}

// WithUpdateTime UpdateTime获取 文件上传时间
func (obj *_EntFileMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["UpdateTime"] = updateTime })
}

// WithCreatTime CreatTime获取 记录创建时间
func (obj *_EntFileMgr) WithCreatTime(creatTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["CreatTime"] = creatTime })
}

// GetByOption 功能选项模式获取
func (obj *_EntFileMgr) GetByOption(opts ...Option) (result EntFile, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(EntFile{}).Where(options.query).Find(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_EntFileMgr) GetByOptions(opts ...Option) (results []*EntFile, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(EntFile{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromAutoID 通过AutoID获取内容
func (obj *_EntFileMgr) GetFromAutoID(autoID int64) (result EntFile, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EntFile{}).Where("`AutoID` = ?", autoID).Find(&result).Error

	return
}

// GetBatchFromAutoID 批量查找
func (obj *_EntFileMgr) GetBatchFromAutoID(autoIDs []int64) (results []*EntFile, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EntFile{}).Where("`AutoID` IN (?)", autoIDs).Find(&results).Error

	return
}

// GetFromFileID 通过FileID获取内容  文件（档案）ID
func (obj *_EntFileMgr) GetFromFileID(fileID string) (results []*EntFile, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EntFile{}).Where("`FileID` = ?", fileID).Find(&results).Error

	return
}

// GetBatchFromFileID 批量查找  文件（档案）ID
func (obj *_EntFileMgr) GetBatchFromFileID(fileIDs []string) (results []*EntFile, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EntFile{}).Where("`FileID` IN (?)", fileIDs).Find(&results).Error

	return
}

// GetFromCategoryID 通过CategoryID获取内容  档案类别
func (obj *_EntFileMgr) GetFromCategoryID(categoryID int) (results []*EntFile, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EntFile{}).Where("`CategoryID` = ?", categoryID).Find(&results).Error

	return
}

// GetBatchFromCategoryID 批量查找  档案类别
func (obj *_EntFileMgr) GetBatchFromCategoryID(categoryIDs []int) (results []*EntFile, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EntFile{}).Where("`CategoryID` IN (?)", categoryIDs).Find(&results).Error

	return
}

// GetFromProjectID 通过ProjectID获取内容 文件所属项目
func (obj *_EntFileMgr) GetFromProjectID(projectID string) (results []*EntFile, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EntFile{}).Where("`ProjectID` = ?", projectID).Find(&results).Error

	return
}

// GetBatchFromProjectID 批量查找 文件所属项目
func (obj *_EntFileMgr) GetBatchFromProjectID(projectIDs []string) (results []*EntFile, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EntFile{}).Where("`ProjectID` IN (?)", projectIDs).Find(&results).Error

	return
}

// GetFromEnterpriseID 通过EnterpriseID获取内容 文件所属公司ID
func (obj *_EntFileMgr) GetFromEnterpriseID(enterpriseID string) (results []*EntFile, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EntFile{}).Where("`EnterpriseID` = ?", enterpriseID).Find(&results).Error

	return
}

// GetBatchFromEnterpriseID 批量查找 文件所属公司ID
func (obj *_EntFileMgr) GetBatchFromEnterpriseID(enterpriseIDs []string) (results []*EntFile, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EntFile{}).Where("`EnterpriseID` IN (?)", enterpriseIDs).Find(&results).Error

	return
}

// GetFromFileName 通过FileName获取内容 文件名
func (obj *_EntFileMgr) GetFromFileName(fileName string) (results []*EntFile, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EntFile{}).Where("`FileName` = ?", fileName).Find(&results).Error

	return
}

// GetBatchFromFileName 批量查找 文件名
func (obj *_EntFileMgr) GetBatchFromFileName(fileNames []string) (results []*EntFile, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EntFile{}).Where("`FileName` IN (?)", fileNames).Find(&results).Error

	return
}

// GetFromFileURL 通过FileURL获取内容 文件存放地址
func (obj *_EntFileMgr) GetFromFileURL(fileURL string) (results []*EntFile, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EntFile{}).Where("`FileURL` = ?", fileURL).Find(&results).Error

	return
}

// GetBatchFromFileURL 批量查找 文件存放地址
func (obj *_EntFileMgr) GetBatchFromFileURL(fileURLs []string) (results []*EntFile, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EntFile{}).Where("`FileURL` IN (?)", fileURLs).Find(&results).Error

	return
}

// GetFromFileTypeID 通过FileTypeID获取内容 文件类型
func (obj *_EntFileMgr) GetFromFileTypeID(fileTypeID int) (results []*EntFile, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EntFile{}).Where("`FileTypeID` = ?", fileTypeID).Find(&results).Error

	return
}

// GetBatchFromFileTypeID 批量查找 文件类型
func (obj *_EntFileMgr) GetBatchFromFileTypeID(fileTypeIDs []int) (results []*EntFile, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EntFile{}).Where("`FileTypeID` IN (?)", fileTypeIDs).Find(&results).Error

	return
}

// GetFromStatus 通过Status获取内容 0表示没有被借出，1表示已经借出
func (obj *_EntFileMgr) GetFromStatus(status int8) (results []*EntFile, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EntFile{}).Where("`Status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找 0表示没有被借出，1表示已经借出
func (obj *_EntFileMgr) GetBatchFromStatus(statuss []int8) (results []*EntFile, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EntFile{}).Where("`Status` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromUserID 通过UserID获取内容 文件上传人
func (obj *_EntFileMgr) GetFromUserID(userID string) (results []*EntFile, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EntFile{}).Where("`UserID` = ?", userID).Find(&results).Error

	return
}

// GetBatchFromUserID 批量查找 文件上传人
func (obj *_EntFileMgr) GetBatchFromUserID(userIDs []string) (results []*EntFile, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EntFile{}).Where("`UserID` IN (?)", userIDs).Find(&results).Error

	return
}

// GetFromFileCabinet 通过FileCabinet获取内容  存放档案柜编号
func (obj *_EntFileMgr) GetFromFileCabinet(fileCabinet string) (results []*EntFile, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EntFile{}).Where("`FileCabinet` = ?", fileCabinet).Find(&results).Error

	return
}

// GetBatchFromFileCabinet 批量查找  存放档案柜编号
func (obj *_EntFileMgr) GetBatchFromFileCabinet(fileCabinets []string) (results []*EntFile, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EntFile{}).Where("`FileCabinet` IN (?)", fileCabinets).Find(&results).Error

	return
}

// GetFromIsDelete 通过IsDelete获取内容 是否删除
func (obj *_EntFileMgr) GetFromIsDelete(isDelete bool) (results []*EntFile, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EntFile{}).Where("`IsDelete` = ?", isDelete).Find(&results).Error

	return
}

// GetBatchFromIsDelete 批量查找 是否删除
func (obj *_EntFileMgr) GetBatchFromIsDelete(isDeletes []bool) (results []*EntFile, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EntFile{}).Where("`IsDelete` IN (?)", isDeletes).Find(&results).Error

	return
}

// GetFromUpdateTime 通过UpdateTime获取内容 文件上传时间
func (obj *_EntFileMgr) GetFromUpdateTime(updateTime time.Time) (results []*EntFile, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EntFile{}).Where("`UpdateTime` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找 文件上传时间
func (obj *_EntFileMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*EntFile, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EntFile{}).Where("`UpdateTime` IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromCreatTime 通过CreatTime获取内容 记录创建时间
func (obj *_EntFileMgr) GetFromCreatTime(creatTime time.Time) (results []*EntFile, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EntFile{}).Where("`CreatTime` = ?", creatTime).Find(&results).Error

	return
}

// GetBatchFromCreatTime 批量查找 记录创建时间
func (obj *_EntFileMgr) GetBatchFromCreatTime(creatTimes []time.Time) (results []*EntFile, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EntFile{}).Where("`CreatTime` IN (?)", creatTimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_EntFileMgr) FetchByPrimaryKey(autoID int64) (result EntFile, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EntFile{}).Where("`AutoID` = ?", autoID).Find(&result).Error

	return
}
