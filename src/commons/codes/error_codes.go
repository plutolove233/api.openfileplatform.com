package codes

const (
	//通用错误
	OK = 2000                            //成功

	NotData = 6001                       //无数据
	DataExist = 6002                     //数据已存在
	DataError = 6003                     //数据错误

	ParamError = 4000                    //参数错误
	SessionError = 4001                  //用户未登录
	LoginError = 4002                    //用户登录失败
	UserError = 4003                     //用户不存在或未激活
	RoleError = 4004                     //用户身份错误
	PwdError = 4005                      //密码错误
	REQError = 4006                      //非法请求或请求次数受限
	IPError = 4007                       //IP受限
	UpdateError  = 4008					 //上传失败

	InternetError = 5000                 //服务器内部错误
	DBError = 5001                       //数据库错误
	ThirdError = 5002                    //第三方系统错误
	IOError = 5003                       //文件读写错误
	UnknowError = 5004                   //未知错误
)