// coding: utf-8
// @Author : lryself
// @Date : 2021/12/29 1:11
// @Software: GoLand

package platUsers

import (
	"api.openfileplatform.com/internal/globals/codes"
	"api.openfileplatform.com/internal/globals/responseParser"
	"api.openfileplatform.com/internal/globals/snowflake"
	"api.openfileplatform.com/internal/models/ginModels"
	"api.openfileplatform.com/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type RegisterParser struct {
	UserName string `form:"UserName" json:"UserName" binding:""`
	Account  string `form:"Account" json:"Account" binding:"required"`
	Password string `form:"Password" json:"Password" binding:"required"`
	Phone    string `form:"Phone" json:"Phone" binding:""`
	Email    string `form:"Email" json:"Email" binding:""`
}

func (*PlatformUserApi) Register(c *gin.Context) {
	var Parser RegisterParser
	var err error
	//解析参数
	err = c.ShouldBind(&Parser)
	if err != nil {
		responseParser.JsonParameterIllegal(c, "", err)
		return
	}

	var platUsersService services.PlatUsersService

	// 检验此注册方式是否已经注册过
	platUsersService.Account = Parser.Account
	err = platUsersService.Get()
	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    codes.DataExist,
			"message": "账号已被注册！",
		})
		return
	} else if err.Error() != "record not found" {
		responseParser.JsonDBError(c, "", err)
	}
	// 未注册过则注册此登录方式
	hash, err := bcrypt.GenerateFromPassword([]byte(Parser.Password), bcrypt.DefaultCost)
	if err != nil {
		responseParser.JsonInternalError(c, "密码加密错误", err)
		return
	}
	platUsersService.UserName = Parser.UserName
	platUsersService.Account = Parser.Account
	platUsersService.Password = string(hash)
	platUsersService.UserID = snowflake.GetSnowflakeID()
	platUsersService.Phone = Parser.Phone
	platUsersService.Email = Parser.Email
	err1 := platUsersService.Add()
	if err1 != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    codes.DBError,
			"message": "数据库错误！",
			"err":     err1,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    codes.OK,
		"message": "注册成功！",
	})
}

type changePasswordParser struct {
	UserID      string `form:"UserID" json:"UserID" binding:"required"`
	Password    string `form:"Password" json:"Password" binding:"required"`
	NewPassword string `form:"NewPassword" json:"NewPassword" binding:"required"`
}

func (*PlatformUserApi) ChangePassword(c *gin.Context) {
	var Parser changePasswordParser
	var err error
	//解析参数
	err = c.ShouldBind(&Parser)
	if err != nil {
		responseParser.JsonParameterIllegal(c, "", err)
		return
	}

	//查询账号信息
	temp, ok := c.Get("user")
	if !ok {
		responseParser.JsonNotData(c, "用户未登录", nil)
		return
	}

	user := temp.(ginModels.UserModel)
	userID := Parser.UserID

	if !user.VerifyAdminRole() {
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
		responseParser.JsonDBError(c, "", err)
		return
	}
	//验证密码
	err = bcrypt.CompareHashAndPassword([]byte(platUser.Password), []byte(Parser.Password))
	if err != nil {
		responseParser.JsonDataError(c, "密码错误！", err)
		return
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(Parser.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		responseParser.JsonDBError(c, "", err)
		return
	}

	err = platUser.Update(map[string]interface{}{
		"password": string(hash),
	})
	if err != nil {
		responseParser.JsonDBError(c, "密码出错", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    codes.OK,
		"message": "修改成功！",
	})
}

type refreshPasswordParser struct {
	UserID string `form:"UserID" json:"UserID" binding:"required"`
}

func (*PlatformUserApi) RefreshPassword(c *gin.Context) {
	var Parser refreshPasswordParser
	var err error
	//解析参数
	err = c.ShouldBind(&Parser)
	if err != nil {
		responseParser.JsonParameterIllegal(c, "", err)
		return
	}

	//查询账号信息
	temp, ok := c.Get("user")
	if !ok {
		responseParser.JsonNotData(c, "用户未登录", nil)
		return
	}

	user := temp.(ginModels.UserModel)
	userID := Parser.UserID

	if !user.VerifyAdminRole() {
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
		responseParser.JsonDBError(c, "", err)
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(viper.GetString("user.defaultPassword")), bcrypt.DefaultCost)
	if err != nil {
		responseParser.JsonDBError(c, "", err)
		return
	}

	err = platUser.Update(map[string]interface{}{
		"password": string(hash),
	})
	if err != nil {
		responseParser.JsonDBError(c, "密码出错", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    codes.OK,
		"message": "密码重置成功！",
	})
}

type setEnterpriseAdminParser struct {
	EnterpriseID	string	`json:"EnterpriseID" form:"EnterpriseID" binding:"required"`
	UserID			string	`json:"UserID" form:"UserID" binding:"required"`
}

func (*PlatformUserApi)SetEnterpriseAdmin(c *gin.Context)  {
	var parser setEnterpriseAdminParser
	err := c.ShouldBind(&parser)
	if err != nil {
		responseParser.JsonParameterIllegal(c,"获取企业管理员信息失败",err)
		return
	}
	temp,ok := c.Get("user")
	if !ok {
		responseParser.JsonNotData(c,"用户未登录",nil)
		return
	}
	user := temp.(ginModels.UserModel)
	if !user.IsPlatUser {
		responseParser.JsonAccessDenied(c,"用户没有权限访问")
		return
	}

	var platUserService services.PlatUsersService
	platUserService.UserID = user.UserID
	err = platUserService.Get()
	if err != nil {
		responseParser.JsonNotData(c,"未找到平台用户",err)
		return
	}

	msg,err := platUserService.SetEntUserAdmin(parser.EnterpriseID,parser.UserID)
	if err != nil {
		responseParser.JsonDBError(c,msg,err)
		return
	}

	c.JSON(200,gin.H{
		"code":codes.OK,
		"message":msg,
	})
}

type removeAdminParser struct {
	EnterpriseID	string	`json:"EnterpriseID" form:"EnterpriseID" binding:"required"`
	UserID			string 	`json:"UserID" form:"UserID" binding:"required"`
}

func (*PlatformUserApi) RemoveEnterpriseAdmin(c *gin.Context) {
	var parser removeAdminParser
	err := c.ShouldBind(&parser)
	if err != nil {
		responseParser.JsonParameterIllegal(c,"获取管理员信息失败",err)
		return
	}

	temp,_ := c.Get("user")
	user := temp.(ginModels.UserModel)
	if user.IsPlatUser == false{
		responseParser.JsonAccessDenied(c, "没有权限修改管理员信息")
		return
	}
	platUser := services.PlatUsersService{}
	msg,err := platUser.RemoveEntUserAdmin(parser.EnterpriseID,parser.UserID)
	if err != nil {
		responseParser.JsonInternalError(c,msg,err)
		return
	}
	c.JSON(200,gin.H{
		"code":codes.OK,
		"message":msg,
	})
}