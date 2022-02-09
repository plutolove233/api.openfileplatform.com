// coding: utf-8
// @Author : lryself
// @Date : 2022/2/9 15:12
// @Software: GoLand

package platUsers

import (
	"api.openfileplatform.com/internal/apis/api1_0/platUsers"
	"api.openfileplatform.com/internal/middlewares"
	"github.com/gin-gonic/gin"
)

var (
	Api *gin.RouterGroup
)

func InitPlatUsersRouterGroup(engine *gin.RouterGroup) {
	Api = engine.Group("platUsers")
	initBaseAPI()

	var platformUserApi platUsers.PlatformUserApi
	Api.POST("register", platformUserApi.Register)
	Api.Use(middlewares.TokenRequire())
	Api.POST("refreshPassword", platformUserApi.RefreshPassword)
	Api.POST("changePassword", platformUserApi.ChangePassword)
}
