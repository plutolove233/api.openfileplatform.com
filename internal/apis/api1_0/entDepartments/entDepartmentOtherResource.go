/*
@Coding : utf-8
@Time : 2022/3/5 15:01
@Author : 刘浩宇
@Software: GoLand
*/
package entDepartments

import (
	"api.openfileplatform.com/internal/globals/codes"
	"api.openfileplatform.com/internal/globals/responseParser"
	"api.openfileplatform.com/internal/models/ginModels"
	"api.openfileplatform.com/internal/services"
	"github.com/gin-gonic/gin"
)

type setHeaderParser struct {
	EnterpriseID	string	`json:"EnterpriseID" form:"EnterpriseID" binding:"required"`
	DepartmentID	string	`json:"DepartmentID" form:"DepartmentID" binding:"required"`
	HeadID			string	`json:"HeadID" form:"HeadID" binding:"required"`
}

func (*EnterpriseDepartmentApi)SetHeader(c *gin.Context) {
	var parser setHeaderParser
	err := c.ShouldBind(&parser)
	if err != nil {
		responseParser.JsonParameterIllegal(c,"获取请求信息失败",err)
		return
	}

	temp,ok := c.Get("user")
	if !ok {
		responseParser.JsonNotData(c,"用户未登录",nil)
		return
	}
	user := temp.(ginModels.UserModel)
	if !user.VerifyAdminRole() {
		responseParser.JsonAccessDenied(c,"用户没有权限访问")
		return
	}

	var departmentService services.EntDepartmentService
	departmentService.DepartmentID = parser.DepartmentID
	departmentService.EnterpriseID = parser.EnterpriseID
	err = departmentService.Get()
	if err != nil {
		responseParser.JsonDBError(c,"获取企业部门表失败",err)
		return
	}

	err = departmentService.Update(map[string]interface{}{
		"HeadID":parser.HeadID,
	})

	if err != nil {
		responseParser.JsonDBError(c,"设置部门管理员失败",err)
		return
	}

	c.JSON(200,gin.H{
		"code":codes.OK,
		"message":"设置部门管理员成功",
	})
}