package codes

const (
	//通用错误
	OK 				= 200	//执行成功
	NotData 		= 2001	//没有请求数据
	DataExist 		= 2002	//请求数据存在
	DataError 		= 2003	//请求数据错误

	//用户端错误
	NotLogin		= 1000	//用户未登录

	//网络端错误
	ParamIllegal	= 400	//参数请求错误

	//系统错误
	InternetError 	= 500
	DBError			= 501
)