/*
@Coding : utf-8
@Time : 2022/3/2 20:55
@Author : 刘浩宇
@Software: GoLand
*/
package entProject

import (
	"api.openfileplatform.com/internal/globals/responseParser"
	"api.openfileplatform.com/internal/globals/snowflake"
	"api.openfileplatform.com/internal/models/ginModels"
	"api.openfileplatform.com/internal/services"
	"github.com/gin-gonic/gin"
	"time"
)

type addPreojectParser struct {
	EnterpriseID 	string	`json:"EnterpriseID" form:"EnterpriseID" binding:"required"`
	ProjectName		string	`json:"ProjectName" form:"ProjectName" binding:"required"`
}

func (*EnterpriseProjectApi)AddNewProject(c *gin.Context){
	var parser addPreojectParser
	err := c.ShouldBind(&parser)
	if err != nil {
		responseParser.JsonParameterIllegal(c,"获取档案项目信息失败",err)
		return
	}

	var projectService services.EnterpriseProjectService
	projectService.ProjectName = parser.ProjectName
	projectService.EnterpriseID = parser.EnterpriseID
	projectService.CreatTime = time.Now()
	err = projectService.Get()
	if err == nil {
		responseParser.JsonDataExist(c,"项目信息名称存在")
		return
	} else if err.Error() != "record not found" {
		responseParser.JsonDBError(c,"",err)
		return
	}

	projectService.ProjectID = snowflake.GetSnowflakeID()
	projectService.UpdateTime = time.Now()

	err = projectService.Add()
	if err != nil {
		responseParser.JsonDBError(c,"企业项目信息添加失败",err)
		return
	}

	responseParser.JsonOK(c,"",projectService)
}

type changeProjectParser struct {
	ProjectID	string 	`json:"ProjectID" form:"ProjectID" binding:"required"`
	ProjectName	string	`json:"ProjectName" form:"ProjectName" binding:""`
}

func (*EnterpriseProjectApi)ChangeProjectName(c *gin.Context){
	var parser changeProjectParser
	err := c.ShouldBind(&parser)
	if err != nil {
		responseParser.JsonParameterIllegal(c,"获取项目信息失败",err)
		return
	}

	var projectService services.EnterpriseProjectService
	projectService.ProjectID = parser.ProjectID
	err = projectService.Update(map[string]interface{}{
		"ProjectName":parser.ProjectName,
	})
	if err != nil {
		responseParser.JsonDBError(c,"项目信息更新失败",err)
		return
	}
	responseParser.JsonOK(c,"信息更新成功",nil)
}

type projectParser struct {
	ProjectID   string `json:"ProjectID"`
	ProjectName string `json:"ProjectName"`
}

func (*EnterpriseProjectApi) GetAllProject(c *gin.Context) {
	temp,ok := c.Get("user")
	if !ok {
		responseParser.JsonNotData(c,"用户未登录",nil)
		return
	}
	user := temp.(ginModels.UserModel)
	entUser := services.EntUserService{}
	entUser.UserID = user.UserID
	if err := entUser.Get(); err != nil{
		responseParser.JsonNotData(c,"该用户不存在",err)
		return
	}

	projectService := services.EnterpriseProjectService{}
	projectinfo,err1 := projectService.GetAll(entUser.EnterpriseID)
	if err1 != nil {
		responseParser.JsonDBError(c,"获取所有分类失败",err1)
		return
	}
	data := []projectParser{}
	for _,item := range projectinfo {
		x := projectParser{
			ProjectID: item.ProjectID,
			ProjectName: item.ProjectName,
		}
		data = append(data, x)
	}
	responseParser.JsonOK(c,"获取所有分类成功",data)
}