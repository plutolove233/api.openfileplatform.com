package routers

import (
	"api.openfileplatform.com/internal/routers/api1_0"
	"github.com/gin-gonic/gin"
)

func InitRouter(engine *gin.Engine) {
	api1_0.InitAPI1_0Router(engine)
}
