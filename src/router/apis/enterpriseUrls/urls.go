package enterprise

import (
	"api.openfileplatform.com/api/enterprise"
	"github.com/gin-gonic/gin"
)

func InitEnterpriseApiGroup(r *gin.Engine){
	ent := r.Group("enterprise")
	{
		file := ent.Group("file")
		{
			file.POST("upload",enterprise.Upload)
			file.POST("borrow/:id",enterprise.BorrowFile)
			file.POST("return/:id",enterprise.ReturnFile)
		}
	}
}