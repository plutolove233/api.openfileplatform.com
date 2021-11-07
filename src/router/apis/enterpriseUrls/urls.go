package enterprise

import (
	"api.openfileplatform.com/api/enterprise"
	"github.com/gin-gonic/gin"
)

func InitEnterpriseApiGroup(r *gin.Engine){
	ent := r.Group("enterprise")
	{
		user := ent.Group("user")
		{
			user.POST("login",enterprise.UserLogin)
			user.POST("register",enterprise.UserRegister)
			user.POST("list",enterprise.GetUsers)
			file := user.Group("file")
			{
				file.POST("upload",enterprise.Upload)
				file.POST("borrow/:id",enterprise.BorrowFile)
				file.POST("return/:id",enterprise.ReturnFile)
				file.DELETE("delete/:id",enterprise.DeleteFile)
			}
			user.PUT("logo",enterprise.ChangeFace)
		}
		role:=ent.Group("role")
		{
			role.POST("new",enterprise.NewRole)
			role.GET("list",enterprise.GetRoleList)
			role.DELETE("delete/:id",enterprise.DeleteRole)
		}
		department := ent.Group("department")
		{
			department.POST("new",enterprise.NewDepartment)
			department.DELETE("delete/:id",enterprise.DeleteDepartment)
		}
	}
}