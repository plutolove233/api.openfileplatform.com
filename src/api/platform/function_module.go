package platform

import (
	"api.openfileplatform.com/commons/codes"
	"api.openfileplatform.com/dao"
	"api.openfileplatform.com/models"
	"github.com/gin-gonic/gin"
	"strconv"
)

func NewFunctionModule(c *gin.Context){
	var module models.PlatModule
	err := c.ShouldBind(&module)
	if err != nil {
		c.JSON(200,gin.H{
			"code":codes.ParamError,
			"error":err,
			"msg":"功能模块信息获取失败",
		})
		return
	}

	err = module.NewModule()
	if err != nil {
		c.JSON(200,gin.H{
			"code":codes.DBError,
			"error":err,
			"msg":"信息上传数据库失败",
		})
		return
	}

	c.JSON(200,gin.H{
		"code":codes.OK,
		"msg":module,
	})
}

func DeleteFunctionModule(c *gin.Context){
	id,err := strconv.ParseInt(c.Param("id"),10,10)
	if err != nil {
		c.JSON(200,gin.H{
			"code":codes.ParamError,
			"error":err,
			"msg":"获取删除模块对象失败",
		})
		return
	}

	var module models.PlatModule
	err = dao.DB.Model(&models.PlatModule{}).Where("ModuleID = ?",id).Find(&module).Error
	if err != nil {
		c.JSON(200,gin.H{
			"code":codes.DBError,
			"error":err,
			"msg":"获取模块表单信息失败",
		})
		return
	}

	err = module.DeleteModule()
	if err != nil {
		c.JSON(200,gin.H{
			"code":codes.DBError,
			"error":err,
			"msg":"数据库表单删除失败",
		})
		return
	}

	c.JSON(200,gin.H{
		"code":codes.OK,
		"msg":module,
	})
}

func GetModuleList(c *gin.Context){
	var modules []models.PlatModule
	err := dao.DB.Model(&models.PlatModule{}).Find(&modules).Error
	if err != nil {
		c.JSON(200,gin.H{
			"code":codes.DBError,
			"error":err,
			"msg":"模块信息获取失败",
		})
		return
	}

	c.JSON(200,gin.H{
		"code":codes.OK,
		"msg":modules,
	})
}

func FindModule(c *gin.Context){
	msg := c.PostForm("message")
	key := "%"+msg+"%"
	var modules []models.PlatModule

	err := dao.DB.Model(&models.PlatModule{}).Where("ModuleName like ?",key).Find(&modules).Error
	if err != nil {
		c.JSON(200,gin.H{
			"code":codes.DBError,
			"error":err,
			"msg":"数据库信息请求错误",
		})
		return
	}

	c.JSON(200,gin.H{
		"code":codes.OK,
		"msg":modules,
	})
}