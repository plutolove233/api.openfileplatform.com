package apis

import (
	enterpriseApi "DocumentSystem/router/apis/enterpriseUrls"
	platformApi "DocumentSystem/router/apis/platformUrls"
	"github.com/gin-gonic/gin"
)

func InitApiGroup(r *gin.Engine){
	enterpriseApi.InitEnterpriseApiGroup(r)
	platformApi.InitPlatformApiGroup(r)
}