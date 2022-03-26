/*
@Coding : utf-8
@Time : 2022/2/15 17:21
@Author : 刘浩宇
@Software: GoLand
*/
package entProject

import (
	"api.openfileplatform.com/internal/apis/api1_0/entProject"
	"api.openfileplatform.com/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func InitEnterpriseProjectRouterGroup(engine *gin.RouterGroup) {
	Api := engine.Group("project")
	initBaseUrls(Api)

	var projectApi entProject.EnterpriseProjectApi
	Api.Use(middlewares.TokenRequire())
	Api.POST("add", middlewares.AuthenticationMiddleware(),projectApi.AddNewProject)
	Api.PUT("changeName",middlewares.AuthenticationMiddleware(),projectApi.ChangeProjectName)
	Api.GET("all",projectApi.GetAllProject)
}