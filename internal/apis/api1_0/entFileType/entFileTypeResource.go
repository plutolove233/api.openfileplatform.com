/*
@Coding : utf-8
@Time : 2022/2/15 16:30
@Author : 刘浩宇
@Software: GoLand
*/
package entFileType

import (
	"api.openfileplatform.com/internal/globals/responseParser"
	"api.openfileplatform.com/internal/services"
	"api.openfileplatform.com/internal/utils/structs"
	"github.com/gin-gonic/gin"
)

type EnterpriseFileTypeApi struct {}

func (*EnterpriseFileTypeApi) EnterpriseFileTypeApi(c *gin.Context) {
	var err error
	var entFileTypeService services.EnterpriseFileTypeService
	err = c.ShouldBind(&entFileTypeService)

	if c.Request.Method == "GET" {
		err = entFileTypeService.Get()
		if err != nil {
			responseParser.JsonDBError(c, "", err)
			return
		}
	} else if c.Request.Method == "POST" {
		err = entFileTypeService.Add()
		if err != nil {
			responseParser.JsonDBError(c, "", err)
			return
		}
	} else if c.Request.Method == "PUT" {
		args, err := structs.StructToMap(entFileTypeService.EntFileType.EntFileType, "json")
		if err != nil {
			responseParser.JsonParameterIllegal(c, "", err)
		}
		// todo enterpriseId为业务主键名
		delete(args, "fileTypeId")

		temp := services.EnterpriseFileTypeService{}
		temp.FileTypeID = entFileTypeService.FileTypeID
		err = temp.Update(args)
		if err != nil {
			responseParser.JsonDBError(c, "", err)
			return
		}
	} else if c.Request.Method == "DELETE" {
		err = entFileTypeService.Delete()
		if err != nil {
			responseParser.JsonDBError(c, "", err)
			return
		}
	}

	responseParser.JsonOK(c, "", entFileTypeService)
}