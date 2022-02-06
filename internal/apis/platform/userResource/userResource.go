// coding: utf-8
// @Author : lryself
// @Date : 2021/12/29 1:11
// @Software: GoLand

package userResource

import (
	"api.openfileplatform.com/internal/dao"
	"api.openfileplatform.com/internal/globals/codes"
	"api.openfileplatform.com/internal/globals/snowflake"
	"api.openfileplatform.com/internal/models/ginModels/platform"
	"api.openfileplatform.com/internal/services"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type UserApiImpl struct{}

type RegisterParser struct {
	UserName string `form:"UserName" json:"UserName" binding:""`
	Account  string `form:"Account" json:"Account" binding:"required"`
	Password string `form:"Password" json:"Password" binding:"required"`
	Phone    string `form:"Phone" json:"Phone" binding:""`
	Email    string `form:"Email" json:"Email" binding:""`
}

func (*UserApiImpl) Register(c *gin.Context) {
	var Parser RegisterParser
	var err error
	//解析参数
	err = c.ShouldBind(&Parser)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    codes.OK,
			"message": "参数错误！",
			"err":     err,
		})
		return
	}

	userInfo := services.PlatUsersService{}

	// 检验此注册方式是否已经注册过
	userInfo.Account = Parser.Account
	err = userInfo.Get()
	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    codes.DataExist,
			"message": "账号已被注册！",
		})
		return
	} else if err.Error() == "record not found" {
		// 未注册过则注册此登录方式
		hash, err := bcrypt.GenerateFromPassword([]byte(Parser.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":    codes.InternalError,
				"message": "密码加密错误！",
				"err":     err,
			})
			return
		}
		userInfo.UserName = Parser.UserName
		userInfo.Account = Parser.Account
		userInfo.Password = string(hash)
		userInfo.UserID = snowflake.GetSnowflakeID()
		userInfo.Phone = Parser.Phone
		userInfo.Email = Parser.Email
		err1 := userInfo.Add()
		if err1 != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":    codes.DBError,
				"message": "数据库错误！",
				"err":     err1,
			})
			return
		}
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    codes.DBError,
			"message": "数据库错误！",
			"err":     err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    codes.OK,
		"message": "注册成功！",
	})
}

func (*UserApiImpl) Get (c *gin.Context) {
	var usersInfo []dao.PlatUsers
	var platusers dao.PlatUsers
	var err error
	err,usersInfo = platusers.GetAll()
	if err != nil {
		c.JSON(http.StatusOK,gin.H{
			"code":codes.DBError,
			"message":"数据库错误",
			"error":err,
		})
		return
	}
	userList := []platform.UserList{}
	list := platform.UserList{}
	for i := 0; i< len(usersInfo); i++{
		list.UserID = usersInfo[i].UserID
		list.UserName = usersInfo[i].UserName
		list.Phone = usersInfo[i].Phone
		list.Email = usersInfo[i].Email
		userList = append(userList,list)
	}

	c.JSON(http.StatusOK,gin.H{
		"code":codes.OK,
		"message":userList,
	})
}