// coding: utf-8
// @Author : lryself
// @Date : 2022/2/9 15:12
// @Software: GoLand

package platUsers

import (
	"api.openfileplatform.com/internal/apis/api1_0/platUsers"
	"github.com/gin-gonic/gin"
)

var (
	Api *gin.RouterGroup
)

func InitPlatUsersRouterGroup(engine *gin.RouterGroup) {
	Api = engine.Group("platUsers")
	initBaseAPI()

	var platformUserApi platUsers.PlatformUserApi
	Api.POST("refresh", platformUserApi.RefreshPassword)
	Api.POST("changePassword", platformUserApi.ChangePassword)
	Api.POST("register", platformUserApi.Register)
}
