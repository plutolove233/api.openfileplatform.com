// coding:utf-8
// @Author:PigKnight
// @Date:2022/2/10 17:10
// @Software: GoLand

package platEnterprises

import (
	"api.openfileplatform.com/internal/apis/api1_0/platEnterprises"
)

func initBaseAPI() {
	var impl platEnterprises.PlatformEnterpriseApi
	Api.Any("", impl.PlatformEnterpriseApi)
}
