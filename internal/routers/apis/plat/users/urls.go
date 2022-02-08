// coding: utf-8
// @Author : lryself
// @Date : 2021/12/28 21:26
// @Software: GoLand

package users

import (
	"api.openfileplatform.com/internal/apis/platform/userResource"
	"api.openfileplatform.com/internal/middlewares"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	Api *gin.RouterGroup
)

func InitUsersRouterGroup(engine *gin.RouterGroup) {
	Api = engine.Group("user")

	var userApi userResource.UserApiImpl
	Api.POST("register", userApi.Register)

	Api.Use(middlewares.PlatformTokenRequire())
	Api.POST("/changePassword", userApi.ChangePassword)
	Api.GET("ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, "pong")
	})
	Api.GET("list",userApi.GetUserList)
}
