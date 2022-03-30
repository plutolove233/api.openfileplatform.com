/*
@Coding : utf-8
@Time : 2022/2/15 17:16
@Author : 刘浩宇
@Software: GoLand
*/
package fileType

import (
	"api.openfileplatform.com/internal/apis/api1_0/entFileType"
	"api.openfileplatform.com/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func InitFileTypeRouterGroup(engine *gin.RouterGroup){
	Api := engine.Group("type")
	initBaseUrls(Api)

	var FileTypeApi entFileType.EnterpriseFileTypeApi
	Api.Use(middlewares.TokenRequire())
	Api.GET("getAll",FileTypeApi.GetAllFileType)
}