/*
@Coding : utf-8
@Time : 2022/2/12 14:51
@Author : 刘浩宇
@Software: GoLand
*/
package enterpriseUsers

import (
	"api.openfileplatform.com/internal/apis/api1_0/entUsers"
	"api.openfileplatform.com/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func InitEnterpriseUsersRouterGroup(engine *gin.RouterGroup){
	Api := engine.Group("entusers")
	initBaseApi(Api)

	var entUserImpl entUsers.EnterpriseUserApi
	Api.POST("register",entUserImpl.Register)
	Api.Use(middlewares.TokenRequire())
	Api.POST("refreshPassword",entUserImpl.RefreshPassword)
	Api.POST("changePassword",entUserImpl.ChangePassword)
	Api.GET("getUsers",entUserImpl.GetAllUsersList)
	Api.POST("setAdmin",middlewares.AuthenticationMiddleware(),entUserImpl.SetAdmin)
	Api.PUT("removeAdmin",middlewares.AuthenticationMiddleware(),entUserImpl.RemoveAdmin)
	Api.GET("getLendInfo",entUserImpl.GetUserLendInformation)
}