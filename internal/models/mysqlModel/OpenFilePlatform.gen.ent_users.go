package mysqlModel

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _EntUsersMgr struct {
	*_BaseMgr
}

// EntUsersMgr open func
func EntUsersMgr(db *gorm.DB) *_EntUsersMgr {
	if db == nil {
		panic(fmt.Errorf("EntUsersMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EntUsersMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ent_users"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EntUsersMgr) GetTableName() string {
	return "ent_users"
}

// Reset 重置gorm会话
func (obj *_EntUsersMgr) Reset() *_EntUsersMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_EntUsersMgr) Get() (result EntUsers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EntUsers{}).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_EntUsersMgr) Gets() (results []*EntUsers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EntUsers{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_EntUsersMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(EntUsers{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithAutoID AutoID获取
func (obj *_EntUsersMgr) WithAutoID(autoID int64) Option {
	return optionFunc(func(o *options) { o.query["AutoID"] = autoID })
}

// WithUserID UserID获取
func (obj *_EntUsersMgr) WithUserID(userID string) Option {
	return optionFunc(func(o *options) { o.query["UserID"] = userID })
}

// WithEnterpriseID EnterpriseID获取 用户所属企业ID
func (obj *_EntUsersMgr) WithEnterpriseID(enterpriseID string) Option {
	return optionFunc(func(o *options) { o.query["EnterpriseID"] = enterpriseID })
}

// WithAccount Account获取  账号（默认手机号）
func (obj *_EntUsersMgr) WithAccount(account string) Option {
	return optionFunc(func(o *options) { o.query["Account"] = account })
}

// WithPassword Password获取
func (obj *_EntUsersMgr) WithPassword(password string) Option {
	return optionFunc(func(o *options) { o.query["Password"] = password })
}

// WithUserName UserName获取 用户姓名
func (obj *_EntUsersMgr) WithUserName(userName string) Option {
	return optionFunc(func(o *options) { o.query["UserName"] = userName })
}

// WithPhone Phone获取
func (obj *_EntUsersMgr) WithPhone(phone string) Option {
	return optionFunc(func(o *options) { o.query["Phone"] = phone })
}

// WithEmail Email获取
func (obj *_EntUsersMgr) WithEmail(email string) Option {
	return optionFunc(func(o *options) { o.query["Email"] = email })
}

// WithFacePicURL FacePicUrl获取 头像
func (obj *_EntUsersMgr) WithFacePicURL(facePicURL string) Option {
	return optionFunc(func(o *options) { o.query["FacePicUrl"] = facePicURL })
}

// WithIsAdmin IsAdmin获取
func (obj *_EntUsersMgr) WithIsAdmin(isAdmin int) Option {
	return optionFunc(func(o *options) { o.query["IsAdmin"] = isAdmin })
}

// WithIsDeleted IsDeleted获取 是否已删除：0--未删除；1--已经删除
func (obj *_EntUsersMgr) WithIsDeleted(isDeleted bool) Option {
	return optionFunc(func(o *options) { o.query["IsDeleted"] = isDeleted })
}

// WithCreateTime CreateTime获取
func (obj *_EntUsersMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["CreateTime"] = createTime })
}

// WithToken Token获取
func (obj *_EntUsersMgr) WithToken(token string) Option {
	return optionFunc(func(o *options) { o.query["Token"] = token })
}

// GetByOption 功能选项模式获取
func (obj *_EntUsersMgr) GetByOption(opts ...Option) (result EntUsers, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(EntUsers{}).Where(options.query).Find(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_EntUsersMgr) GetByOptions(opts ...Option) (results []*EntUsers, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(EntUsers{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromAutoID 通过AutoID获取内容
func (obj *_EntUsersMgr) GetFromAutoID(autoID int64) (result EntUsers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EntUsers{}).Where("`AutoID` = ?", autoID).Find(&result).Error

	return
}

// GetBatchFromAutoID 批量查找
func (obj *_EntUsersMgr) GetBatchFromAutoID(autoIDs []int64) (results []*EntUsers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EntUsers{}).Where("`AutoID` IN (?)", autoIDs).Find(&results).Error

	return
}

// GetFromUserID 通过UserID获取内容
func (obj *_EntUsersMgr) GetFromUserID(userID string) (results []*EntUsers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EntUsers{}).Where("`UserID` = ?", userID).Find(&results).Error

	return
}

// GetBatchFromUserID 批量查找
func (obj *_EntUsersMgr) GetBatchFromUserID(userIDs []string) (results []*EntUsers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EntUsers{}).Where("`UserID` IN (?)", userIDs).Find(&results).Error

	return
}

// GetFromEnterpriseID 通过EnterpriseID获取内容 用户所属企业ID
func (obj *_EntUsersMgr) GetFromEnterpriseID(enterpriseID string) (results []*EntUsers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EntUsers{}).Where("`EnterpriseID` = ?", enterpriseID).Find(&results).Error

	return
}

// GetBatchFromEnterpriseID 批量查找 用户所属企业ID
func (obj *_EntUsersMgr) GetBatchFromEnterpriseID(enterpriseIDs []string) (results []*EntUsers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EntUsers{}).Where("`EnterpriseID` IN (?)", enterpriseIDs).Find(&results).Error

	return
}

// GetFromAccount 通过Account获取内容  账号（默认手机号）
func (obj *_EntUsersMgr) GetFromAccount(account string) (results []*EntUsers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EntUsers{}).Where("`Account` = ?", account).Find(&results).Error

	return
}

