/*
@Coding : utf-8
@Time : 2022/2/14 10:23
@Author : 刘浩宇
@Software: GoLand
*/
package entFiles

import (
	"api.openfileplatform.com/internal/globals/responseParser"
	"api.openfileplatform.com/internal/services"
	"api.openfileplatform.com/internal/utils/structs"
	"github.com/gin-gonic/gin"
)

type EntFilesApi struct{}

func (*EntFilesApi) EntFilesApi(c *gin.Context) {
	var err error
	var entFilesService services.EnterpriseFilesService
	err = c.ShouldBind(&entFilesService)

	if c.Request.Method == "GET" {
		err = entFilesService.Get()
		if err != nil {
			responseParser.JsonDBError(c, "", err)
			return
		}
	} else if c.Request.Method == "POST" {
		err = entFilesService.Add()
		if err != nil {
			responseParser.JsonDBError(c, "", err)
			return
		}
	} else if c.Request.Method == "PUT" {
		args, err := structs.StructToMap(entFilesService.EntFiles.EntFile, "json")
		if err != nil {
			responseParser.JsonParameterIllegal(c, "", err)
		}
		// todo enterpriseId为业务主键名
		delete(args, "userId")

		temp := services.EnterpriseFilesService{}
		temp.EnterpriseID = entFilesService.EnterpriseID
		err = temp.Update(args)
		if err != nil {
			responseParser.JsonDBError(c, "", err)
			return
		}
	} else if c.Request.Method == "DELETE" {
		err = entFilesService.Delete()
		if err != nil {
			responseParser.JsonDBError(c, "", err)
			return
		}
	}

	responseParser.JsonOK(c, "", entFilesService)
}
