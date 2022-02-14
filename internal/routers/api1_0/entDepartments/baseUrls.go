/*
@Coding : utf-8
@Time : 2022/2/14 15:37
@Author : 刘浩宇
@Software: GoLand
*/
package entDepartments

import (
	"api.openfileplatform.com/internal/apis/api1_0/entFiles"
	"github.com/gin-gonic/gin"
)

func initDepartmentsRouterGroup(engine *gin.RouterGroup){
	var impl entFiles.EntFilesApi
	engine.Any("",impl.EntFilesApi)
}