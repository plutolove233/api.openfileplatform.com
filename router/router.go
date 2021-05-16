package router

import (
	"DocumentSystem/api/normal"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func StartEngine()(r *gin.Engine){
	r = gin.Default()
	r.Use(cors.Default())

	_normal := r.Group("normal")
	{
		_normal.POST("login",normal.Login)
		_normal.POST("register",normal.Register)
		_normal.GET("view",normal.View)
	}

	return
}