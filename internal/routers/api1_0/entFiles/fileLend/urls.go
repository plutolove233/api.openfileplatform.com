/*
@Coding : utf-8
@Time : 2022/2/15 17:19
@Author : 刘浩宇
@Software: GoLand
*/
package fileLend

import (
	"api.openfileplatform.com/internal/apis/api1_0/entFileLend"
	"api.openfileplatform.com/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func InitFileLendRouterGroup(engine *gin.RouterGroup){
	Api := engine.Group("lend")
	initBaseUrls(Api)

	var fileLendApi entFileLend.EnterpriseFileLendApi
	Api.Use(middlewares.TokenRequire())
	Api.POST("borrow",middlewares.AuthenticationMiddleware(),fileLendApi.BorrowFile)
	Api.POST("return",middlewares.AuthenticationMiddleware(),fileLendApi.ReturnFile)
}