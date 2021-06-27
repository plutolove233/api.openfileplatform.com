package platform

import (
	"api.openfileplatform.com/api/platform"
	"github.com/gin-gonic/gin"
)

func InitPlatformApiGroup(r *gin.Engine){
	plat := r.Group("plat")
	{
		plat.GET("list",platform.PlatGetUserList)
		plat.POST("login",platform.PlatUserLogin)
		plat.POST("register",platform.PlatUserRegister)
		plat.POST("reset/pwd",platform.PlatResetPwd)
		plat.POST("reset/phone",platform.PlatResetPhone)
		plat.POST("reset/email",platform.PlatResetEmail)
	}
}