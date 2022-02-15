package api1_0

import (
	"api.openfileplatform.com/internal/apis/api1_0"
	"api.openfileplatform.com/internal/routers/api1_0/entDepartments"
	"api.openfileplatform.com/internal/routers/api1_0/entFiles"
	"api.openfileplatform.com/internal/routers/api1_0/entProject"
	"api.openfileplatform.com/internal/routers/api1_0/enterpriseUsers"
	"api.openfileplatform.com/internal/routers/api1_0/platEnterprises"
	"api.openfileplatform.com/internal/routers/api1_0/platUsers"
	"github.com/gin-gonic/gin"
)

var (
	Api *gin.RouterGroup
)

func InitAPI1_0Router(engine *gin.Engine) {
	Api = engine.Group("api1_0")
	Api.GET("version", api1_0.GetVersion)

	platUsers.InitPlatUsersRouterGroup(Api)
	var userApi api1_0.UserApi
	Api.POST("loginByPassword", userApi.LoginByPassword)
	Api.POST("refreshToken", userApi.RefreshToken)

	platEnterprises.InitPlatEnterprisesRouterGroup(Api)
	enterpriseUsers.InitEnterpriseUsersRouterGroup(Api)
	entProject.InitEnterpriseProjectRouterGroup(Api)
	entDepartments.InitEntDepartmentsRouterGroup(Api)
	entFiles.InitEnterpriseFileRouterGroup(Api)
}
