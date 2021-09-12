package enterprise

import (
	"api.openfileplatform.com/commons/codes"
	"api.openfileplatform.com/dao"
	"api.openfileplatform.com/models"
	"github.com/gin-gonic/gin"
	"time"
)

func NewRole(c *gin.Context){
	var role,last models.EntRole
	err := c.ShouldBind(&role)
	if err != nil {
		c.JSON(200,gin.H{
			"code":codes.ParamError,
			"error":err,
			"msg":"获取信息失败",
		})
		return
	}

	dao.DB.Model(&models.EntRole{}).Last(&last)
	role.RoleID = last.AutoID+101
	role.CreatTime = time.Now()
	role.IsDeleted = false

	var _role models.EntRole
	err = dao.DB.Model(&models.EntRole{}).Where("RoleName = ?",role.RoleName).Find(&_role).Error
	if err == nil {
		c.JSON(200,gin.H{
			"code":codes.DataExist,
			"error":"data exist",
			"msg":"该角色信息存在",
		})
		return
	}

	err = role.NewRole()
	if err != nil {
		c.JSON(200,gin.H{
			"code":codes.DBError,
			"error":err,
			"msg":"角色新建失败",
		})
		return
	}

	c.JSON(200,gin.H{
		"code":codes.OK,
		"error":"nil",
		"msg":role,
	})
}

func GetRoleList(c *gin.Context){
	var role []models.EntRole
	err := dao.DB.Model(&models.EntRole{}).Find(&role).Error
	if err != nil {
		c.JSON(200,gin.H{
			"code":codes.DBError,
			"error":err,
			"msg":"获取员工角色信息错误",
		})
		return
	}

	c.JSON(200,gin.H{
		"code":codes.OK,
		"error":"nil",
		"msg":role,
	})
}

func DeleteRole(c *gin.Context){

}