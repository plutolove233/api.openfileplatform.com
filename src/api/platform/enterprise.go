package platform

import (
	"api.openfileplatform.com/commons/codes"
	"api.openfileplatform.com/dao"
	"api.openfileplatform.com/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func NewEnterprise(c *gin.Context){
	var ent models.PlatEnterprise
	err := c.ShouldBind(&ent)
	if err != nil {
		c.JSON(200,gin.H{
			"code":codes.ParamError,
			"error":err,
			"msg":"企业信息获取失败",
		})
		return
	}

	var last models.PlatEnterprise
	dao.DB.Model(&models.PlatEnterprise{}).Last(&last)
	ent.EnterpriseID = last.AutoID+1
	ent.EnterpriseUrl = "1"//?
	ent.LogoPicUrl = "pic/index.png"

	err = ent.Add()
	if err != nil {
		c.JSON(200,gin.H{
			"code":codes.DBError,
			"error":err,
			"msg":"数据保存错误",
		})
		return
	}

	c.JSON(200,gin.H{
		"code":codes.OK,
	})
}

func DeleteEnterprise(c *gin.Context){
	entID,err := strconv.ParseInt(c.Param("id"),10,10)
	if err != nil {
		c.JSON(200,gin.H{
			"code":codes.ParamError,
			"error":err,
			"msg":"企业信息获取失败",
		})
		return
	}

	var ent models.PlatEnterprise
	err = dao.DB.Model(&models.PlatEnterprise{}).Where("EnterpriseID = ? ",entID).Find(&ent).Error
	if err != nil {
		c.JSON(200,gin.H{
			"code":codes.DBError,
			"error":err,
			"msg":"数据库错误",
		})
		return
	}

	err = ent.Delete()
	if err != nil {
		c.JSON(200,gin.H{
			"code":codes.DBError,
			"error":err,
			"msg":"数据删除失败",
		})
		return
	}

	c.JSON(200,gin.H{
		"code":codes.OK,
		"msg":ent,
	})
}

func GetEnterpriseList(c *gin.Context){
	var list []models.PlatEnterprise
	err := dao.DB.Model(&models.PlatEnterprise{}).Find(&list).Error
	if err != nil {
		c.JSON(200,gin.H{
			"code":codes.DBError,
			"error":err,
			"msg":"平台部门信息表获取失败",
		})
		return
	}

	c.JSON(200,gin.H{
		"code":codes.OK,
		"msg":list,
	})
}

func ChangePWD(c *gin.Context){
	pwd,ok := c.GetPostForm("pwd")
	if !ok {
		c.JSON(200,gin.H{
			"code":codes.ParamError,
			"msg":"请输入密码",
		})
		return
	}

	dao.DB.Model(&models.PlatEnterprise{}).Where("EnterpriseID",c.MustGet("EnterpriseID").(int64)).Update("EnterprisePwd",pwd)
	c.JSON(200,gin.H{
		"code":codes.OK,
		"msg":"密码修改成功",
	})
}

func ChangeLogo(c *gin.Context){
	file,err := c.FormFile("pic")
	if err != nil {
		c.JSON(200,gin.H{
			"code":codes.UpdateError,
			"error":err,
			"msg":"企业logo上传失败",
		})
		return
	}

	dst := fmt.Sprintf("./pic/logo/%s",file.Filename)
	err = c.SaveUploadedFile(file,dst)
	if err != nil {
		c.JSON(200,gin.H{
			"code":codes.UpdateError,
			"error":err,
			"msg":"企业logo保存失败",
		})
		return
	}

	entId := c.Param("id")
	var ent models.PlatEnterprise
	dao.DB.Model(&models.PlatEnterprise{}).Where("UserID = ?", entId).Find(&ent)
	err = ent.ChangeLogo(dst)
	if err != nil {
		c.JSON(200,gin.H{
			"code":codes.DBError,
			"error":err,
			"msg":"企业logo更新失败",
		})
		return
	}

	c.JSON(200,gin.H{
		"code":codes.OK,
		"msg":dst,
	})
}