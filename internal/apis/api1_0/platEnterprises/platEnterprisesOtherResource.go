// coding:utf-8
// @Author:PigKnight
// @Date:2022/2/10 16:22
// @Software: GoLand

package platEnterprises

import (
	"api.openfileplatform.com/internal/globals/codes"
	"api.openfileplatform.com/internal/globals/responseParser"
	"api.openfileplatform.com/internal/globals/snowflake"
	"api.openfileplatform.com/internal/models/ginModels"
	"api.openfileplatform.com/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
)

type platformEnterpriseParser struct {
	EnterpriseName 		string 	`json:"EnterpriseName" form:"EnterpriseName" binding:"required"`
	EnterprisePassword 	string	`json:"EnterprisePassword" form:"EnterprisePassword" binding:"required"`
	Location 			string	`json:"Location" form:"Location" binding:"required"`
	EnterprisePhone 	string	`json:"EnterprisePhone" form:"EnterprisePhone" binding:"required"`
}

func (*PlatformEnterpriseApi)Register(c *gin.Context){
	var parser platformEnterpriseParser
	err := c.ShouldBind(&parser)
	if err != nil {
		responseParser.JsonParameterIllegal(c,"平台企业注册参数不合法",err)
		return
	}

	var platEnterpriseService services.PlatEnterpriseService
	platEnterpriseService.EnterpriseName = parser.EnterpriseName
	err = platEnterpriseService.Get()
	if err == nil {
		responseParser.JsonDataExist(c,"企业名称已存在")
		return
	} else if err.Error() != "record not found" {
		responseParser.JsonDBError(c,"",err)
		return
	}

	hash,err := bcrypt.GenerateFromPassword([]byte(parser.EnterprisePassword),bcrypt.DefaultCost)
	if err != nil {
		responseParser.JsonInternalError(c,"企业密码加密失败",err)
		return
	}

	platEnterpriseService.EnterpriseID = snowflake.GetSnowflakeID()
	platEnterpriseService.EnterprisePhone = parser.EnterprisePhone
	platEnterpriseService.EnterpriseURL = ""
	platEnterpriseService.Location = parser.Location
	platEnterpriseService.AdminID = ""
	platEnterpriseService.LogoPicURL = "pic/logo/default.png"
	platEnterpriseService.EnterprisePassword = string(hash)

	err1 := platEnterpriseService.Add()
	if err1 != nil {
		responseParser.JsonDBError(c,"企业数据信息保存失败",err1)
		return
	}

	responseParser.JsonOK(c,"企业信息注册成功",platEnterpriseService)
}

type changePasswordParser struct {
	EnterpriseID 	string	`json:"EnterpriseID" form:"EnterpriseID" binding:"required"`
	Password		string	`json:"Password" form:"Password" binding:"required"`
	NewPassword		string	`json:"NewPassword" form:"NewPassword" binding:"required"`
}

func (*PlatformEnterpriseApi) ChangePassword(c *gin.Context)  {
	var parser changePasswordParser
	err := c.ShouldBind(&parser)
	if err != nil {
		responseParser.JsonParameterIllegal(c,"获取修改密码参数错误",err)
		return
	}

	temp,ok := c.Get("user")
	if !ok {
		responseParser.JsonNotData(c,"用户未登录",nil)
		return
	}

	user := temp.(ginModels.UserModel)

	var platEnterpriseService services.PlatEnterpriseService
	platEnterpriseService.EnterpriseID = parser.EnterpriseID
	err = platEnterpriseService.Get()

	if err != nil {
		responseParser.JsonNotData(c,"没有该企业的信息",err)
		return
	}

	if !user.IsPlatUser {
		if user.UserID != platEnterpriseService.AdminID {
			c.JSON(200, gin.H{
				"code":    codes.AccessDenied,
				"message": "用户没有权限修改企业密码",
			})
			return
		}
	}

	err = bcrypt.CompareHashAndPassword([]byte(platEnterpriseService.EnterprisePassword),[]byte(parser.Password))
	if err != nil {
		responseParser.JsonDataError(c,"密码错误",err)
		return
	}

	hash,err1 := bcrypt.GenerateFromPassword([]byte(parser.NewPassword),bcrypt.DefaultCost)
	if err1 != nil {
		responseParser.JsonInternalError(c,"密码加密失败",err1)
		return
	}

	err = platEnterpriseService.Update(map[string]interface{}{
		"EnterprisePassword":string(hash),
	})
	if err != nil {
		responseParser.JsonDBError(c,"密码更新失败",err)
		return
	}

	c.JSON(200,gin.H{
		"code":codes.OK,
		"message":"密码修改成功",
	})
}

type refreshPasswordParser struct {
	EnterpriseID 	string	`json:"EnterpriseID" form:"EnterpriseID" binding:"required"`
}

func (*PlatformEnterpriseApi)RefreshPassword(c *gin.Context)  {
	var parser refreshPasswordParser
	err := c.ShouldBind(&parser)
	if err != nil{
		responseParser.JsonParameterIllegal(c,"重置密码参数获取错误",err)
		return
	}

	var platEnterpriseService services.PlatEnterpriseService
	platEnterpriseService.EnterpriseID = parser.EnterpriseID
	err = platEnterpriseService.Get()
	if err != nil {
		responseParser.JsonNotData(c,"企业信息不存在",err)
		return
	}

	temp,ok := c.Get("user")
	if !ok {
		responseParser.JsonNotData(c,"用户未登录",err)
		return
	}

	user := temp.(ginModels.UserModel)

	if !user.IsPlatUser {
		c.JSON(200,gin.H{
			"code":codes.UnauthorizedUserId,
			"message":"该用户没有权限重置密码",
		})
		return
	}

	hash,err := bcrypt.GenerateFromPassword([]byte(viper.GetString("user.defaultPassword")),bcrypt.DefaultCost)
	if err != nil {
		responseParser.JsonInternalError(c,"密码重置加密错误",err)
		return
	}

	err = platEnterpriseService.Update(map[string]interface{}{
		"EnterprisePassword":string(hash),
	})

	if err != nil {
		responseParser.JsonDBError(c,"密码重置失败",err)
		return
	}

	c.JSON(200,gin.H{
		"code":codes.OK,
		"message":"密码重置成功",
	})
}