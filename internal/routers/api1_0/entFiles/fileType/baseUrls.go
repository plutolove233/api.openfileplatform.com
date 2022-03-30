/*
@Coding : utf-8
@Time : 2022/2/15 17:16
@Author : 刘浩宇
@Software: GoLand
*/
package fileType

import (
	"api.openfileplatform.com/internal/apis/api1_0/entFileType"
	"github.com/gin-gonic/gin"
)

func initBaseUrls(engine *gin.RouterGroup) {
	var impl entFileType.EnterpriseFileTypeApi

	engine.Any("", impl.EnterpriseFileTypeApi)
}