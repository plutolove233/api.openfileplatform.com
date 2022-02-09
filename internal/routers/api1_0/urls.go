package api1_0

import (
	"api.openfileplatform.com/internal/apis/api1_0"
	"api.openfileplatform.com/internal/routers/api1_0/platUsers"
	"github.com/gin-gonic/gin"
)

var (
	Api *gin.RouterGroup
)

func InitAPI1_0Router(engine *gin.Engine) {
	Api = engine.Group("api1_0")
	Api.GET("version", api1_0.GetVersion)

	platUsers.InitPlatUsersRouterGroup(Api)
}
