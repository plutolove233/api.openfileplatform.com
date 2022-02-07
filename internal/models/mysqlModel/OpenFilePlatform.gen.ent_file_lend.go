package mysqlModel

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _EntFileLendMgr struct {
	*_BaseMgr
}

// EntFileLendMgr open func
func EntFileLendMgr(db *gorm.DB) *_EntFileLendMgr {
	if db == nil {
		panic(fmt.Errorf("EntFileLendMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EntFileLendMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ent_file_lend"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EntFileLendMgr) GetTableName() string {
	return "ent_file_lend"
}

// Reset 重置gorm会话
func (obj *_EntFileLendMgr) Reset() *_EntFileLendMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_EntFileLendMgr) Get() (result EntFileLend, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EntFileLend{}).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_EntFileLendMgr) Gets() (results []*EntFileLend, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EntFileLend{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_EntFileLendMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(EntFileLend{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithAutoID AutoID获取
func (obj *_EntFileLendMgr) WithAutoID(autoID int64) Option {
	return optionFunc(func(o *options) { o.query["AutoID"] = autoID })
}

// WithFileID FileID获取  文件（档案）ID
func (obj *_EntFileLendMgr) WithFileID(fileID string) Option {
	return optionFunc(func(o *options) { o.query["FileID"] = fileID })
}

// WithEnterpriseID EnterpriseID获取 文件所属公司ID
func (obj *_EntFileLendMgr) WithEnterpriseID(enterpriseID string) Option {
	return optionFunc(func(o *options) { o.query["EnterpriseID"] = enterpriseID })
}

// WithBorrowerID BorrowerID获取 文件借阅人
func (obj *_EntFileLendMgr) WithBorrowerID(borrowerID string) Option {
	return optionFunc(func(o *options) { o.query["BorrowerID"] = borrowerID })
}

// WithBorrowTime BorrowTime获取 借出时间
func (obj *_EntFileLendMgr) WithBorrowTime(borrowTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["BorrowTime"] = borrowTime })
}

// WithBorrowTerm BorrowTerm获取  借阅周期
func (obj *_EntFileLendMgr) WithBorrowTerm(borrowTerm int8) Option {
	return optionFunc(func(o *options) { o.query["BorrowTerm"] = borrowTerm })
}

// WithReturnTime ReturnTime获取 归还时间
func (obj *_EntFileLendMgr) WithReturnTime(returnTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["ReturnTime"] = returnTime })
}

// WithIsDelete IsDelete获取 是否删除
func (obj *_EntFileLendMgr) WithIsDelete(isDelete bool) Option {
	return optionFunc(func(o *options) { o.query["IsDelete"] = isDelete })
}

// WithCreatTime CreatTime获取 记录创建时间
func (obj *_EntFileLendMgr) WithCreatTime(creatTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["CreatTime"] = creatTime })
}

// GetByOption 功能选项模式获取
func (obj *_EntFileLendMgr) GetByOption(opts ...Option) (result EntFileLend, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(EntFileLend{}).Where(options.query).Find(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_EntFileLendMgr) GetByOptions(opts ...Option) (results []*EntFileLend, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(EntFileLend{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromAutoID 通过AutoID获取内容
func (obj *_EntFileLendMgr) GetFromAutoID(autoID int64) (result EntFileLend, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EntFileLend{}).Where("`AutoID` = ?", autoID).Find(&result).Error

	return
}

// GetBatchFromAutoID 批量查找
func (obj *_EntFileLendMgr) GetBatchFromAutoID(autoIDs []int64) (results []*EntFileLend, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EntFileLend{}).Where("`AutoID` IN (?)", autoIDs).Find(&results).Error

	return
}

// GetFromFileID 通过FileID获取内容  文件（档案）ID
func (obj *_EntFileLendMgr) GetFromFileID(fileID string) (results []*EntFileLend, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EntFileLend{}).Where("`FileID` = ?", fileID).Find(&results).Error

	return
}

// GetBatchFromFileID 批量查找  文件（档案）ID
func (obj *_EntFileLendMgr) GetBatchFromFileID(fileIDs []string) (results []*EntFileLend, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EntFileLend{}).Where("`FileID` IN (?)", fileIDs).Find(&results).Error

	return
}

