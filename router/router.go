package router

import (
	"github.com/diguacheng/mycaptcha"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init(){
	// 导入字体
	mycaptcha.LoadFonts("router/fonts")
}

func StartEngine()(r *gin.Engine){
	r = gin.Default()
	r.Use(cors.Default())



	r.GET("verification", func(c *gin.Context) {
		base64image,_ := mycaptcha.GetCaptchaBase64(300,100,4)
		c.JSON(200,gin.H{
			"code":200,
			"verification":base64image,
			//"ans":key,
		})
	})

	return
}