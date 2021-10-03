package jwt

import (
	"api.openfileplatform.com/commons/codes"
	"api.openfileplatform.com/utils/authority"
	"github.com/gin-gonic/gin"
)

func AuthorityMiddleware(permissionID int64) func (ctx gin.Context){
	return func(ctx gin.Context) {
		temp,ok := ctx.Get("UserID")
		if ok == false{
			ctx.JSON(200,gin.H{
				"code":codes.InternetError,
				"msg":"用户信息获取失败",
			})
			return
		}
		userID := temp.(int64)
		if !authority.CheckAuthority(userID, permissionID) {
			ctx.JSON(200,gin.H{
				"code":codes.UserError,
				"msg":"用户没有权限访问",
			})
			ctx.Abort()
			return
		}
		ctx.Next()
		return
	}
}