// GetFromEnterpriseID 通过EnterpriseID获取内容 文件所属公司ID
func (obj *_EntFileLendMgr) GetFromEnterpriseID(enterpriseID string) (results []*EntFileLend, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EntFileLend{}).Where("`EnterpriseID` = ?", enterpriseID).Find(&results).Error

	return
}

// GetBatchFromEnterpriseID 批量查找 文件所属公司ID
func (obj *_EntFileLendMgr) GetBatchFromEnterpriseID(enterpriseIDs []string) (results []*EntFileLend, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EntFileLend{}).Where("`EnterpriseID` IN (?)", enterpriseIDs).Find(&results).Error

	return
}

// GetFromBorrowerID 通过BorrowerID获取内容 文件借阅人
func (obj *_EntFileLendMgr) GetFromBorrowerID(borrowerID string) (results []*EntFileLend, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EntFileLend{}).Where("`BorrowerID` = ?", borrowerID).Find(&results).Error

	return
}

// GetBatchFromBorrowerID 批量查找 文件借阅人
func (obj *_EntFileLendMgr) GetBatchFromBorrowerID(borrowerIDs []string) (results []*EntFileLend, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EntFileLend{}).Where("`BorrowerID` IN (?)", borrowerIDs).Find(&results).Error

	return
}

// GetFromBorrowTime 通过BorrowTime获取内容 借出时间
func (obj *_EntFileLendMgr) GetFromBorrowTime(borrowTime time.Time) (results []*EntFileLend, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EntFileLend{}).Where("`BorrowTime` = ?", borrowTime).Find(&results).Error

	return
}

// GetBatchFromBorrowTime 批量查找 借出时间
func (obj *_EntFileLendMgr) GetBatchFromBorrowTime(borrowTimes []time.Time) (results []*EntFileLend, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EntFileLend{}).Where("`BorrowTime` IN (?)", borrowTimes).Find(&results).Error

	return
}

// GetFromBorrowTerm 通过BorrowTerm获取内容  借阅周期
func (obj *_EntFileLendMgr) GetFromBorrowTerm(borrowTerm int8) (results []*EntFileLend, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EntFileLend{}).Where("`BorrowTerm` = ?", borrowTerm).Find(&results).Error

	return
}

// GetBatchFromBorrowTerm 批量查找  借阅周期
func (obj *_EntFileLendMgr) GetBatchFromBorrowTerm(borrowTerms []int8) (results []*EntFileLend, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EntFileLend{}).Where("`BorrowTerm` IN (?)", borrowTerms).Find(&results).Error

	return
}

// GetFromReturnTime 通过ReturnTime获取内容 归还时间
func (obj *_EntFileLendMgr) GetFromReturnTime(returnTime time.Time) (results []*EntFileLend, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EntFileLend{}).Where("`ReturnTime` = ?", returnTime).Find(&results).Error

	return
}

// GetBatchFromReturnTime 批量查找 归还时间
func (obj *_EntFileLendMgr) GetBatchFromReturnTime(returnTimes []time.Time) (results []*EntFileLend, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EntFileLend{}).Where("`ReturnTime` IN (?)", returnTimes).Find(&results).Error

	return
}

// GetFromIsDelete 通过IsDelete获取内容 是否删除
func (obj *_EntFileLendMgr) GetFromIsDelete(isDelete bool) (results []*EntFileLend, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EntFileLend{}).Where("`IsDelete` = ?", isDelete).Find(&results).Error

	return
}

// GetBatchFromIsDelete 批量查找 是否删除
func (obj *_EntFileLendMgr) GetBatchFromIsDelete(isDeletes []bool) (results []*EntFileLend, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EntFileLend{}).Where("`IsDelete` IN (?)", isDeletes).Find(&results).Error

	return
}

// GetFromCreatTime 通过CreatTime获取内容 记录创建时间
func (obj *_EntFileLendMgr) GetFromCreatTime(creatTime time.Time) (results []*EntFileLend, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EntFileLend{}).Where("`CreatTime` = ?", creatTime).Find(&results).Error

	return
}

// GetBatchFromCreatTime 批量查找 记录创建时间
func (obj *_EntFileLendMgr) GetBatchFromCreatTime(creatTimes []time.Time) (results []*EntFileLend, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EntFileLend{}).Where("`CreatTime` IN (?)", creatTimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_EntFileLendMgr) FetchByPrimaryKey(autoID int64) (result EntFileLend, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EntFileLend{}).Where("`AutoID` = ?", autoID).Find(&result).Error

	return
}
