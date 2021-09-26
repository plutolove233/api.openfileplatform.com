package platform

import (
	"api.openfileplatform.com/commons/codes"
	"api.openfileplatform.com/dao"
	"api.openfileplatform.com/models"
	"github.com/gin-gonic/gin"
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

	err = dao.DB.Model(&models.PlatEnterprise{}).Create(&ent).Error
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