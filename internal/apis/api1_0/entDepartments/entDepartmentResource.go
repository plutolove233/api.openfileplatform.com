/*
@Coding : utf-8
@Time : 2022/2/12 9:31
@Author : 刘浩宇
@Software: GoLand
*/
package entDepartments

import (
	"api.openfileplatform.com/internal/globals/responseParser"
	"api.openfileplatform.com/internal/services"
	"api.openfileplatform.com/internal/utils/structs"
	"github.com/gin-gonic/gin"
)

type EnterpriseDepartmentApi struct{}

func (*EnterpriseDepartmentApi) EnterpriseDepartmentApi(c *gin.Context) {
	var err error
	var entDepartmentService services.EntDepartmentService
	err = c.ShouldBind(&entDepartmentService)

	if c.Request.Method == "GET" {
		err = entDepartmentService.Get()
		if err != nil {
			responseParser.JsonDBError(c, "", err)
			return
		}
	} else if c.Request.Method == "POST" {
		err = entDepartmentService.Add()
		if err != nil {
			responseParser.JsonDBError(c, "", err)
			return
		}
	} else if c.Request.Method == "PUT" {
		args, err := structs.StructToMap(entDepartmentService.EntDepartment.EntDepartments, "json")
		if err != nil {
			responseParser.JsonParameterIllegal(c, "", err)
		}
		// todo departmentId为业务主键名
		delete(args, "departmentId")

		temp := services.EntDepartmentService{}
		temp.DepartmentID = entDepartmentService.DepartmentID
		err = temp.Update(args)
		if err != nil {
			responseParser.JsonDBError(c, "", err)
			return
		}
	} else if c.Request.Method == "DELETE" {
		err = entDepartmentService.Delete()
		if err != nil {
			responseParser.JsonDBError(c, "", err)
			return
		}
	}

	responseParser.JsonOK(c, "", entDepartmentService)
}
