/*
@Coding : utf-8
@Time : 2022/2/14 16:07
@Author : 刘浩宇
@Software: GoLand
*/
package entFileCategory

import (
	"api.openfileplatform.com/internal/apis/api1_0/entFileCategory"
	"github.com/gin-gonic/gin"
)

func initEnterpriseFileCategory(engine *gin.RouterGroup) {
	var impl entFileCategory.EnterpriseFileCategoryApi

	engine.Any("",impl.EnterpriseFileCategoryApi)
}