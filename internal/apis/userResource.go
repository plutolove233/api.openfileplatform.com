package apis

import (
	"api.openfileplatform.com/internal/globals/codes"
	"api.openfileplatform.com/internal/globals/responseParser"
	"api.openfileplatform.com/internal/globals/snowflake"
	"api.openfileplatform.com/internal/services"
	"api.openfileplatform.com/internal/utils/jwt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type UserApiImpl struct{}

type loginByPasswordParser struct {
	Account  string `form:"Account" json:"Account" binding:"required"`
	Password string `form:"Password" json:"Password" binding:"required"`
}

func (*UserApiImpl) LoginByPassword(c *gin.Context) {
	var parser loginByPasswordParser
	var err error
	//解析参数
	err = c.ShouldBind(&parser)
	if err != nil {
		responseParser.JsonParameterIllegal(c, err)
		return
	}
	//session := sessions.Default(c)
	userType := c.Request.Header.Get("userType")
	if userType == "" {
		responseParser.JsonIncompleteRequest(c,"请求头不完整")
		return
	}

	var user services.UserInterface
	token := ""
	if userType == "platform" {
		user = &services.PlatUsersService{}
	}else if userType == "enterprise"{
		user = &services.EntUserService{}
	}else {
		responseParser.JsonDataError(c,"userType数据错误")
		return
	}
	//查询账号信息
	user.SetAccount(parser.Account)
	err = user.Get()

	if err != nil {
		responseParser.JsonDBError(c, "", err)
		return
	}

	//验证密码
	var password []byte

	//password, err = base64.StdEncoding.DecodeString(parser.Password)
	//if err != nil {
	//	c.JSON(http.StatusOK, gin.H{
	//		"code":    codes.InternalError,
	//		"message": "解码失败！",
	//		"err":     err,
	//	})
	//	return
	//}
	//RSA := rsa.GetRSAHelper()
	//password, err = RSA.Decrypt(password)
	//if err != nil {
	//	c.JSON(http.StatusOK, gin.H{
	//		"code":    codes.InternalError,
	//		"message": "解密失败！",
	//		"err":     err,
	//	})
	//	return
	//}

	password = []byte(parser.Password)
	err = bcrypt.CompareHashAndPassword([]byte(user.GetPassword()), password)

	if err != nil {
		if err.Error() == "crypto/bcrypt: hashedPassword is not the hash of the given password" {
			responseParser.JsonAccessDenied(c,"密码错误！")
			return
		}
		responseParser.JsonInternalError(c, "密码加密失败", err)
		return
	}

	//获取登录的用户信息
	//user.UserID = platUser.UserID
	//user.Account = platUser.Account
	//user.IsPlatUser = true

	//temp, err := json.Marshal(user)
	//if err != nil {
	//	c.JSON(http.StatusOK, gin.H{
	//		"code":    codes.InternalError,
	//		"message": "系统错误！",
	//		"err":     err,
	//	})
	//	return
	//}
	//生成token
	token, err = jwt.MakeToken(user.GetUserID(), true, user.GetIsAdmin())
	if err != nil {
		responseParser.JsonInternalError(c,"Token生成错误！",err)
		return
	}
	//存入session
	//session.Set("tokenID", tokenID)
	//err = session.Save()
	//if err != nil {
	//	c.JSON(http.StatusOK, gin.H{
	//		"code":    codes.InternalError,
	//		"message": "session存储错误！",
	//		"err":     err,
	//	})
	//	return
	//}

	//返回
	responseParser.JsonOK(c, gin.H{
		"user":  user,
		"token": token,
	})
	return
}

type refreshTokenParser struct {
	Token string `form:"Token" json:"Token" binding:"required"`
}

func (*UserApiImpl) RefreshToken(c *gin.Context) {
	var parser refreshTokenParser
	var err error
	err = c.ShouldBindJSON(&parser)
	if err != nil {
		responseParser.JsonParameterIllegal(c,err)
		return
	}

	token := parser.Token

	token, err = jwt.RefreshToken(token)

	if err != nil {
		responseParser.JsonLoginError(c,"token已过期！",err)
		return
	}
	responseParser.JsonOK(c,token)
	return
}

//func (*UserApiImpl) Logout(c *gin.Context) {
//	temp, ok := c.GetUserList("TokenID")
//	if ok == false {
//		c.JSON(http.StatusOK, gin.H{
//			"code":    codes.AccessDenied,
//			"message": "Token错误！",
//		})
//		return
//	}
//	tokenID := temp.(string)
//
//	//删除session中的token
//	session := sessions.Default(c)
//	session.Delete("tokenID")
//	err := session.Save()
//	if err != nil {
//		c.JSON(http.StatusOK, gin.H{
//			"code":    codes.InternalError,
//			"message": "session连接错误！",
//			"err":     err,
//		})
//		return
//	}
//
//	//删除redis中的token
//	redisManager := database.GetRedisManager()
//	err = redisManager.Del("Token_"+tokenID).Err()
//	if err != nil {
//		c.JSON(http.StatusOK, gin.H{
//			"code":    codes.AccessDenied,
//			"message": "登出失败！",
//		})
//		return
//	}
//	c.JSON(http.StatusOK, gin.H{
//		"code":    codes.OK,
//		"message": "登出成功！",
//	})
//	return
//}

type RegisterParser struct {
	UserName string `form:"UserName" json:"UserName" binding:""`
	Account  string `form:"Account" json:"Account" binding:"required"`
	Password string `form:"Password" json:"Password" binding:"required"`
	Phone    string `form:"Phone" json:"Phone" binding:""`
	Email    string `form:"Email" json:"Email" binding:""`
}

func (*UserApiImpl) Register(c *gin.Context) {
	var parser RegisterParser
	var err error
	//解析参数
	err = c.ShouldBind(&parser)
	if err != nil {
		responseParser.JsonParameterIllegal(c,err)
		return
	}
	userType := c.Request.Header.Get("userType")
	if userType == "" {
		responseParser.JsonIncompleteRequest(c,"请求头不完整")
		return
	}
	var userInfo services.UserInterface

	if userType == "platform" {
		userInfo = &services.PlatUsersService{}
	}else if userType == "enterprise"{
		userInfo = &services.EntUserService{}
	}else {
		responseParser.JsonDataError(c,"userType数据错误")
		return
	}

	// 检验此注册方式是否已经注册过
	userInfo.SetAccount(parser.Account)
	err = userInfo.Get()
	if err == nil {
		responseParser.JsonDataExist(c,"账号已被注册！")
		return
	} else if err.Error() == "record not found" {
		// 未注册过则注册此登录方式
		hash, err := bcrypt.GenerateFromPassword([]byte(parser.Password), bcrypt.DefaultCost)
		if err != nil {
			responseParser.JsonInternalError(c,"密码加密错误！",err)
			return
		}
		userInfo.SetUserName(parser.UserName)
		userInfo.SetAccount(parser.Account)
		userInfo.SetPassword(string(hash))
		userInfo.SetUserID(snowflake.GetSnowflakeID())
		userInfo.SetPhone(parser.Phone)
		userInfo.SetEmail(parser.Email)
		err1 := userInfo.Add()
		if err1 != nil {
			responseParser.JsonDBError(c,"",err1)
			return
		}
	} else {
		responseParser.JsonDBError(c,"",err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    codes.OK,
		"message": "注册成功！",
	})
}
