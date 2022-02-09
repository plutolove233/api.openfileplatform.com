package ent

import (
	"api.openfileplatform.com/internal/apis"
	"github.com/gin-gonic/gin"
)

func InitEntRouterGroup(engine *gin.RouterGroup){
	Api := engine.Group("ent")
	var loginApi apis.UserApiImpl
	Api.POST("/login",loginApi.LoginByPassword)
}
