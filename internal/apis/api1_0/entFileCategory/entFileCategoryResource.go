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

type EnterpriseFileCategoryApi struct {}

func (*EnterpriseFileCategoryApi) EnterpriseFileCategoryApi(c *gin.Context) {
	var err error
	var entFileCategoryService services.EnterpriseFileCategoryService
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
		// todo categoryId为业务主键名
		delete(args, "categoryId")

		temp := services.EnterpriseFileCategoryService{}
		temp.CategoryID = entFileCategoryService.CategoryID
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
