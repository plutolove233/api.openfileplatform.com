/*
@Coding : utf-8
@Time : 2022/2/15 16:59
@Author : 刘浩宇
@Software: GoLand
*/
package entFiles

import (
	"api.openfileplatform.com/internal/apis/api1_0/entFiles"
	"github.com/gin-gonic/gin"
)

func initBaseApi(engine *gin.RouterGroup) {
	var impl entFiles.EnterpriseFilesApi
	engine.Any("", impl.EnterpriseFilesApi)
}