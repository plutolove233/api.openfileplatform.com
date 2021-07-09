package enterprise

import (
	"api.openfileplatform.com/api/enterprise"
	"api.openfileplatform.com/utils/jwt"
	"github.com/gin-gonic/gin"
)

func InitEnterpriseApiGroup(r *gin.Engine){
	ent := r.Group("enterprise")
	{
		user := ent.Group("user")
		{
			user.POST("login",enterprise.UserLogin)
			user.POST("register",enterprise.UserRegister)
			user.GET("list",enterprise.GetUsers)
			file := user.Group("file")
			{
				file.POST("upload",jwt.JWTAuthMiddleware(),enterprise.Upload)
				file.POST("borrow/:id",enterprise.BorrowFile)
				file.POST("return/:id",enterprise.ReturnFile)
			}
		}
	}
}