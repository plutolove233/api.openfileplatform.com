// coding: utf-8
// @Author : lryself
// @Date : 2022/2/9 15:15
// @Software: GoLand

package platUsers

import (
	"api.openfileplatform.com/internal/apis/api1_0/platUsers"
)

func initBaseAPI() {
	var impl platUsers.PlatformUserApi
	Api.Any("", impl.PlatformUserApi)
}
