// coding:utf-8
// @Author:PigKnight
// @Date:2022/2/10 16:50
// @Software: GoLand

package platEnterprises

import (
	"api.openfileplatform.com/internal/apis/api1_0/platEnterprises"
	"api.openfileplatform.com/internal/middlewares"
	"github.com/gin-gonic/gin"
)

var (
	Api *gin.RouterGroup
)

func InitPlatEnterprisesRouterGroup(engine *gin.RouterGroup) {
	Api = engine.Group("platEnterprises")
	initBaseAPI()

	var platformEnterpriseApi platEnterprises.PlatformEnterpriseApi
	Api.POST("register",platformEnterpriseApi.Register)
	Api.GET("getAll",platformEnterpriseApi.GetAll)
	Api.POST("getAllUsers",platformEnterpriseApi.GetAllUsers)
	Api.Use(middlewares.TokenRequire())
	Api.POST("refreshPassword",platformEnterpriseApi.RefreshPassword)
	Api.POST("changePassword",platformEnterpriseApi.ChangePassword)
}
