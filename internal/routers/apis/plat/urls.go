package plat

import (
	"api.openfileplatform.com/internal/apis"
	"api.openfileplatform.com/internal/routers/apis/plat/users"
	"github.com/gin-gonic/gin"
)

var (
	Api *gin.RouterGroup
)

func InitPlatRouterGroup(engine *gin.RouterGroup) {
	Api = engine.Group("plat")

	var loginApi apis.UserApiImpl
	Api.POST("/login", loginApi.LoginByPassword)

	//Api.POST("/logout", loginApi.Logout)
	Api.POST("/refreshToken", loginApi.RefreshToken)

	users.InitUsersRouterGroup(Api)
}
