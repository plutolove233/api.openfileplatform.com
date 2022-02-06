package apis

import (
	"api.openfileplatform.com/internal/apis"
	"api.openfileplatform.com/internal/routers/apis/ent"
	"api.openfileplatform.com/internal/routers/apis/plat"
	"github.com/gin-gonic/gin"
)

var (
	Api *gin.RouterGroup
)

func InitApiGroup(engine *gin.Engine) {
	Api = engine.Group("api")
	Api.GET("version", apis.GetVersion)

	plat.InitPlatRouterGroup(Api)
	ent.InitEntRouterGroup(Api)
}
