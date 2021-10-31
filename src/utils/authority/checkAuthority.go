package authority

import (
	"api.openfileplatform.com/commons/codes"
	"api.openfileplatform.com/dao"
	"api.openfileplatform.com/models"
	"errors"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

func CheckAuthority(userID int64,match int64,enterpriseID int64) bool{
	var user models.EntUser
	err :=dao.DB.Model(&models.EntUser{}).Where("UserID = ? AND EnterpriseID = ?",userID,enterpriseID).Find(&user).Error
	if err != nil {
		return false
	}

	var userRole [32]int64
	strArr := strings.FieldsFunc(user.UserRoleID,func(r rune)bool{
		if r == ',' {
			return true
		}
		return false
	})

	for i,x:=range strArr {
		userRole[i], _ = strconv.ParseInt(x, 10, 64)
	}

	var userAuthority []models.EntRoleAuthority
	for _,aRole:=range userRole{
		err = dao.DB.Model(&models.EntRoleAuthority{}).Where("RoleID = ? AND EnterpriseID = ?",
			aRole,enterpriseID).Find(&userAuthority).Error
		if err != nil {
			return false
		}
		for _,oneAuthority := range userAuthority{
			if oneAuthority.AuthorityID == match {
				return true
			}
			if oneAuthority.AuthorityID == codes.Admin{
				return true
			}
		}
	}
	return false
}

func VerifyPermission(ctx *gin.Context,permissionID int64) error {
	userID,err := strconv.ParseInt(ctx.PostForm("UserID"),10,64)
	if err!=nil{
		ctx.JSON(200,gin.H{
			"code":codes.ParamError,
			"error":err,
			"msg":"userID获取失败",
		})
		return errors.New("get userID failed")
	}

	enterpriseID,err := strconv.ParseInt(ctx.PostForm("EnterpriseID"),10,64)
	if err != nil {
		ctx.JSON(200,gin.H{
			"code":codes.ParamError,
			"error":err,
			"msg":"企业ID获取失败",
		})
		return errors.New("get enterpriseID failed")
	}
	if !CheckAuthority(userID, permissionID,enterpriseID) {
		ctx.JSON(200, gin.H{
			"code": codes.UserError,
			"msg":  "用户没有权限访问",
		})
		return errors.New("user does not have permission to access")
	}
	return nil
}