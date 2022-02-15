/*
@Coding : utf-8
@Time : 2022/2/14 15:06
@Author : 刘浩宇
@Software: GoLand
*/
package entFiles

import (
	"api.openfileplatform.com/internal/globals/codes"
	"api.openfileplatform.com/internal/globals/responseParser"
	"api.openfileplatform.com/internal/models/ginModels"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (*EnterpriseFilesApi)GetAllEntFiles(c *gin.Context){
	temp,ok := c.Get("user")
	if !ok {
		responseParser.JsonNotData(c,"用户未登录",nil)
		return
	}

	user := temp.(ginModels.UserModel)

	if user.IsAdmin == false {
		c.JSON(http.StatusOK,gin.H{
			"code":codes.UnauthorizedUserId,
			"message":"用户没有权限查看企业文档",
		})
		return
	}
}