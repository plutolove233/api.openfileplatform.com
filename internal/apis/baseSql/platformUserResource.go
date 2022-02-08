// coding: utf-8
// @Author : lryself
// @Date : 2022/2/8 15:50
// @Software: GoLand

package baseSql

import (
	"api.openfileplatform.com/internal/globals/responseParser"
	"api.openfileplatform.com/internal/services"
	"api.openfileplatform.com/internal/utils/structs"
	"github.com/gin-gonic/gin"
)

type PlatformUserImpl struct{}

func (*PlatformUserImpl) PlatformUserApi(c *gin.Context) {
	var err error
	var platUsersService services.PlatUsersService
	err = c.ShouldBind(&platUsersService)

	if c.Request.Method == "GET" {
		err = platUsersService.Get()
		if err != nil {
			responseParser.JsonDBError(c, "", err)
			return
		}
	} else if c.Request.Method == "POST" {
		err = platUsersService.Add()
		if err != nil {
			responseParser.JsonDBError(c,"", err)
			return
		}
	} else if c.Request.Method == "PUT" {
		args, err := structs.StructToMap(platUsersService.PlatUsers.PlatUsers, "json")
		if err != nil {
			responseParser.JsonParameterIllegal(c, err)
		}
		delete(args, "userId")
		temp := services.PlatUsersService{}
		temp.UserID = platUsersService.UserID
		err = temp.Update(args)
		if err != nil {
			responseParser.JsonDBError(c,"", err)
			return
		}
	} else if c.Request.Method == "DELETE" {
		err = platUsersService.Delete()
		if err != nil {
			responseParser.JsonDBError(c,"", err)
			return
		}
	}

	responseParser.JsonOK(c, platUsersService)
}
