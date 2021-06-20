package router

import (
	"api.openfileplatform.com/router/apis"
	"github.com/diguacheng/mycaptcha"
	"github.com/gin-gonic/gin"
)

func init(){
	// 导入字体
	mycaptcha.LoadFonts("router/fonts")
}

func InitRouter(engine *gin.Engine){
	apis.InitApiGroup(engine)
	engine.GET("/", func(c *gin.Context) {
		c.HTML(200,"index.html",nil)
	})
}