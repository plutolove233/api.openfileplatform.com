/*
@Coding : utf-8
@Time : 2022/2/15 17:19
@Author : 刘浩宇
@Software: GoLand
*/
package fileLend

import "github.com/gin-gonic/gin"

func InitFileLendRouterGroup(engine *gin.RouterGroup){
	Api := engine.Group("lend")
	initBaseUrls(Api)
}