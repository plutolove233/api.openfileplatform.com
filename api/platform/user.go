package platform

import (
	"DocumentSystem/commons/codes"
	"DocumentSystem/dao"
	"DocumentSystem/models/platform"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func PlatUserLogin(c *gin.Context){
	var plat_user platform.PlatUser
	err := c.ShouldBind(&plat_user)
	if err != nil{
		c.JSON(200,gin.H{
			"code":codes.ParamIllegal,
			"error":err,
			"msg":"参数错误",
		})
		return
	}

	user := platform.PlatUser{}
	err = dao.DB.Model(&platform.PlatUser{}).Where("Account = ?",plat_user.Account).Find(&user).Error

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
		"msg":plat_user,
	})
}

func PlatGetUserList(c *gin.Context){
	var users []platform.PlatUser
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
	var plat_user platform.PlatUser
	err:= c.ShouldBind(&plat_user)
	if err != nil{
		c.JSON(200,gin.H{
			"code":codes.ParamIllegal,
			"error":err,
			"msg":"数据绑定错误",
		})
		return
	}

	user := platform.PlatUser{}

	err = dao.DB.Model(&platform.PlatUser{}).Where("Account = ?",plat_user.Account).Find(&user).Error
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