/*
@Coding : utf-8
@Time : 2022/2/15 16:31
@Author : 刘浩宇
@Software: GoLand
*/
package entProject

import (
	"api.openfileplatform.com/internal/globals/responseParser"
	"api.openfileplatform.com/internal/services"
	"api.openfileplatform.com/internal/utils/structs"
	"github.com/gin-gonic/gin"
)

type EnterpriseProjectApi struct {}

func (*EnterpriseProjectApi) EnterpriseProjectApi(c *gin.Context) {
	var err error
	var entFilesService services.EnterpriseProjectService
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
		args, err := structs.StructToMap(entFilesService.EntFileProject.EntProject, "json")
		if err != nil {
			responseParser.JsonParameterIllegal(c, "", err)
		}
		// todo projectId为业务主键名
		delete(args, "projectId")

		temp := services.EnterpriseProjectService{}
		temp.ProjectID = entFilesService.ProjectID
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