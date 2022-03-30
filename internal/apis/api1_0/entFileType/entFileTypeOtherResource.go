/*
@Coding : utf-8
@Time : 2022/3/25 17:12
@Author : 刘浩宇
@Software: GoLand
*/
package entFileType

import (
	"api.openfileplatform.com/internal/globals/responseParser"
	"api.openfileplatform.com/internal/services"
	"github.com/gin-gonic/gin"
)

type fileTypeParser struct {
	FileTypeID   string `json:"FileTypeID"`
	FileTypeName string `json:"FileTypeName"`
}

func (*EnterpriseFileTypeApi) GetAllFileType(c *gin.Context) {
	_,ok := c.Get("user")
	if !ok {
		responseParser.JsonNotData(c,"用户未登录",nil)
		return
	}

	fileTypeService := services.EnterpriseFileTypeService{}
	fileTypeinfo,err1 := fileTypeService.GetAll()
	if err1 != nil {
		responseParser.JsonDBError(c,"获取所有分类失败",err1)
		return
	}
	data := []fileTypeParser{}
	for _,item := range fileTypeinfo {
		x := fileTypeParser{
			FileTypeID: item.FileTypeID,
			FileTypeName: item.FileTypeName,
		}
		data = append(data, x)
	}
	responseParser.JsonOK(c,"获取所有分类成功",data)
}