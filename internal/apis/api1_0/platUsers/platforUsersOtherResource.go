// coding: utf-8
// @Author : lryself
// @Date : 2021/12/29 1:11
// @Software: GoLand

package platUsers

import (
	"api.openfileplatform.com/internal/globals/codes"
	"api.openfileplatform.com/internal/globals/responseParser"
	"api.openfileplatform.com/internal/globals/snowflake"
	"api.openfileplatform.com/internal/models/ginModels/platform"
	"api.openfileplatform.com/internal/services"
	"github.com/gin-gonic/gin"
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
		"password": string(hash),
		//"update_user": user.UserID,
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

	userID := Parser.UserID

	var platUser services.PlatUsersService
	platUser.UserID = userID
	err = platUser.Get()
	if err != nil {
		responseParser.JsonDBError(c, "", err)
		return
	}
	hash, err := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)
	if err != nil {
		responseParser.JsonInternalError(c, "", err)
		return
	}

	err = platUser.Update(map[string]interface{}{
		"password": string(hash),
	})
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    codes.DataError,
			"message": "更新密码出错！",
			"err":     err,
		})
		return
	}

	responseParser.JsonOK(c, "", nil)
}
