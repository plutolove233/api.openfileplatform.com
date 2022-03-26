/*
@Coding : utf-8
@Time : 2022/2/15 17:28
@Author : 刘浩宇
@Software: GoLand
*/
package entDepartments

import (
	"api.openfileplatform.com/internal/apis/api1_0/entDepartments"
	"api.openfileplatform.com/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func InitEntDepartmentsRouterGroup(engine *gin.RouterGroup){
	Api := engine.Group("department")
	initDepartmentsRouterGroup(Api)

	var departmentApi entDepartments.EnterpriseDepartmentApi
	Api.Use(middlewares.TokenRequire())
	Api.POST("setHeader",middlewares.AuthenticationMiddleware(),departmentApi.SetHeader)
	Api.GET("getAll",departmentApi.GetAllDepartment)
}