// GetBatchFromAccount 批量查找  账号（默认手机号）
func (obj *_EntUsersMgr) GetBatchFromAccount(accounts []string) (results []*EntUsers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EntUsers{}).Where("`Account` IN (?)", accounts).Find(&results).Error

	return
}

// GetFromPassword 通过Password获取内容
func (obj *_EntUsersMgr) GetFromPassword(password string) (results []*EntUsers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EntUsers{}).Where("`Password` = ?", password).Find(&results).Error

	return
}

// GetBatchFromPassword 批量查找
func (obj *_EntUsersMgr) GetBatchFromPassword(passwords []string) (results []*EntUsers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EntUsers{}).Where("`Password` IN (?)", passwords).Find(&results).Error

	return
}

// GetFromUserName 通过UserName获取内容 用户姓名
func (obj *_EntUsersMgr) GetFromUserName(userName string) (results []*EntUsers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EntUsers{}).Where("`UserName` = ?", userName).Find(&results).Error

	return
}

// GetBatchFromUserName 批量查找 用户姓名
func (obj *_EntUsersMgr) GetBatchFromUserName(userNames []string) (results []*EntUsers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EntUsers{}).Where("`UserName` IN (?)", userNames).Find(&results).Error

	return
}

// GetFromPhone 通过Phone获取内容
func (obj *_EntUsersMgr) GetFromPhone(phone string) (results []*EntUsers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EntUsers{}).Where("`Phone` = ?", phone).Find(&results).Error

	return
}

// GetBatchFromPhone 批量查找
func (obj *_EntUsersMgr) GetBatchFromPhone(phones []string) (results []*EntUsers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EntUsers{}).Where("`Phone` IN (?)", phones).Find(&results).Error

	return
}

// GetFromEmail 通过Email获取内容
func (obj *_EntUsersMgr) GetFromEmail(email string) (results []*EntUsers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EntUsers{}).Where("`Email` = ?", email).Find(&results).Error

	return
}

// GetBatchFromEmail 批量查找
func (obj *_EntUsersMgr) GetBatchFromEmail(emails []string) (results []*EntUsers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EntUsers{}).Where("`Email` IN (?)", emails).Find(&results).Error

	return
}

// GetFromFacePicURL 通过FacePicUrl获取内容 头像
func (obj *_EntUsersMgr) GetFromFacePicURL(facePicURL string) (results []*EntUsers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EntUsers{}).Where("`FacePicUrl` = ?", facePicURL).Find(&results).Error

	return
}

// GetBatchFromFacePicURL 批量查找 头像
func (obj *_EntUsersMgr) GetBatchFromFacePicURL(facePicURLs []string) (results []*EntUsers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EntUsers{}).Where("`FacePicUrl` IN (?)", facePicURLs).Find(&results).Error

	return
}

// GetFromIsAdmin 通过IsAdmin获取内容
func (obj *_EntUsersMgr) GetFromIsAdmin(isAdmin int) (results []*EntUsers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EntUsers{}).Where("`IsAdmin` = ?", isAdmin).Find(&results).Error

	return
}

// GetBatchFromIsAdmin 批量查找
func (obj *_EntUsersMgr) GetBatchFromIsAdmin(isAdmins []int) (results []*EntUsers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EntUsers{}).Where("`IsAdmin` IN (?)", isAdmins).Find(&results).Error

	return
}

// GetFromIsDeleted 通过IsDeleted获取内容 是否已删除：0--未删除；1--已经删除
func (obj *_EntUsersMgr) GetFromIsDeleted(isDeleted bool) (results []*EntUsers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EntUsers{}).Where("`IsDeleted` = ?", isDeleted).Find(&results).Error

	return
}

// GetBatchFromIsDeleted 批量查找 是否已删除：0--未删除；1--已经删除
func (obj *_EntUsersMgr) GetBatchFromIsDeleted(isDeleteds []bool) (results []*EntUsers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EntUsers{}).Where("`IsDeleted` IN (?)", isDeleteds).Find(&results).Error

	return
}

// GetFromCreateTime 通过CreateTime获取内容
func (obj *_EntUsersMgr) GetFromCreateTime(createTime time.Time) (results []*EntUsers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EntUsers{}).Where("`CreateTime` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找
func (obj *_EntUsersMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*EntUsers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EntUsers{}).Where("`CreateTime` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromToken 通过Token获取内容
func (obj *_EntUsersMgr) GetFromToken(token string) (results []*EntUsers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EntUsers{}).Where("`Token` = ?", token).Find(&results).Error

	return
}

// GetBatchFromToken 批量查找
func (obj *_EntUsersMgr) GetBatchFromToken(tokens []string) (results []*EntUsers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EntUsers{}).Where("`Token` IN (?)", tokens).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_EntUsersMgr) FetchByPrimaryKey(autoID int64) (result EntUsers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EntUsers{}).Where("`AutoID` = ?", autoID).Find(&result).Error

	return
}
