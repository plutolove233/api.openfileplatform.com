package platform

import (
	"DocumentSystem/api/platform"
	"github.com/gin-gonic/gin"
)

func InitPlatformApiGroup(r *gin.Engine){
	plat := r.Group("plat")
	{
		plat.GET("list",platform.PlatGetUserList)
		plat.POST("login",platform.PlatUserLogin)
		plat.POST("register",platform.PlatUserRegister)
	}
}