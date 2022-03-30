/*
@Coding : utf-8
@Time : 2022/3/5 15:33
@Author : 刘浩宇
@Software: GoLand
*/
package entFileCategory

import (
	"api.openfileplatform.com/internal/globals/codes"
	"api.openfileplatform.com/internal/globals/responseParser"
	"api.openfileplatform.com/internal/globals/snowflake"
	"api.openfileplatform.com/internal/models/ginModels"
	"api.openfileplatform.com/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"time"
)

type getFileCategoryParser struct {
	CategoryID	string	`json:"CategoryID" form:"CategoryID" binding:"required"`
}

func (*EnterpriseFileCategoryApi)GetFileCategoryPath(c *gin.Context) {
	var parser getFileCategoryParser
	err := c.ShouldBind(&parser)
	if err != nil {
		responseParser.JsonParameterIllegal(c,"获取分类信息失败",err)
		return
	}
	var fileCategoryService services.EnterpriseFileCategoryService
	fileCategoryService.CategoryID = parser.CategoryID

	name_path,id_path,err := fileCategoryService.GetPath()

	if err != nil {
		responseParser.JsonDBError(c,"获取文件分类信息路径失败",err)
		return
	}

	responseParser.JsonOK(c,"获取文件分类信息成功",gin.H{
		"name path":name_path,
		"id path":id_path,
	})
}

type addFileCategoryParser struct {
	CategoryParentID	string 	`json:"CategoryParentID" form:"CategoryParentID" binding:""`
	ProjectID			string	`json:"ProjectID" form:"ProjectID" binding:"required"`
	CategoryName		string	`json:"CategoryName" form:"CategoryName" binding:"required"`
}

func (*EnterpriseFileCategoryApi)AddFileCategory(c *gin.Context)  {
	var parser addFileCategoryParser
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
	entUserService := services.EntUserService{}
	entUserService.UserID = user.UserID
	err = entUserService.Get()
	if err != nil {
		responseParser.JsonNotData(c,"该用户不存在",err)
		return
	}

	var fileCategoryService services.EnterpriseFileCategoryService
	fileCategoryService.CategoryName = parser.CategoryName
	fileCategoryService.EnterpriseID = entUserService.EnterpriseID
	fileCategoryService.ProjectID = parser.ProjectID
	err = fileCategoryService.Get()
	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    codes.DataExist,
			"message": "文件分类信息已存在",
		})
		return
	} else if err.Error() != "record not found" {
		responseParser.JsonDBError(c, "", err)
	}

	fileCategoryService.CategoryID = snowflake.GetSnowflakeID()
	fileCategoryService.CategoryParentID = parser.CategoryParentID
	fileCategoryService.CreatTime = time.Now()
	path := ""
	if parser.CategoryParentID == "" {
		_,id,err1 := fileCategoryService.GetRootPath()
		if err1 != nil {
			responseParser.JsonDBError(c,"获取文件分类根分类失败",err1)
			return
		}
		fileCategoryService.CategoryParentID = id[0:len(id)-1]
	}
	err = fileCategoryService.Add()
	if err != nil {
		responseParser.JsonDBError(c,"创建新的文件分类信息失败",err)
		return
	}
	path,_,err = fileCategoryService.GetPath()
	if err != nil {
		responseParser.JsonDBError(c,"获取相关路径失败",err)
		return
	}
	err = os.Mkdir("save/"+path,os.ModePerm)
	if err != nil {
		responseParser.JsonInternalError(c,"文件夹创建失败",err)
		return
	}

	responseParser.JsonOK(c,"添加新的文件分类信息成功",fileCategoryService)
}

type categoryParser struct {
	CategoryID	    string	`json:"CategoryID"`
	CategoryName	string	`json:"CategoryName"`
}

type queryCategoryParser struct {
	CategoryID	string	`json:"CategoryID" form:"CategoryID" binding:""`
	ProjectID	string	`json:"ProjectID" form:"ProjectID" binding:""`
}

func (*EnterpriseFileCategoryApi) GetAllCategory(c *gin.Context) {
	var parser queryCategoryParser
	err := c.ShouldBind(&parser)
	if err != nil {
		responseParser.JsonParameterIllegal(c,"获取文件类型查询信息失败",err)
		return
	}

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

	categoryService := services.EnterpriseFileCategoryService{}
	categoryService.CategoryID = parser.CategoryID
	categoryService.ProjectID = parser.ProjectID
	categoryService.EnterpriseID = entUser.EnterpriseID
	categoryinfo,err1 := categoryService.GetAll()
	if err1 != nil {
		responseParser.JsonDBError(c,"获取所有分类失败",err1)
		return
	}
	data := []categoryParser{}
	for _,item := range categoryinfo{
		x := categoryParser{
			CategoryID: item.CategoryID,
			CategoryName: item.CategoryName,
		}
		data = append(data, x)
	}
	responseParser.JsonOK(c,"获取所有分类成功",data)
}