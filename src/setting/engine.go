package setting

import (
	"api.openfileplatform.com/src/routers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitEngine() (*gin.Engine, error) {
	//新建路由引擎
	r := gin.Default()
	r.Use(cors.Default())

	//导入static以及template模板
	//	r.Static("/static","static")
	//	r.LoadHTMLGlob("templates/*")

	//路由配置
	routers.InitRouter(r)

	return r, nil
}
