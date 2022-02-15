/*
@Coding : utf-8
@Time : 2022/2/15 17:16
@Author : 刘浩宇
@Software: GoLand
*/
package fileType

import "github.com/gin-gonic/gin"

func InitFileTypeRouterGroup(engine *gin.RouterGroup){
	Api := engine.Group("type")
	initBaseUrls(Api)
}