package api1_0

import (
	"api.openfileplatform.com/internal/globals/responseParser"
	"api.openfileplatform.com/internal/services"
	"api.openfileplatform.com/internal/utils/jwt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserApi struct{}
type usersService interface {
	Get() error
	Add() error
	GetPassword() string
	GetUserName() string
	GetIsAdmin() bool
	GetUserID() string
	SetAccount(string)
}

type EnterprisesApi struct{}
type enterpriseService interface {
}

type loginByPasswordParser struct {
	Account  string `form:"Account" json:"Account" binding:"required"`
	Password string `form:"Password" json:"Password" binding:"required"`
	UserType int    `form:"UserType" json:"UserType" binding:"required"`
}

func (*UserApi) LoginByPassword(c *gin.Context) {
	var parser loginByPasswordParser
	var err error
	//解析参数
	err = c.ShouldBind(&parser)
	if err != nil {
		responseParser.JsonParameterIllegal(c, "", err)
		return
	}

	var user usersService
	IsPlatUser := false
	token := ""
	if parser.UserType == 1 {
		user = &services.PlatUsersService{}
		IsPlatUser = true
	} else if parser.UserType == 2 {
		user = &services.EntUserService{}
	} else {
		responseParser.JsonParameterIllegal(c, "userType错误", nil)
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
			responseParser.JsonDataError(c, "密码错误！", err)
			return
		}
		responseParser.JsonInternalError(c, "", err)
		return
	}

	//生成token
	token, err = jwt.MakeToken(user.GetUserID(), IsPlatUser, user.GetIsAdmin())
	if err != nil {
		responseParser.JsonInternalError(c, "Token生成错误！", err)
		return
	}

	//返回
	responseParser.JsonOK(c, "登录成功！", gin.H{
		"UserName":user.GetUserName(),
		"token": token,
	})
	return
}

type refreshTokenParser struct {
	Token string `form:"Token" json:"Token" binding:"required"`
}

func (*UserApi) RefreshToken(c *gin.Context) {
	var parser refreshTokenParser
	var err error
	err = c.ShouldBind(&parser)
	if err != nil {
		responseParser.JsonParameterIllegal(c, "", err)
		return
	}

	token := parser.Token

	token, err = jwt.RefreshToken(token)
	if err != nil {
		responseParser.JsonDataError(c, "token已过期！", err)
		return
	}
	responseParser.JsonOK(c, "更新token成功！", gin.H{
		"token": token,
	})
	return
}
