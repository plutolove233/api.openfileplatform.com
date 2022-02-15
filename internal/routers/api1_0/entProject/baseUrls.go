/*
@Coding : utf-8
@Time : 2022/2/15 17:21
@Author : 刘浩宇
@Software: GoLand
*/
package entProject

import (
	"api.openfileplatform.com/internal/apis/api1_0/entProject"
	"github.com/gin-gonic/gin"
)

func initBaseUrls(engine *gin.RouterGroup) {
	var impl entProject.EnterpriseProjectApi

	engine.Any("",impl.EnterpriseProjectApi)
}