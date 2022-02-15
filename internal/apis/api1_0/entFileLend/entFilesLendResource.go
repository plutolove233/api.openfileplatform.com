/*
@Coding : utf-8
@Time : 2022/2/15 15:28
@Author : 刘浩宇
@Software: GoLand
*/
package entFileLend

import (
	"api.openfileplatform.com/internal/globals/responseParser"
	"api.openfileplatform.com/internal/services"
	"api.openfileplatform.com/internal/utils/structs"
	"github.com/gin-gonic/gin"
)

type EnterpriseFileLendApi struct{}

func (*EnterpriseFileLendApi) EnterpriseFileLendApi(c *gin.Context) {
	var err error
	var entFileLendService services.EnterpriseFileLendService
	err = c.ShouldBind(&entFileLendService)

	if c.Request.Method == "GET" {
		err = entFileLendService.Get()
		if err != nil {
			responseParser.JsonDBError(c, "", err)
			return
		}
	} else if c.Request.Method == "POST" {
		err = entFileLendService.Add()
		if err != nil {
			responseParser.JsonDBError(c, "", err)
			return
		}
	} else if c.Request.Method == "PUT" {
		args, err := structs.StructToMap(entFileLendService.EntFileLend.EntFileLend, "json")
		if err != nil {
			responseParser.JsonParameterIllegal(c, "", err)
		}
		// todo fileID为业务主键名
		delete(args, "fileId")

		temp := services.EnterpriseFileLendService{}
		temp.FileID = entFileLendService.FileID
		err = temp.Update(args)
		if err != nil {
			responseParser.JsonDBError(c, "", err)
			return
		}
	} else if c.Request.Method == "DELETE" {
		err = entFileLendService.Delete()
		if err != nil {
			responseParser.JsonDBError(c, "", err)
			return
		}
	}

	responseParser.JsonOK(c, "", entFileLendService)
}