// coding:utf-8
// @Author:PigKnight
// @Date:2022/2/10 16:50
// @Software: GoLand

package platEnterprises

import (
	"github.com/gin-gonic/gin"
)

var (
	Api *gin.RouterGroup
)

func InitPlatEnterprisesRouterGroup(engine *gin.RouterGroup) {
	Api = engine.Group("platEnterprises")
	initBaseAPI()

	//var platformEnterpriseApi platEnterprises.PlatformEnterpriseApi
}
