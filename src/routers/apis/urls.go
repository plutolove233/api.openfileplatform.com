package apis

import (
	"api.openfileplatform.com/src/apis"
	"api.openfileplatform.com/src/routers/apis/plat"
	"github.com/gin-gonic/gin"
)

var (
	Api *gin.RouterGroup
)

func InitApiGroup(engine *gin.Engine) {
	Api = engine.Group("api")
	Api.GET("version", apis.GetVersion)

	plat.InitPlatRouterGroup(Api)
}
