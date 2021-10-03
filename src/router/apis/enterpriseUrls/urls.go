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
				file.POST("borrow/:id",jwt.JWTAuthMiddleware(),enterprise.BorrowFile)
				file.POST("return/:id",enterprise.ReturnFile)
				file.DELETE("delete/:id",jwt.JWTAuthMiddleware(),enterprise.DeleteFile)
			}
			user.PUT("logo",enterprise.ChangeFace)
		}
		role:=ent.Group("role")
		{
			role.POST("new",jwt.JWTAuthMiddleware(),enterprise.NewRole)
			role.GET("list",jwt.JWTAuthMiddleware(),enterprise.GetRoleList)
			role.DELETE("delete/:id",jwt.JWTAuthMiddleware(),enterprise.DeleteRole)
		}
		department := ent.Group("department")
		{
			department.POST("new",jwt.JWTAuthMiddleware(),enterprise.NewDepartment)
			department.DELETE("delete/:id",jwt.JWTAuthMiddleware(),enterprise.DeleteDepartment)
		}
	}
}