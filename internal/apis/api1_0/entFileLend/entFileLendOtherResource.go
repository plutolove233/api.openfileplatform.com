/*
@Coding : utf-8
@Time : 2022/3/18 15:00
@Author : 刘浩宇
@Software: GoLand
*/
package entFileLend

import (
	"api.openfileplatform.com/internal/globals/responseParser"
	"api.openfileplatform.com/internal/models/ginModels"
	"api.openfileplatform.com/internal/services"
	"github.com/gin-gonic/gin"
)

type borrowFileParser struct {
	FileID		string 	`form:"FileID" json:"FileID" binding:"required"`
	BorrowTerm 	int8 	`form:"BorrowTerm" json:"BorrowTerm" binding:"required"`
}

func (*EnterpriseFileLendApi) BorrowFile(c *gin.Context) {
	var parser borrowFileParser
	err := c.ShouldBind(&parser)
	if err != nil {
		responseParser.JsonParameterIllegal(c,"获取借阅信息失败",err)
		return
	}

	temp,ok := c.Get("user")
	if !ok {
		responseParser.JsonNotData(c,"用户未登录",nil)
		return
	}

	user := temp.(ginModels.UserModel)
	entUserService := services.EntUserService{}
	entUserService.UserID = user.UserID
	err = entUserService.Get()
	if err != nil {
		responseParser.JsonNotData(c,"该用户信息不存在",err)
		return
	}

	info, err := entUserService.BorrowFile(parser.FileID,parser.BorrowTerm)
	if err != nil {
		if err.Error() == "Enterprise is not matched" {
			responseParser.JsonAccessDenied(c,"文件所属企业与用户所属企业不匹配")
			return
		}
		if err.Error() == "This book has been borrowed" {
			responseParser.JsonNotData(c,"该书已经被借出",err)
			return
		}
		responseParser.JsonDBError(c,"借阅图书失败",err)
		return
	}

	responseParser.JsonOK(c,"图书借阅成功",info)
}