package entResource

import (
	"api.openfileplatform.com/internal/globals/codes"
	"api.openfileplatform.com/internal/globals/responseParser"
	"api.openfileplatform.com/internal/globals/snowflake"
	"api.openfileplatform.com/internal/services"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type PlatEntApiImpl struct {}

type EntRegisterParser struct {
	EnterpriseName 	string	`form:"EnterpriseName" json:"EnterpriseName" binding:"required"`
	EnterprisePwd 	string	`form:"EnterprisePwd" json:"EnterprisePwd" binding:"required"`
	Location 		string	`form:"Location" json:"Location" binding:"required"`
	EnterprisePhone string	`form:"EnterprisePhone" json:"EnterprisePhone" binding:"required"`
}

//用于在平台注册企业
func ( *EntRegisterParser) Register(c *gin.Context)  {
	var parser EntRegisterParser
	err := c.ShouldBindJSON(&parser)
	if err != nil {
		responseParser.JsonParameterIllegal(c,err)
		return
	}

	platEnt := services.PlatEnterpriseService{}
	platEnt.EnterpriseName = parser.EnterpriseName
	err = platEnt.Get()
	if err == nil {
		c.JSON(200,gin.H{
			"code":codes.DataExist,
			"message":"企业名称已存在",
		})
		return
	}else if err.Error() == "record not found" {
		hash, err1 := bcrypt.GenerateFromPassword([]byte(parser.EnterprisePwd),bcrypt.DefaultCost)
		if err1 != nil {
			c.JSON(200,gin.H{
				"code":codes.InternalError,
				"message":"密码加密失败",
				"error":err.Error(),
			})
			return
		}
		platEnt.EnterprisePwd = string(hash)
		platEnt.EnterpriseID = snowflake.GetSnowflakeID()
		platEnt.EnterprisePhone = parser.EnterprisePhone
		platEnt.Location = parser.Location

		err1 = platEnt.Add()
		if err1 != nil {
			responseParser.JsonDBError(c,err1)
			return
		}
	}else {
		responseParser.JsonDBError(c,err)
		return
	}
	c.JSON(200,gin.H{
		"code":codes.OK,
		"message":"注册成功",
	})
}