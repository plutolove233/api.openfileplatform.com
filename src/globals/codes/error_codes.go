package codes

const (
	////通用错误
	//OK = 2000 //成功
	//
	//NotData   = 6001 //无数据
	//DataExist = 6002 //数据已存在
	//DataError = 6003 //数据错误
	//
	//ParamError   = 4000 //参数错误
	//SessionError = 4001 //用户未登录
	//LoginError   = 4002 //用户登录失败
	//UserError    = 4003 //用户不存在或未激活
	//RoleError    = 4004 //用户身份错误
	//PwdError     = 4005 //密码错误
	//REQError     = 4006 //非法请求或请求次数受限
	//IPError      = 4007 //IP受限
	//UpdateError  = 4008 //上传失败
	//
	//InternetError = 5000 //服务器内部错误
	//DBError       = 5001 //数据库错误
	//ThirdError    = 5002 //第三方系统错误
	//IOError       = 5003 //文件读写错误
	//UnknowError   = 5004 //未知错误

	//用户错误
	NotLoggedIn        = 1000 // 未登录
	UnauthorizedUserId = 1002 // 非法的用户Id
	Unauthorized       = 1003 // 未授权
	OperationFailure   = 1009 // 操作失败

	// 通用错误
	OK        = 200  // Success
	NotData   = 2001 // 没有数据
	DataExist = 2002 // 数据已存在
	DataError = 2003 // 数据错误

	// 网络级错误
	ParameterIllegal = 400 // 参数不合法
	RequestOverDue   = 402 // 请求已过期
	AccessDenied     = 403 // 拒绝访问
	RoutingNotExist  = 404 // 路由不存在
	RequestError     = 406 // 非法访问
	IPError          = 407 // IP受限

	// 系统级错误
	InternalError = 500  // 系统错误
	DBError       = 5001 // 数据库错误
)
