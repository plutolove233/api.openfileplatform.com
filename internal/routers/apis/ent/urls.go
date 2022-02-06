package ent

import "github.com/gin-gonic/gin"

func InitEntRouterGroup(engine *gin.RouterGroup){
	Api := engine.Group("ent")
	Api.POST("/login",)
}
