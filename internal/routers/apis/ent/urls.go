package ent

import (
	"api.openfileplatform.com/internal/apis/api1_0"
	"github.com/gin-gonic/gin"
)

func InitEntRouterGroup(engine *gin.RouterGroup) {
	Api := engine.Group("ent")
	var loginApi api1_0.UserApi
	Api.POST("/login", loginApi.LoginByPassword)
}
