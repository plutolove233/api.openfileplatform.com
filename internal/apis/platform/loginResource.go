// coding: utf-8
// @Author : lryself
// @Date : 2021/4/10 2:43
// @Software: GoLand

package platform

import (
	"api.openfileplatform.com/internal/models/ginModels/platform"
	"net/http"

	"api.openfileplatform.com/internal/globals/codes"
	"api.openfileplatform.com/internal/globals/responseParser"
	"api.openfileplatform.com/internal/services"
	"api.openfileplatform.com/internal/utils/jwt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type LoginApiImpl struct{}

type loginByPasswordParser struct {
	Account  string `form:"Account" json:"Account" binding:"required"`
	Password string `form:"Password" json:"Password" binding:"required"`
}

func (*LoginApiImpl) LoginByPassword(c *gin.Context) {
	var Parser loginByPasswordParser
	var err error
	//解析参数
	err = c.ShouldBind(&Parser)
	if err != nil {
		responseParser.JsonParameterIllegal(c, err)
		return
	}
	//session := sessions.Default(c)

	user := platform.UserModel{}
	token := ""
	//查询账号信息
	var platUser services.PlatUsersService
	platUser.Account = Parser.Account
	err = platUser.Get()

	if err != nil {
		responseParser.JsonDBError(c, err)
		return
	}

	//验证密码
	var password []byte

	//password, err = base64.StdEncoding.DecodeString(Parser.Password)
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

	password = []byte(Parser.Password)
	err = bcrypt.CompareHashAndPassword([]byte(platUser.Password), password)

	if err != nil {
		if err.Error() == "crypto/bcrypt: hashedPassword is not the hash of the given password" {
			c.JSON(http.StatusOK, gin.H{
				"code":    codes.AccessDenied,
				"message": "密码错误！",
			})
			return
		}
		responseParser.JsonInternalError(c, err)
		return
	}

	//获取登录的用户信息
	user.UserID = platUser.UserID
	user.Account = platUser.Account
	user.IsPlatUser = true

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
	token, err = jwt.MakeToken(user.UserID, user.IsPlatUser, user.IsAdmin)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    codes.InternalError,
			"message": "Token生成错误！",
			"err":     err.Error(),
		})
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

type changePasswordParser struct {
	UserID      string `form:"UserID" json:"UserID" binding:"required"`
	Password    string `form:"Password" json:"Password" binding:"required"`
	NewPassword string `form:"NewPassword" json:"NewPassword" binding:"required"`
}

func (*LoginApiImpl) ChangePassword(c *gin.Context) {
	var Parser changePasswordParser
	var err error
	//解析参数
	err = c.ShouldBindJSON(&Parser)
	if err != nil {
		responseParser.JsonParameterIllegal(c, err)
		return
	}

	//查询账号信息
	temp, ok := c.Get("user")
	if !ok {
		c.JSON(http.StatusOK, gin.H{
			"code":    codes.NotData,
			"message": "用户未登录！",
		})
		return
	}

	user, _ := temp.(platform.UserModel)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    codes.ParameterIllegal,
			"message": "UserID错误！",
			"err":     err,
		})
		return
	}
	userID := Parser.UserID

	if user.VerifyAdminRole() {
		if user.UserID != userID {
			c.JSON(http.StatusOK, gin.H{
				"code":    codes.UnauthorizedUserId,
				"message": "只能修改自己的密码！",
			})
			return
		}
	}

	var platUser services.PlatUsersService
	platUser.UserID = userID
	err = platUser.Get()
	if err != nil {
		if err.Error() == "record not found" {
			c.JSON(http.StatusOK, gin.H{
				"code":    codes.NotData,
				"message": "无数据！",
				"err":     err,
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code":    codes.DBError,
			"message": "数据库错误！",
			"err":     err,
		})
		return
	}
	//验证密码
	err = bcrypt.CompareHashAndPassword([]byte(platUser.Password), []byte(Parser.Password))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    codes.DataError,
			"message": "密码错误！",
		})
		return
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(Parser.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    codes.DBError,
			"message": "数据库错误！",
			"err":     err,
		})
		return
	}

	err = platUser.Update(map[string]interface{}{
		"password":    string(hash),
		"update_user": user.UserID,
	})
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    codes.DBError,
			"message": "更新密码出错！",
			"err":     err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    codes.OK,
		"message": "修改成功！",
	})
}

type refreshTokenParser struct {
	Token string `form:"Token" json:"Token" binding:"required"`
}

func (*LoginApiImpl) RefreshToken(c *gin.Context) {
	var parser refreshTokenParser
	var err error
	err = c.ShouldBindJSON(&parser)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    codes.ParameterIllegal,
			"message": "参数错误！",
			"err":     err,
		})
		return
	}

	token := parser.Token

	token, err = jwt.RefreshToken(token)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    codes.AccessDenied,
			"message": "token已过期！",
			"err":     err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    codes.OK,
		"message": "成功！",
		"data":    token,
	})
	return
}

//func (*LoginApiImpl) Logout(c *gin.Context) {
//	temp, ok := c.Get("TokenID")
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
