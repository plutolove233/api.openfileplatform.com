package enterprise

import (
	"api.openfileplatform.com/commons/codes"
	"api.openfileplatform.com/dao"
	"api.openfileplatform.com/models"
	"api.openfileplatform.com/utils/jwt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func UserLogin(c *gin.Context){
	var ent_user models.EntUser
	err := c.ShouldBind(&ent_user)
	if err != nil {
		c.JSON(200,gin.H{
			"code":codes.ParamError,
			"error":err,
			"msg":"获取员工登录信息失败",
		})
		return
	}

	user := models.EntUser{}
	err = dao.DB.Model(&models.EntUser{}).Where("Account = ?",ent_user.Account).Find(&user).Error
	if err != nil {
		c.JSON(200,gin.H{
			"code":codes.NotData,
			"error":err,
			"msg":"员工信息不存在",
		})
		return
	}
	
	err = bcrypt.CompareHashAndPassword([]byte(user.Pwd),[]byte(ent_user.Pwd))
	if err != nil {
		c.JSON(200,gin.H{
			"code":codes.DataError,
			"error":err,
			"msg":"登录账号或密码错误",
		})
		return
	}

	user.Token,err = jwt.GetToken(user)
	if err != nil{
		c.JSON(200,gin.H{
			"code":codes.InternetError,
			"error":err,
			"msg":"token生成失败",
		})
		return
	}

	err = dao.DB.Model(&models.EntUser{}).Where("UserID = ?",user.UserID).Update("Token",user.Token).Error
	if err != nil {
		c.JSON(200,gin.H{
			"code":codes.DBError,
			"error":err,
			"msg":"用户token更新失败",
		})
		return
	}

	c.JSON(200,gin.H{
		"code":codes.OK,
		"error":nil,
		"msg":user,
	})


	//日志记录
	reqIP := c.ClientIP()//获取IP
	if reqIP == "::1" {
		reqIP = "127.0.0.1"
	}

	ent_user_log := models.EntUserLog{
		UserID:user.UserID,
		UserName:user.UserName,
		Account:user.Account,
		OperationIP:reqIP,
		OperationType:"1",
		OperationContent:"员工登录",
		OperationResult:1,
		OperationStatus:1,
		CreateTime:time.Now(),
	}
	err = dao.DB.Model(&models.EntUserLog{}).Create(ent_user_log).Error

	if err != nil {
		c.JSON(200,gin.H{
			"code":codes.DBError,
			"error":err,
			"msg":"日志记录失败",
		})
	}
}

func UserRegister(c *gin.Context){

}

func GetUsers(c *gin.Context){
	var ent_users []models.EntUser
	err := dao.DB.Model(&models.EntUser{}).Find(&ent_users)
	if err != nil {
		c.JSON(200,gin.H{
			"code":codes.DBError,
			"error":err,
			"msg":"员工信息获取失败",
		})
		return
	}

	c.JSON(200,gin.H{
		"code":codes.OK,
		"error":nil,
		"msg":ent_users,
	})
}