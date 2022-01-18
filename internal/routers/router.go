package routers

import (
	"api.openfileplatform.com/internal/routers/apis"
	"github.com/gin-gonic/gin"
)

func InitRouter(engine *gin.Engine) {
	apis.InitApiGroup(engine)
}
