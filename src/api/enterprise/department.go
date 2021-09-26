package enterprise

import (
	"api.openfileplatform.com/commons/codes"
	"api.openfileplatform.com/dao"
	"api.openfileplatform.com/models"
	"api.openfileplatform.com/utils/authority"
	"github.com/gin-gonic/gin"
	"strconv"
)

func NewDepartment(c *gin.Context) {
	if authority.CheckAuthority(c.MustGet("UserID").(int64),codes.NewDepartmentPermission)==false{
		c.JSON(200,gin.H{
			"code":codes.RoleError,
			"msg":"用户没有权限创建部门",
		})
		return
	}

	var dep models.EntDepartment
	err := c.ShouldBind(&dep)
	if err != nil {
		c.JSON(200,gin.H{
			"code":codes.ParamError,
			"error":err,
			"msg":"获取表单参数失败",
		})
		return
	}

	var last models.EntDepartment
	err = dao.DB.Model(&models.EntDepartment{}).Last(&last).Error
	dep.DepartmentID = last.AutoID+101
	dep.IsDeleted = false

	err = dep.AddDepart()
	if err != nil {
		c.JSON(200,gin.H{
			"code":codes.DBError,
			"error":err,
			"msg":"数据上传失败",
		})
		return
	}

	c.JSON(200,gin.H{
		"code":codes.OK,
		"err":"nil",
		"msg":"部门创建成功",
	})
}

func DeleteDepartment(c *gin.Context){
	if authority.CheckAuthority(c.MustGet("UserID").(int64),codes.DeleteDepartmentPermission)==false{
		c.JSON(200,gin.H{
			"code":codes.RoleError,
			"msg":"用户没有权限删除部门",
		})
		return
	}

	departmentID,err := strconv.ParseInt(c.Param("id"),10,10)
	if err != nil {
		c.JSON(200,gin.H{
			"code":codes.ParamError,
			"error":err,
			"msg":"文档编号获取失败",
		})
		return
	}

	var dep models.EntDepartment
	err = dao.DB.Model(&models.EntDepartment{}).Where("DepartmentID = ?",departmentID).Find(&dep).Error
	if err != nil {
		c.JSON(200,gin.H{
			"code":codes.NotData,
			"error":err,
			"msg":"文件信息不存在",
		})
		return
	}

	err = dao.DB.Unscoped().Delete(&dep).Error
	if err != nil {
		c.JSON(200,gin.H{
			"code":codes.DBError,
			"error":err,
			"msg":"文件删除失败",
		})
		return
	}

	c.JSON(200,gin.H{
		"code":codes.OK,
		"error":"nil",
		"msg":dep,
	})
}