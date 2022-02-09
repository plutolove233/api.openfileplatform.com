package entResource

import (
	"api.openfileplatform.com/internal/globals/codes"
	"api.openfileplatform.com/internal/globals/responseParser"
	"api.openfileplatform.com/internal/globals/snowflake"
	"api.openfileplatform.com/internal/models/ginModels"
	"api.openfileplatform.com/internal/services"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type PlatEntApiImpl struct {}

type EntRegisterParser struct {
	EnterpriseName     string `form:"EnterpriseName" json:"EnterpriseName" binding:"required"`
	EnterprisePassword string `form:"EnterprisePassword" json:"EnterprisePassword" binding:"required"`
	Location           string `form:"Location" json:"Location" binding:"required"`
	EnterprisePhone    string `form:"EnterprisePhone" json:"EnterprisePhone" binding:"required"`
}

//用于在平台注册企业
func ( *EntRegisterParser) Register(c *gin.Context)  {
	var parser EntRegisterParser
	err := c.ShouldBind(&parser)
	if err != nil {
		responseParser.JsonParameterIllegal(c,err)
		return
	}

	platEnt := services.PlatEnterpriseService{}
	platEnt.EnterpriseName = parser.EnterpriseName
	err = platEnt.Get()
	if err == nil {
		responseParser.JsonDataExist(c,"企业名称已存在")
		return
	}else if err.Error() == "record not found" {
		hash, err1 := bcrypt.GenerateFromPassword([]byte(parser.EnterprisePassword),bcrypt.DefaultCost)
		if err1 != nil {
			responseParser.JsonInternalError(c,"密码加密失败",err1)
			return
		}
		platEnt.EnterprisePassword = string(hash)
		platEnt.EnterpriseID = snowflake.GetSnowflakeID()
		platEnt.EnterprisePhone = parser.EnterprisePhone
		platEnt.Location = parser.Location

		err1 = platEnt.Add()
		if err1 != nil {
			responseParser.JsonDBError(c,"",err1)
			return
		}
	}else {
		responseParser.JsonDBError(c,"",err)
		return
	}
	c.JSON(200,gin.H{
		"code":codes.OK,
		"message":"注册成功",
	})
}

type changePwdParser struct {
	EnterpriseID 	string	`form:"EnterpriseID" json:"EnterpriseID" binding:"required"`
	NewPassword 	string	`form:"NewPassword" json:"NewPassword" binding:"required"`
}

//平台企业密码仅可通过管理员来进行修改
func (*PlatEntApiImpl) ChangePwd(c *gin.Context) {
	var parser changePwdParser
	err := c.ShouldBind(&parser)
	if err != nil {
		responseParser.JsonParameterIllegal(c,err)
		return
	}

	temp,ok := c.Get("user")
	if !ok {
		responseParser.JsonLoginError(c,"用户未登录",nil)
		return
	}
	user,_ := temp.(ginModels.UserModel)
	if user.IsPlatUser == false{
		responseParser.JsonAccessDenied(c,"仅允许平台修改密码")
		return
	}

	enterpriseInfo := services.PlatEnterpriseService{}
	enterpriseInfo.EnterpriseID = parser.EnterpriseID
	err = enterpriseInfo.Get()
	if err != nil {
		if err.Error() == "record not found" {
			responseParser.JsonNotData(c,err)
			return
		}
		responseParser.JsonDBError(c,"",err)
		return
	}
	hash,err := bcrypt.GenerateFromPassword([]byte(parser.NewPassword),bcrypt.DefaultCost)
	if err != nil {
		responseParser.JsonInternalError(c,"密码加密失败",err)
		return
	}
	err = enterpriseInfo.Update(map[string]interface{}{
		"EnterprisePassword":hash,
	})
	if err != nil {
		responseParser.JsonDBError(c,"数据库密码更新失败",err)
		return
	}
	responseParser.JsonOK(c,"密码更新成功")
}