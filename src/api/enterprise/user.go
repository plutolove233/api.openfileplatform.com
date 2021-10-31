package enterprise

import (
	"api.openfileplatform.com/commons/codes"
	"api.openfileplatform.com/dao"
	"api.openfileplatform.com/models"
	"api.openfileplatform.com/utils/authority"
	"api.openfileplatform.com/utils/jwt"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"strconv"
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
			"token":user.Token,
		})
		return
	}

	c.JSON(200,gin.H{
		"code":codes.OK,
		"error":nil,
		"msg":user,
	})


	////日志记录
	//reqIP := c.ClientIP()//获取IP
	//if reqIP == "::1" {
	//	reqIP = "127.0.0.1"
	//}
	//
	//ent_user_log := models.EntUserLog{
	//	UserID:user.UserID,
	//	UserName:user.UserName,
	//	Account:user.Account,
	//	OperationIP:reqIP,
	//	OperationType:"1",
	//	OperationContent:"员工登录",
	//	OperationResult:1,
	//	OperationStatus:1,
	//	CreateTime:time.Now(),
	//}
	//err = dao.DB.Model(&models.EntUserLog{}).Create(ent_user_log).Error
	//
	//if err != nil {
	//	c.JSON(200,gin.H{
	//		"code":codes.DBError,
	//		"error":err,
	//		"msg":"日志记录失败",
	//	})
	//}
}

func UserRegister(c *gin.Context){
	var ent_user models.EntUser
	err := c.ShouldBind(&ent_user)
	if err != nil {
		c.JSON(200,gin.H{
			"code":codes.ParamError,
			"error":err,
			"msg":"注册用户信息获取失败",
		})
		return
	}

	var user models.EntUser
	err1  := dao.DB.Model(&models.EntUser{}).Where("Account = ?",ent_user.Account).Find(&user).Error
	if err1 == nil {
		c.JSON(200,gin.H{
			"code":codes.DataExist,
			"error":err,
			"msg":"用户已注册",
		})
		return
	}

	var last models.EntUser
	dao.DB.Model(&models.EntUser{}).Last(&last)
	ent_user.CreateTime = time.Now()
	ent_user.UserID = last.AutoID+101

	user = ent_user
	cipherText,err1 := bcrypt.GenerateFromPassword([]byte(ent_user.Pwd),bcrypt.DefaultCost)
	if err1 != nil {
		c.JSON(200,gin.H{
			"code":codes.InternetError,
			"error":err1,
			"msg":"密码加密错误",
		})
		return
	}
	ent_user.Pwd = string(cipherText)

	ent_user.FacePicUrl = "pic/index.png"
	err1  = ent_user.Register()
	if err1 != nil {
		c.JSON(200,gin.H{
			"code":codes.DBError,
			"error":err1,
			"msg":"数据库存储失败",
		})
		return
	}

	c.JSON(200,gin.H{
		"code":codes.OK,
		"error":"nil",
		"msg":ent_user,
	})
}

func GetUsers(c *gin.Context){
	var ent_users []models.EntUser
	enterpriseID,_ := strconv.ParseInt(c.PostForm("EnterpriseID"),10,64)
	err := dao.DB.Model(&models.EntUser{}).Where("EnterpriseID = ?",enterpriseID).Find(&ent_users)
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
		"error":"nil",
		"msg":ent_users,
	})
}

func ChangeFace(c *gin.Context){
	file,err := c.FormFile("pic")
	if err != nil {
		c.JSON(200,gin.H{
			"code":codes.UpdateError,
			"error":err,
			"msg":"头像上传失败",
		})
		return
	}

	dst := fmt.Sprintf("./pic/user/%s",file.Filename)
	err = c.SaveUploadedFile(file,dst)
	if err != nil {
		c.JSON(200,gin.H{
			"code":codes.UpdateError,
			"error":err,
			"msg":"头像上传失败",
		})
		return
	}

	userId := c.MustGet("UserID")
	var entUser models.EntUser
	dao.DB.Model(&models.EntUser{}).Where("UserID = ?", userId).Find(&entUser)
	err = entUser.ChangeFace(dst)
	if err != nil {
		c.JSON(200,gin.H{
			"code":codes.DBError,
			"error":err,
			"msg":"头像信息更新失败",
		})
		return
	}

	c.JSON(200,gin.H{
		"code":codes.OK,
		"msg":dst,
	})
}

func FindUserInformation(c *gin.Context){
	enterpriseID,_ := strconv.ParseInt(c.PostForm("EnterpriseID"),10,64)
	message := c.PostForm("message")
	infor := "%"+message+"%"
	var x []models.EntUser
	err := dao.DB.Model(models.EntUser{}).
		Where("EnterpriseID = ? AND UserName LIKE ?",enterpriseID,infor).Find(&x).Error
	if err != nil {
		c.JSON(200,gin.H{
			"code":codes.NotData,
			"error":err,
			"msg":"数据信息不存在",
		})
		return
	}
	c.JSON(200,gin.H{
		"code":codes.OK,
		"msg":x,
	})
}

func AddUserRole(c *gin.Context){
	userID,_ := strconv.ParseInt(c.Param("id"),10,64)
	roleID := c.PostForm("roleID")

	err := authority.VerifyPermission(c,codes.AddRolePermission)
	if err != nil {
		return
	}

	var user models.EntUser
	err = dao.DB.Model(&models.EntUser{}).Where("UserID = ?",userID).Find(&user).Error
	if err != nil {
		c.JSON(200,gin.H{
			"code":codes.NotData,
			"error":err,
			"msg":"该用户信息不存在",
		})
		return
	}

	err = dao.DB.Model(&models.EntUser{}).Where("UserID = ?",userID).Update("UserRoleID",user.UserRoleID+
		","+roleID).Error
	if err != nil {
		c.JSON(200,gin.H{
			"code":codes.UpdateError,
			"error":err,
			"msg":"添加角色信息失败",
		})
		return
	}

	c.JSON(200,gin.H{
		"code":codes.OK,
		"msg":user,
	})
}