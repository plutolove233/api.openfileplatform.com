/*
@Coding : utf-8
@Time : 2022/2/14 16:13
@Author : 刘浩宇
@Software: GoLand
*/
package entFileCategory

import (
	"api.openfileplatform.com/internal/globals/responseParser"
	"api.openfileplatform.com/internal/services"
	"api.openfileplatform.com/internal/utils/structs"
	"github.com/gin-gonic/gin"
)

type EntFileCategoryApi struct {}

func (*EntFileCategoryApi) EntDepartmenrtApi(c *gin.Context) {
	var err error
	var entFileCategoryService services.EntFileCategoryService
	err = c.ShouldBind(&entFileCategoryService)

	if c.Request.Method == "GET" {
		err = entFileCategoryService.Get()
		if err != nil {
			responseParser.JsonDBError(c, "", err)
			return
		}
	} else if c.Request.Method == "POST" {
		err = entFileCategoryService.Add()
		if err != nil {
			responseParser.JsonDBError(c, "", err)
			return
		}
	} else if c.Request.Method == "PUT" {
		args, err := structs.StructToMap(entFileCategoryService.EntFileCategory.EntFileCategory, "json")
		if err != nil {
			responseParser.JsonParameterIllegal(c, "", err)
		}
		// todo enterpriseId为业务主键名
		delete(args, "userId")

		temp := services.EntFileCategoryService{}
		temp.EnterpriseID = entFileCategoryService.EnterpriseID
		err = temp.Update(args)
		if err != nil {
			responseParser.JsonDBError(c, "", err)
			return
		}
	} else if c.Request.Method == "DELETE" {
		err = entFileCategoryService.Delete()
		if err != nil {
			responseParser.JsonDBError(c, "", err)
			return
		}
	}

	responseParser.JsonOK(c, "", entFileCategoryService)
}
