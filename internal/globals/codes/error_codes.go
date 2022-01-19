package codes

const (
	//用户错误
	NotLoggedIn        = "1000" // 未登录
	UnauthorizedUserId = "1002" // 非法的用户Id
	Unauthorized       = "1003" // 未授权
	OperationFailure   = "1009" // 操作失败

	// 通用错误
	OK        = "200"  // Success
	NotData   = "2001" // 没有数据
	DataExist = "2002" // 数据已存在
	DataError = "2003" // 数据错误

	// 网络级错误
	ParameterIllegal = "400" // 参数不合法
	RequestOverDue   = "402" // 请求已过期
	AccessDenied     = "403" // 拒绝访问
	RoutingNotExist  = "404" // 路由不存在
	RequestError     = "406" // 非法访问
	IPError          = "407" // IP受限

	// 系统级错误
	InternalError = "500"  // 系统错误
	DBError       = "5001" // 数据库错误
)
