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
		enterprise := plat.Group("enterprise")
		{
			enterprise.PUT("new",platform.NewEnterprise)
			enterprise.DELETE("delete/:id",platform.DeleteEnterprise)
			enterprise.GET("list",platform.GetEnterpriseList)
			enterprise.PUT("logo",platform.ChangeLogo)
			//enterprise.PUT("change_pwd",platform.ChangePWD)
		}
		module := plat.Group("module")
		{
			module.PUT("new",platform.NewFunctionModule)
			module.DELETE("delete/:id",platform.DeleteFunctionModule)
			module.GET("list",platform.GetModuleList)
			module.PUT("find",platform.FindModule)
		}
	}
}