/*
@Coding : utf-8
@Time : 2022/2/12 15:00
@Author : 刘浩宇
@Software: GoLand
*/
package enterpriseUsers

import (
	"api.openfileplatform.com/internal/apis/api1_0/entUsers"
	"github.com/gin-gonic/gin"
)

func initBaseApi(engine *gin.RouterGroup) {
	var impl entUsers.EnterpriseUserApi

	engine.Any("",impl.EnterpriseUserApi)
}