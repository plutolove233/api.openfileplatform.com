package apis

import (
	"api.openfileplatform.com/api/common"
	enterpriseApi "api.openfileplatform.com/router/apis/enterpriseUrls"
	platformApi "api.openfileplatform.com/router/apis/platformUrls"
	"github.com/gin-gonic/gin"
)

func InitApiGroup(r *gin.Engine){
	enterpriseApi.InitEnterpriseApiGroup(r)
	platformApi.InitPlatformApiGroup(r)
	r.POST("/email", common.SMTPSendEmail)
	r.GET("api/version", common.Version)
}