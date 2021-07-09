//平台用户UserID与AutoID一致
//平台用户Account由平台管理员分配
package platform

import (
	"api.openfileplatform.com/commons/codes"
	"api.openfileplatform.com/dao"
	"api.openfileplatform.com/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func PlatUserLogin(c *gin.Context){
	var plat_user models.PlatUser
	err := c.ShouldBind(&plat_user)
	if err != nil{
		c.JSON(200,gin.H{
			"code":codes.ParamError,
			"error":err,
			"msg":"参数错误",
		})
		return
	}

	user := models.PlatUser{}
	err = dao.DB.Model(&models.PlatUser{}).Where("Account = ?",plat_user.Account).Find(&user).Error

	//验证登录信息是否存在
	if err != nil{
		c.JSON(200,gin.H{
			"code":codes.NotData,
			"error":err,
			"msg":"登录信息不存在",
		})
		return
	}

	//解密
	err = bcrypt.CompareHashAndPassword([]byte(user.Pwd),[]byte(plat_user.Pwd))
	if err !=nil {
		c.JSON(200,gin.H{
			"code":codes.DataError,
			"error":err,
			"msg":"密码错误",
		})
		return
	}

	c.JSON(200,gin.H{
		"code":codes.OK,
		"error":"nil",
		"msg":user,
	})
}

func PlatGetUserList(c *gin.Context){
	var users []models.PlatUser
	err := dao.DB.Find(&users).Error
	if err != nil {
		c.JSON(200,gin.H{
			"code":codes.DBError,
			"error":err,
			"msg":"数据导出失败",
		})
		return
	}

	c.JSON(200,gin.H{
		"code":codes.OK,
		"error":"nil",
		"msg":users,
	})
}

func PlatUserRegister(c *gin.Context){
	var plat_user models.PlatUser
	var last models.PlatUser
	dao.DB.Model(&models.PlatUser{}).Last(&last)
	err:= c.ShouldBind(&plat_user)
	if err != nil{
		c.JSON(200,gin.H{
			"code":codes.ParamError,
			"error":err,
			"msg":"数据绑定错误",
		})
		return
	}

	plat_user.CreateTime = time.Now()
	plat_user.UserID = last.AutoID+101

	user := models.PlatUser{}

	err = dao.DB.Model(&models.PlatUser{}).Where("Account = ?",plat_user.Account).Find(&user).Error
	if err == nil {
		c.JSON(200,gin.H{
			"code":codes.DataExist,
			"error":"data exist",
			"msg":"注册信息存在",
		})
		return
	}

	user = plat_user

	cipherTEXT,err := bcrypt.GenerateFromPassword([]byte(plat_user.Pwd),bcrypt.DefaultCost)
	if err != nil {
		c.JSON(200,gin.H{
			"code":codes.InternetError,
			"error":err,
			"msg":"密码加密出错",
		})
		return
	}
	plat_user.Pwd = string(cipherTEXT)

	err = plat_user.Add()
	if err != nil {
		c.JSON(200,gin.H{
			"code":codes.DBError,
			"error":err,
			"msg":"数据库存贮错误",
		})
		return
	}

	c.JSON(200,gin.H{
		"code":codes.OK,
		"error":"nil",
		"msg":user,
	})
}

type ResetPwd struct {
	Email string `form:"email" binding:"required"`
	Verification string `form:"verify" binding:"required"`
	Pwd string `form:"pwd" binding:"required"`
}

func PlatResetPwd(c *gin.Context){
	var change ResetPwd
	err := c.ShouldBind(&change)
	if err != nil {
		c.JSON(200,gin.H{
			"code":codes.ParamError,
			"error":err,
			"msg":"数据绑定错误",
		})
		return
	}

	var plat_user models.PlatUser
	err = dao.DB.Model(&models.PlatUser{}).Where("Email = ?", change.Email).Find(&plat_user).Error
	if err != nil {
		c.JSON(200,gin.H{
			"code":codes.NotData,
			"error":err,
			"msg":"该邮箱用户不存在",
		})
		return
	}

	cipherTEXT,err := bcrypt.GenerateFromPassword([]byte(plat_user.Pwd),bcrypt.DefaultCost)
	if err != nil {
		c.JSON(200,gin.H{
			"code":codes.InternetError,
			"error":err,
			"msg":"密码加密出错",
		})
		return
	}

	err = plat_user.ChangePwd(string(cipherTEXT))
	if err != nil {
		c.JSON(200,gin.H{
			"code":codes.InternetError,
			"error":err,
			"msg":"数据库存储失败",
		})
		return
	}

	c.JSON(200,gin.H{
		"code":codes.OK,
		"error":"nil",
		"msg":plat_user,
	})
}

type ResetPhone struct {
	Email string `form:"email" binding:"required"`
	Verification string `form:"verify" binding:"required"`
	Phone string `form:"phone" binding:"required"`
}

func PlatResetPhone(c *gin.Context){
	var change ResetPhone
	err := c.ShouldBind(&change)
	if err != nil {
		c.JSON(200,gin.H{
			"code":codes.ParamError,
			"error":err,
			"msg":"数据绑定错误",
		})
		return
	}

	var plat_user models.PlatUser
	err = dao.DB.Model(&models.PlatUser{}).Where("Email = ?", change.Email).Find(&plat_user).Error
	if err != nil {
		c.JSON(200,gin.H{
			"code":codes.NotData,
			"error":err,
			"msg":"该邮箱用户不存在",
		})
		return
	}

	err = plat_user.ChangePhone(change.Phone)
	if err != nil {
		c.JSON(200,gin.H{
			"code":codes.InternetError,
			"error":err,
			"msg":"数据库存储失败",
		})
		return
	}

	c.JSON(200,gin.H{
		"code":codes.OK,
		"error":"nil",
		"msg":plat_user,
	})
}

type ResetEmail struct {
	Email string `form:"email" binding:"required"`
	Verification string `form:"verify" binding:"required"`
	NewEmail string `form:"pwd" binding:"required"`
}

func PlatResetEmail(c *gin.Context){
	var change ResetEmail
	err := c.ShouldBind(&change)
	if err != nil {
		c.JSON(200,gin.H{
			"code":codes.ParamError,
			"error":err,
			"msg":"数据绑定错误",
		})
		return
	}

	var plat_user models.PlatUser
	err = dao.DB.Model(&models.PlatUser{}).Where("Email = ?", change.Email).Find(&plat_user).Error
	if err != nil {
		c.JSON(200,gin.H{
			"code":codes.NotData,
			"error":err,
			"msg":"该邮箱用户不存在",
		})
		return
	}

	err = plat_user.ChangeEmail(change.NewEmail)
	if err != nil {
		c.JSON(200,gin.H{
			"code":codes.DBError,
			"error":err,
			"msg":"数据库存储失败",
		})
		return
	}

	c.JSON(200,gin.H{
		"code":codes.OK,
		"error":"nil",
		"msg":plat_user,
	})
}