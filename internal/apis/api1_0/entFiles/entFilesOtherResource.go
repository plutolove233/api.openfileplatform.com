/*
@Coding : utf-8
@Time : 2022/2/14 15:06
@Author : 刘浩宇
@Software: GoLand
*/
package entFiles

import (
	"api.openfileplatform.com/internal/globals/codes"
	"api.openfileplatform.com/internal/globals/responseParser"
	"api.openfileplatform.com/internal/globals/snowflake"
	"api.openfileplatform.com/internal/models/ginModels"
	"api.openfileplatform.com/internal/services"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func (*EnterpriseFilesApi) GetAllEntFiles(c *gin.Context) {
	temp, ok := c.Get("user")
	if !ok {
		responseParser.JsonNotData(c, "用户未登录", nil)
		return
	}

	user := temp.(ginModels.UserModel)

	entUser := services.EntUserService{}
	entUser.UserID = user.UserID
	err := entUser.Get()
	if err != nil {
		responseParser.JsonNotData(c, "该用户id不存在", err)
		return
	}

	file := services.EnterpriseFilesService{}
	info, err := file.GetAll(entUser.EnterpriseID)
	if err != nil {
		responseParser.JsonDBError(c, "获取企业文件信息失败", err)
		return
	}
	responseParser.JsonOK(c, "获取企业文件信息成功", info)
}

type fileUploadParser struct {
	CategoryID  string `json:"CategoryID" form:"CategoryID" binding:""`
	ProjectID   string `json:"ProjectID" form:"ProjectID" binding:""`
	FileTypeID  string `json:"FileTypeID" form:"FileTypeID" binding:""`
	FileCabinet string `json:"FileCabinet" form:"FileCabinet" binding:"required"`
}

func (*EnterpriseFilesApi) UploadFile(c *gin.Context) {
	var parser fileUploadParser
	err := c.ShouldBind(&parser)
	if err != nil {
		responseParser.JsonParameterIllegal(c, "获取文件上传参数成功", err)
		return
	}

	temp, ok := c.Get("user")
	if !ok {
		responseParser.JsonNotData(c, "用户未登录", nil)
		return
	}
	user := temp.(ginModels.UserModel)
	entUser := services.EntUserService{}
	entUser.UserID = user.UserID
	err = entUser.Get()
	if err != nil {
		responseParser.JsonNotData(c,"该用户不存在",err)
		return
	}

	form, err := c.MultipartForm()
	if err != nil {
		responseParser.JsonParameterIllegal(c, "获取上传文件列表失败", err)
		return
	}

	files := form.File["file"]
	for _, file := range files {
		filePath := ""
		file_category := services.EnterpriseFileCategoryService{}
		if parser.CategoryID != "" {
			file_category.CategoryID = parser.CategoryID
			path, _, err := file_category.GetPath()
			if err != nil {
				responseParser.JsonInternalError(c, "获取文件保存路径失败", err)
				return
			}
			filePath ="save/" + path + file.Filename
			fmt.Println(filePath)
		} else {
			file_category.EnterpriseID = entUser.EnterpriseID
			path, _, err := file_category.GetRootPath()
			if err != nil {
				responseParser.JsonInternalError(c, "获取文件保存路径失败", err)
				return
			}
			filePath = "save/" + path + file.Filename
		}

		if err = c.SaveUploadedFile(file, filePath); err != nil {
			responseParser.JsonInternalError(c, file.Filename+"上传失败", err)
			return
		}

		file_service := services.EnterpriseFilesService{}
		file_service.FileID = snowflake.GetSnowflakeID()
		file_service.FileCabinet = parser.FileCabinet
		file_service.FileTypeID = parser.FileTypeID
		file_service.CategoryID = parser.CategoryID
		file_service.FileName = file.Filename
		file_service.UserID = entUser.UserID
		file_service.EnterpriseID = entUser.EnterpriseID
		file_service.ProjectID = parser.ProjectID
		file_service.FileURL = filePath
		file_service.CreatTime = time.Now()
		if err = file_service.Add(); err != nil {
			responseParser.JsonDBError(c,"上传文件信息数据库保存失败",err)
			return
		}
	}
	c.JSON(200,gin.H{
		"code":codes.OK,
		"message":"文件上传成功",
	})
}
