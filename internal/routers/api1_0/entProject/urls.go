/*
@Coding : utf-8
@Time : 2022/2/15 17:21
@Author : 刘浩宇
@Software: GoLand
*/
package entProject

import "github.com/gin-gonic/gin"

func InitEnterpriseProjectRouterGroup(engine *gin.RouterGroup) {
	Api := engine.Group("project")
	initBaseUrls(Api)
}