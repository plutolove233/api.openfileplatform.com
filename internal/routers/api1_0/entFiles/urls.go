/*
@Coding : utf-8
@Time : 2022/2/15 16:59
@Author : 刘浩宇
@Software: GoLand
*/
package entFiles

import (
	"api.openfileplatform.com/internal/routers/api1_0/entFiles/entFileCategory"
	"api.openfileplatform.com/internal/routers/api1_0/entFiles/fileLend"
	"api.openfileplatform.com/internal/routers/api1_0/entFiles/fileType"
	"github.com/gin-gonic/gin"
)

func InitEnterpriseFileRouterGroup(engine *gin.RouterGroup){
	Api := engine.Group("file")

	initBaseApi(Api)

	entFileCategory.InitEnterpriseFileCategoryRouterGroup(Api)
	fileType.InitFileTypeRouterGroup(Api)
	fileLend.InitFileLendRouterGroup(Api)
}