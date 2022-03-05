/*
@Coding : utf-8
@Time : 2022/2/14 17:21
@Author : 刘浩宇
@Software: GoLand
*/
package entFileCategory

import (
	"api.openfileplatform.com/internal/apis/api1_0/entFileCategory"
	"api.openfileplatform.com/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func InitEnterpriseFileCategoryRouterGroup(engine *gin.RouterGroup){
	Api := engine.Group("category")
	initEnterpriseFileCategory(Api)

	var fileCategoryApi entFileCategory.EnterpriseFileCategoryApi
	Api.POST("path",fileCategoryApi.GetFileCategoryPath)
	Api.Use(middlewares.TokenRequire())
	Api.POST("add",fileCategoryApi.AddFileCategory)
}