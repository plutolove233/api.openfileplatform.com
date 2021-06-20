package apis

import (
	"DocumentSystem/api"
	"DocumentSystem/api/common"
	enterpriseApi "DocumentSystem/router/apis/enterpriseUrls"
	platformApi "DocumentSystem/router/apis/platformUrls"
	"github.com/gin-gonic/gin"
)

func InitApiGroup(r *gin.Engine){
	enterpriseApi.InitEnterpriseApiGroup(r)
	platformApi.InitPlatformApiGroup(r)
	r.POST("/email",api.SMTPSendEmail)
	r.GET("api/version", common.Version)
}