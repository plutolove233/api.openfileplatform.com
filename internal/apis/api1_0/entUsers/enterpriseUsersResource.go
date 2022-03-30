package entUsers

import (
	"api.openfileplatform.com/internal/globals/responseParser"
	"api.openfileplatform.com/internal/services"
	"api.openfileplatform.com/internal/utils/structs"
	"github.com/gin-gonic/gin"
)

type EnterpriseUserApi struct{}

func (*EnterpriseUserApi) EnterpriseUserApi(c *gin.Context) {
	var err error
	var entUsersService services.EntUserService
	err = c.ShouldBind(&entUsersService)

	if c.Request.Method == "GET" {
		err = entUsersService.Get()
		if err != nil {
			responseParser.JsonDBError(c, "", err)
			return
		}
	} else if c.Request.Method == "POST" {
		err = entUsersService.Add()
		if err != nil {
			responseParser.JsonDBError(c, "", err)
			return
		}
	} else if c.Request.Method == "PUT" {
		args, err := structs.StructToMap(entUsersService.EntUsers.EntUsers, "json")
		if err != nil {
			responseParser.JsonParameterIllegal(c, "", err)
		}
		// todo userId为业务主键名
		delete(args, "userId")

		temp := services.EntUserService{}
		temp.UserID = entUsersService.UserID
		err = temp.Update(args)
		if err != nil {
			responseParser.JsonDBError(c, "", err)
			return
		}
	} else if c.Request.Method == "DELETE" {
		err = entUsersService.Delete()
		if err != nil {
			responseParser.JsonDBError(c, "", err)
			return
		}
	}

	responseParser.JsonOK(c, "", entUsersService)
}
