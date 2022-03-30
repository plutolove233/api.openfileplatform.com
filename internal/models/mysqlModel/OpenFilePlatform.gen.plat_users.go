package mysqlModel

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _PlatUsersMgr struct {
	*_BaseMgr
}

// PlatUsersMgr open func
func PlatUsersMgr(db *gorm.DB) *_PlatUsersMgr {
	if db == nil {
		panic(fmt.Errorf("PlatUsersMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_PlatUsersMgr{_BaseMgr: &_BaseMgr{DB: db.Table("plat_users"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_PlatUsersMgr) GetTableName() string {
	return "plat_users"
}

// Reset 重置gorm会话
func (obj *_PlatUsersMgr) Reset() *_PlatUsersMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_PlatUsersMgr) Get() (result PlatUsers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PlatUsers{}).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_PlatUsersMgr) Gets() (results []*PlatUsers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PlatUsers{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_PlatUsersMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(PlatUsers{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithAutoID AutoID获取
func (obj *_PlatUsersMgr) WithAutoID(autoID int64) Option {
	return optionFunc(func(o *options) { o.query["AutoID"] = autoID })
}

// WithUserID UserID获取
func (obj *_PlatUsersMgr) WithUserID(userID string) Option {
	return optionFunc(func(o *options) { o.query["UserID"] = userID })
}

// WithUserName UserName获取
func (obj *_PlatUsersMgr) WithUserName(userName string) Option {
	return optionFunc(func(o *options) { o.query["UserName"] = userName })
}

// WithAccount Account获取
func (obj *_PlatUsersMgr) WithAccount(account string) Option {
	return optionFunc(func(o *options) { o.query["Account"] = account })
}

// WithPassword Password获取
func (obj *_PlatUsersMgr) WithPassword(password string) Option {
	return optionFunc(func(o *options) { o.query["Password"] = password })
}

// WithPhone Phone获取
func (obj *_PlatUsersMgr) WithPhone(phone string) Option {
	return optionFunc(func(o *options) { o.query["Phone"] = phone })
}

// WithEmail Email获取
func (obj *_PlatUsersMgr) WithEmail(email string) Option {
	return optionFunc(func(o *options) { o.query["Email"] = email })
}

// WithIsDeleted IsDeleted获取
func (obj *_PlatUsersMgr) WithIsDeleted(isDeleted bool) Option {
	return optionFunc(func(o *options) { o.query["IsDeleted"] = isDeleted })
}

// WithCreateTime CreateTime获取
func (obj *_PlatUsersMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["CreateTime"] = createTime })
}

// GetByOption 功能选项模式获取
func (obj *_PlatUsersMgr) GetByOption(opts ...Option) (result PlatUsers, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(PlatUsers{}).Where(options.query).Find(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_PlatUsersMgr) GetByOptions(opts ...Option) (results []*PlatUsers, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(PlatUsers{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromAutoID 通过AutoID获取内容
func (obj *_PlatUsersMgr) GetFromAutoID(autoID int64) (result PlatUsers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PlatUsers{}).Where("`AutoID` = ?", autoID).Find(&result).Error

	return
}

// GetBatchFromAutoID 批量查找
func (obj *_PlatUsersMgr) GetBatchFromAutoID(autoIDs []int64) (results []*PlatUsers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PlatUsers{}).Where("`AutoID` IN (?)", autoIDs).Find(&results).Error

	return
}

// GetFromUserID 通过UserID获取内容
func (obj *_PlatUsersMgr) GetFromUserID(userID string) (results []*PlatUsers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PlatUsers{}).Where("`UserID` = ?", userID).Find(&results).Error

	return
}

// GetBatchFromUserID 批量查找
func (obj *_PlatUsersMgr) GetBatchFromUserID(userIDs []string) (results []*PlatUsers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PlatUsers{}).Where("`UserID` IN (?)", userIDs).Find(&results).Error

	return
}

// GetFromUserName 通过UserName获取内容
func (obj *_PlatUsersMgr) GetFromUserName(userName string) (results []*PlatUsers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PlatUsers{}).Where("`UserName` = ?", userName).Find(&results).Error

	return
}

// GetBatchFromUserName 批量查找
func (obj *_PlatUsersMgr) GetBatchFromUserName(userNames []string) (results []*PlatUsers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PlatUsers{}).Where("`UserName` IN (?)", userNames).Find(&results).Error

	return
}

// GetFromAccount 通过Account获取内容
func (obj *_PlatUsersMgr) GetFromAccount(account string) (results []*PlatUsers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PlatUsers{}).Where("`Account` = ?", account).Find(&results).Error

	return
}

// GetBatchFromAccount 批量查找
func (obj *_PlatUsersMgr) GetBatchFromAccount(accounts []string) (results []*PlatUsers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PlatUsers{}).Where("`Account` IN (?)", accounts).Find(&results).Error

	return
}

// GetFromPassword 通过Password获取内容
func (obj *_PlatUsersMgr) GetFromPassword(password string) (results []*PlatUsers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PlatUsers{}).Where("`Password` = ?", password).Find(&results).Error

	return
}

// GetBatchFromPassword 批量查找
func (obj *_PlatUsersMgr) GetBatchFromPassword(passwords []string) (results []*PlatUsers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PlatUsers{}).Where("`Password` IN (?)", passwords).Find(&results).Error

	return
}

// GetFromPhone 通过Phone获取内容
func (obj *_PlatUsersMgr) GetFromPhone(phone string) (results []*PlatUsers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PlatUsers{}).Where("`Phone` = ?", phone).Find(&results).Error

	return
}

// GetBatchFromPhone 批量查找
func (obj *_PlatUsersMgr) GetBatchFromPhone(phones []string) (results []*PlatUsers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PlatUsers{}).Where("`Phone` IN (?)", phones).Find(&results).Error

	return
}

// GetFromEmail 通过Email获取内容
func (obj *_PlatUsersMgr) GetFromEmail(email string) (results []*PlatUsers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PlatUsers{}).Where("`Email` = ?", email).Find(&results).Error

	return
}

// GetBatchFromEmail 批量查找
func (obj *_PlatUsersMgr) GetBatchFromEmail(emails []string) (results []*PlatUsers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PlatUsers{}).Where("`Email` IN (?)", emails).Find(&results).Error

	return
}

// GetFromIsDeleted 通过IsDeleted获取内容
func (obj *_PlatUsersMgr) GetFromIsDeleted(isDeleted bool) (results []*PlatUsers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PlatUsers{}).Where("`IsDeleted` = ?", isDeleted).Find(&results).Error

	return
}

// GetBatchFromIsDeleted 批量查找
func (obj *_PlatUsersMgr) GetBatchFromIsDeleted(isDeleteds []bool) (results []*PlatUsers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PlatUsers{}).Where("`IsDeleted` IN (?)", isDeleteds).Find(&results).Error

	return
}

// GetFromCreateTime 通过CreateTime获取内容
func (obj *_PlatUsersMgr) GetFromCreateTime(createTime time.Time) (results []*PlatUsers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PlatUsers{}).Where("`CreateTime` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找
func (obj *_PlatUsersMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*PlatUsers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PlatUsers{}).Where("`CreateTime` IN (?)", createTimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_PlatUsersMgr) FetchByPrimaryKey(autoID int64) (result PlatUsers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PlatUsers{}).Where("`AutoID` = ?", autoID).Find(&result).Error

	return
}
