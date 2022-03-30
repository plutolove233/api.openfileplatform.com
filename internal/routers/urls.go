package routers

import (
	"api.openfileplatform.com/internal/routers/api1_0"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func InitStaticRouterGroup(engine *gin.Engine){
	Api := engine.Group("statics")
	Api.GET("/ping", func(c *gin.Context) {
		c.JSON(200,"pong")
	})
	Api.StaticFS("/save",gin.Dir(viper.GetString("system.FsPath"),false))
	Api.StaticFS("/saveList",gin.Dir(viper.GetString("system.FsPath"),true))
	Api.StaticFS("/log",gin.Dir("./logs",true))
}

func InitRouter(engine *gin.Engine) {
	api1_0.InitAPI1_0Router(engine)
}
