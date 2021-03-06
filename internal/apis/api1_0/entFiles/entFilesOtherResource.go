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
	"os"
	"time"
)

type getAllFileParser struct {
	FileID		string	`json:"FileID" form:"FileID" binding:""`
	CategoryID	string 	`json:"CategoryID" form:"CategoryID" binding:""`
	FileTypeID	string 	`json:"FileTypeID" form:"FileTypeID" binding:""`
	ProjectID	string 	`json:"ProjectID" form:"ProjectID" binding:""`
	Status		int8	`json:"Status" form:"Status" binding:""`
}

func (*EnterpriseFilesApi) GetAllEntFiles(c *gin.Context) {
	var parser getAllFileParser
	err := c.ShouldBind(&parser)
	if err != nil {
		responseParser.JsonParameterIllegal(c,"获取请求失败",err)
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
		responseParser.JsonNotData(c, "该用户id不存在", err)
		return
	}

	file := services.EnterpriseFilesService{}
	file.EnterpriseID = entUser.EnterpriseID
	file.FileID = parser.FileID
	file.FileTypeID = parser.FileTypeID
	file.Status = parser.Status
	file.CategoryID = parser.CategoryID
	info, err := file.GetFileInformation()
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
		responseParser.JsonParameterIllegal(c, "获取文件上传参数失败", err)
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

	data := []services.EnterpriseFilesService{}

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
		data = append(data, file_service)
	}
	responseParser.JsonOK(c,"文件上传成功",data)
}

type moveFileParser struct {
	FileID        string	`json:"FileID" form:"FileID" binding:"required" `
	NewCategoryID string	`json:"NewCategoryID" form:"NewCategoryID" binding:"required"`
}

func (*EnterpriseFilesApi) MoveFile(c *gin.Context) {
	var parser moveFileParser
	err := c.ShouldBind(&parser)
	if err != nil {
		responseParser.JsonParameterIllegal(c,"获取转移信息失败",err)
		return
	}
	fileService := services.EnterpriseFilesService{}
	fileService.FileID = parser.FileID
	err = fileService.Get()
	from := fileService.FileURL
	if err != nil {
		responseParser.JsonNotData(c,"该文件信息不存在",err)
		return
	}
	categoryService := services.EnterpriseFileCategoryService{}
	categoryService.CategoryID = parser.NewCategoryID
	to,_,err1 := categoryService.GetPath()
	if err1 != nil {
		responseParser.JsonInternalError(c,"获取文件路径失败",err1)
		return
	}
	to = "save/" + to + fileService.FileName

	err = os.Rename(from,to)
	if err != nil {
		responseParser.JsonInternalError(c,"移动文件路径失败",err)
		return
	}

	err = fileService.Update(map[string]interface{}{
		"CategoryID":parser.NewCategoryID,
		"FileURL":to,
	})
	if err != nil {
		responseParser.JsonDBError(c,"修改文件所属类别失败",err)
		return
	}

	c.JSON(200,gin.H{
		"code":codes.OK,
		"message":"修改文件路径成功",
	})
}

type deleteFileParser struct {
	FileID	string	`json:"FileID" form:"FileID" binding:"required"`
}

func (*EnterpriseFilesApi)DeleteFile(c *gin.Context){
	var parser deleteFileParser
	err := c.ShouldBind(&parser)
	if err != nil {
		responseParser.JsonParameterIllegal(c,"获取删除的文件信息失败",err)
		return
	}

	fileService := services.EnterpriseFilesService{}
	fileService.FileID = parser.FileID
	err = fileService.Get()
	if err != nil {
		responseParser.JsonNotData(c,"该文件不存在",err)
		return
	}
	categoryID := fileService.CategoryID

	categoryService := services.EnterpriseFileCategoryService{}
	categoryService.CategoryID = categoryID
	path,_,err := categoryService.GetPath()
	if err != nil {
		responseParser.JsonInternalError(c,"获取被删文件路径失败",err)
		return
	}
	err = os.Rename(fileService.FileURL,"save/"+path+parser.FileID)
	if err != nil {
		responseParser.JsonInternalError(c,"删除文件失败",err)
		return
	}
	err = fileService.Delete()
	if err != nil {
		responseParser.JsonDBError(c,"删除文件失败",err)
		return
	}

	c.JSON(200,gin.H{
		"code":codes.OK,
		"meesage":"删除文件成功",
	})
}

func (*EnterpriseFilesApi) Confirm(c *gin.Context) {
	var parser []services.EnterpriseFilesService
	err := c.ShouldBind(&parser)
	if err != nil {
		responseParser.JsonParameterIllegal(c,"获取文件信息失败",err)
		return
	}
	for _,item := range parser{
		item.UpdateTime = time.Now()
		if err1 := item.Add(); err1 != nil{
			responseParser.JsonDBError(c,"保存到数据库失败",err1)
			return
		}
	}
	c.JSON(200,gin.H{
		"code":codes.OK,
		"message":"保存到数据库成功",
	})
}