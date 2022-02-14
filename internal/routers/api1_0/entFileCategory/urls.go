/*
@Coding : utf-8
@Time : 2022/2/14 17:21
@Author : 刘浩宇
@Software: GoLand
*/
package entFileCategory

import "github.com/gin-gonic/gin"

func InitEnterpriseFileCategoryRouterGroup(engine *gin.RouterGroup){
	Api := engine.Group("category")
	initEnterpriseFileCategory(Api)
}