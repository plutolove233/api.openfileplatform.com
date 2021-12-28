package mysqlModel

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _EntFileTypeMgr struct {
	*_BaseMgr
}

// EntFileTypeMgr open func
func EntFileTypeMgr(db *gorm.DB) *_EntFileTypeMgr {
	if db == nil {
		panic(fmt.Errorf("EntFileTypeMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EntFileTypeMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ent_file_type"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EntFileTypeMgr) GetTableName() string {
	return "ent_file_type"
}

// Reset 重置gorm会话
func (obj *_EntFileTypeMgr) Reset() *_EntFileTypeMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_EntFileTypeMgr) Get() (result EntFileType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EntFileType{}).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_EntFileTypeMgr) Gets() (results []*EntFileType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EntFileType{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_EntFileTypeMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(EntFileType{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithAutoID AutoID获取
func (obj *_EntFileTypeMgr) WithAutoID(autoID int64) Option {
	return optionFunc(func(o *options) { o.query["AutoID"] = autoID })
}

// WithFileTypeID FileTypeID获取 文件(格式)类型
func (obj *_EntFileTypeMgr) WithFileTypeID(fileTypeID int) Option {
	return optionFunc(func(o *options) { o.query["FileTypeID"] = fileTypeID })
}

// WithFileTypeName FileTypeName获取 文件格式名称：DOC，PDF
func (obj *_EntFileTypeMgr) WithFileTypeName(fileTypeName string) Option {
	return optionFunc(func(o *options) { o.query["FileTypeName"] = fileTypeName })
}

// GetByOption 功能选项模式获取
func (obj *_EntFileTypeMgr) GetByOption(opts ...Option) (result EntFileType, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(EntFileType{}).Where(options.query).Find(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_EntFileTypeMgr) GetByOptions(opts ...Option) (results []*EntFileType, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(EntFileType{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromAutoID 通过AutoID获取内容
func (obj *_EntFileTypeMgr) GetFromAutoID(autoID int64) (result EntFileType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EntFileType{}).Where("`AutoID` = ?", autoID).Find(&result).Error

	return
}

// GetBatchFromAutoID 批量查找
func (obj *_EntFileTypeMgr) GetBatchFromAutoID(autoIDs []int64) (results []*EntFileType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EntFileType{}).Where("`AutoID` IN (?)", autoIDs).Find(&results).Error

	return
}

// GetFromFileTypeID 通过FileTypeID获取内容 文件(格式)类型
func (obj *_EntFileTypeMgr) GetFromFileTypeID(fileTypeID int) (results []*EntFileType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EntFileType{}).Where("`FileTypeID` = ?", fileTypeID).Find(&results).Error

	return
}

// GetBatchFromFileTypeID 批量查找 文件(格式)类型
func (obj *_EntFileTypeMgr) GetBatchFromFileTypeID(fileTypeIDs []int) (results []*EntFileType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EntFileType{}).Where("`FileTypeID` IN (?)", fileTypeIDs).Find(&results).Error

	return
}

// GetFromFileTypeName 通过FileTypeName获取内容 文件格式名称：DOC，PDF
func (obj *_EntFileTypeMgr) GetFromFileTypeName(fileTypeName string) (results []*EntFileType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EntFileType{}).Where("`FileTypeName` = ?", fileTypeName).Find(&results).Error

	return
}

// GetBatchFromFileTypeName 批量查找 文件格式名称：DOC，PDF
func (obj *_EntFileTypeMgr) GetBatchFromFileTypeName(fileTypeNames []string) (results []*EntFileType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EntFileType{}).Where("`FileTypeName` IN (?)", fileTypeNames).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_EntFileTypeMgr) FetchByPrimaryKey(autoID int64) (result EntFileType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EntFileType{}).Where("`AutoID` = ?", autoID).Find(&result).Error

	return
}
