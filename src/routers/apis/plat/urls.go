package plat

import (
	"api.openfileplatform.com/src/apis/platform"
	"api.openfileplatform.com/src/routers/apis/plat/users"
	"github.com/gin-gonic/gin"
)

var (
	Api *gin.RouterGroup
)

func InitPlatRouterGroup(engine *gin.RouterGroup) {
	Api = engine.Group("plat")

	var loginApi platform.LoginApiImpl
	Api.POST("/login", loginApi.LoginByPassword)
	Api.POST("/changePassword", loginApi.ChangePassword)
	Api.POST("/logout", loginApi.Logout)
	Api.POST("/refreshToken", loginApi.RefreshToken)

	users.InitUsersRouterGroup(Api)
}
