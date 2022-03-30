/*
@Coding : utf-8
@Time : 2022/2/15 17:19
@Author : 刘浩宇
@Software: GoLand
*/
package fileLend

import (
	"api.openfileplatform.com/internal/apis/api1_0/entFileLend"
	"github.com/gin-gonic/gin"
)

func initBaseUrls(engine *gin.RouterGroup) {
	var impl entFileLend.EnterpriseFileLendApi

	engine.Any("", impl.EnterpriseFileLendApi)
}