// coding:utf-8
// @Author:PigKnight
// @Date:2022/2/10 16:21
// @Software: GoLand

package platEnterprises

import (
	"api.openfileplatform.com/internal/globals/responseParser"
	"api.openfileplatform.com/internal/services"
	"api.openfileplatform.com/internal/utils/structs"
	"github.com/gin-gonic/gin"
)

type PlatformEnterpriseApi struct{}

func (*PlatformEnterpriseApi) PlatformEnterpriseApi(c *gin.Context) {
	var err error
	var platEnterprisesService services.PlatEnterpriseService
	err = c.ShouldBind(&platEnterprisesService)

	if c.Request.Method == "GET" {
		err = platEnterprisesService.Get()
		if err != nil {
			responseParser.JsonDBError(c, "", err)
			return
		}
	} else if c.Request.Method == "POST" {
		err = platEnterprisesService.Add()
		if err != nil {
			responseParser.JsonDBError(c, "", err)
			return
		}
	} else if c.Request.Method == "PUT" {
		args, err := structs.StructToMap(platEnterprisesService.PlatEnterprise.PlatEnterprises, "json")
		if err != nil {
			responseParser.JsonParameterIllegal(c, "", err)
		}
		// todo enterpriseId为业务主键名
		delete(args, "userId")

		temp := services.PlatEnterpriseService{}
		temp.EnterpriseID = platEnterprisesService.EnterpriseID
		err = temp.Update(args)
		if err != nil {
			responseParser.JsonDBError(c, "", err)
			return
		}
	} else if c.Request.Method == "DELETE" {
		err = platEnterprisesService.Delete()
		if err != nil {
			responseParser.JsonDBError(c, "", err)
			return
		}
	}

	responseParser.JsonOK(c, "", platEnterprisesService)
}
