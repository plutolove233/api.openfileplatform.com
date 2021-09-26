package enterprise

import (
	"api.openfileplatform.com/commons/codes"
	"api.openfileplatform.com/dao"
	"api.openfileplatform.com/models"
	"api.openfileplatform.com/utils/authority"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

func NewRole(c *gin.Context){
	if authority.CheckAuthority(c.MustGet("UserID").(int64),codes.NewRolePermission)==false{
		c.JSON(200,gin.H{
			"code":codes.RoleError,
			"msg":"用户没有权限创建新的角色",
		})
		return
	}
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
	if authority.CheckAuthority(c.MustGet("UserID").(int64),codes.GetRoleListPermission)==false{
		c.JSON(200,gin.H{
			"code":codes.RoleError,
			"msg":"用户没有权限获取角色信息表",
		})
		return
	}
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
	if authority.CheckAuthority(c.MustGet("UserID").(int64),codes.DeleteRolePermission)==false{
		c.JSON(200,gin.H{
			"code":codes.RoleError,
			"msg":"没有权限删除角色表",
		})
		return
	}

	roleID,err := strconv.ParseInt(c.Param("id"),10,10)
	if err != nil {
		c.JSON(200,gin.H{
			"code":codes.ParamError,
			"error":err,
			"msg":"删除文档编号获取失败",
		})
		return
	}

	var role models.EntRole
	err = dao.DB.Model(&models.EntRole{}).Where("RoleID = ?",roleID).Find(&role).Error
	if err != nil {
		c.JSON(200,gin.H{
			"code":codes.NotData,
			"error":err,
			"msg":"相关角色信息不存在",
		})
		return
	}

	err = dao.DB.Unscoped().Delete(&role).Error
	if err != nil {
		c.JSON(200,gin.H{
			"code":codes.DBError,
			"error":err,
			"msg":"角色信息删除失败",
		})
		return
	}

	c.JSON(200,gin.H{
		"code":codes.OK,
		"msg":role,
	})
}