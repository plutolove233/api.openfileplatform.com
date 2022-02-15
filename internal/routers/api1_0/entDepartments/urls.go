/*
@Coding : utf-8
@Time : 2022/2/15 17:28
@Author : 刘浩宇
@Software: GoLand
*/
package entDepartments

import "github.com/gin-gonic/gin"

func InitEntDepartmentsRouterGroup(engine *gin.RouterGroup){
	Api := engine.Group("department")
	initDepartmentsRouterGroup(Api)
}