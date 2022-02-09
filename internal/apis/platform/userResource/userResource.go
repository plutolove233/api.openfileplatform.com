// coding: utf-8
// @Author : lryself
// @Date : 2021/12/29 1:11
// @Software: GoLand

package userResource

import (
	"api.openfileplatform.com/internal/globals/codes"
	"api.openfileplatform.com/internal/globals/responseParser"
	"api.openfileplatform.com/internal/models/ginModels"
	"api.openfileplatform.com/internal/services"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)
//platform
type UserApiImpl struct{}

//type RegisterParser struct {
//	UserName string `form:"UserName" json:"UserName" binding:""`
//	Account  string `form:"Account" json:"Account" binding:"required"`
//	Password string `form:"Password" json:"Password" binding:"required"`
//	Phone    string `form:"Phone" json:"Phone" binding:""`
//	Email    string `form:"Email" json:"Email" binding:""`
//}
//
//func (*UserApiImpl) Register(c *gin.Context) {
//	var parser RegisterParser
//	var err error
//	//解析参数
//	err = c.ShouldBind(&parser)
//	if err != nil {
//		responseParser.JsonParameterIllegal(c,err)
//		return
//	}
//
//	userInfo := services.PlatUsersService{}
//
//	// 检验此注册方式是否已经注册过
//	userInfo.Account = parser.Account
//	err = userInfo.Get()
//	if err == nil {
//		responseParser.JsonDataExist(c,"账号已被注册！")
//		return
//	} else if err.Error() == "record not found" {
//		// 未注册过则注册此登录方式
//		hash, err := bcrypt.GenerateFromPassword([]byte(parser.Password), bcrypt.DefaultCost)
//		if err != nil {
//			responseParser.JsonInternalError(c,"密码加密错误！",err)
//			return
//		}
//		userInfo.UserName = parser.UserName
//		userInfo.Account = parser.Account
//		userInfo.Password = string(hash)
//		userInfo.UserID = snowflake.GetSnowflakeID()
//		userInfo.Phone = parser.Phone
//		userInfo.Email = parser.Email
//		err1 := userInfo.Add()
//		if err1 != nil {
//			responseParser.JsonDBError(c,"",err1)
//			return
//		}
//	} else {
//		responseParser.JsonDBError(c,"",err)
//		return
//	}
//
//	c.JSON(http.StatusOK, gin.H{
//		"code":    codes.OK,
//		"message": "注册成功！",
//	})
//}

type userListResponser struct {
	UserName string `form:"UserName" json:"UserName" binding:""`
	UserID   string `form:"UserID" json:"UserID" binding:"required"`
	Phone    string `form:"Phone" json:"Phone" binding:""`
	Email    string `form:"Email" json:"Email" binding:""`
}

func (*UserApiImpl) GetUserList(c *gin.Context) {
	var platusers services.PlatUsersService
	usersInfo,err := platusers.GetAll()
	if err != nil {
		responseParser.JsonDBError(c,"",err)
		return
	}
	userList := []userListResponser{}
	list := userListResponser{}
	for i := 0; i< len(usersInfo); i++{
		list.UserID = usersInfo[i].UserID
		list.UserName = usersInfo[i].UserName
		list.Phone = usersInfo[i].Phone
		list.Email = usersInfo[i].Email
		userList = append(userList,list)
	}

	responseParser.JsonOK(c,userList)
}

type changePasswordParser struct {
	UserID      string `form:"UserID" json:"UserID" binding:"required"`
	Password    string `form:"Password" json:"Password" binding:"required"`
	NewPassword string `form:"NewPassword" json:"NewPassword" binding:"required"`
}

func (*UserApiImpl) ChangePassword(c *gin.Context) {
	var parser changePasswordParser
	var err error
	//解析参数
	err = c.ShouldBind(&parser)
	if err != nil {
		responseParser.JsonParameterIllegal(c, err)
		return
	}

	//查询账号信息
	temp, ok := c.Get("user")
	if !ok {
		responseParser.JsonLoginError(c,"用户未登录",nil)
		return
	}

	user, _ := temp.(ginModels.UserModel)
	if err != nil {
		responseParser.JsonParameterIllegal(c,err)
		return
	}
	userID := parser.UserID

	if user.VerifyAdminRole() {
		if user.UserID != userID || user.IsAdmin != true {
			responseParser.JsonUnauthorizedUserId(c,"只能修改自己的密码！")
			return
		}
	}

	var platUser services.PlatUsersService
	platUser.UserID = userID
	err = platUser.Get()
	if err != nil {
		if err.Error() == "record not found" {
			responseParser.JsonNotData(c,err)
			return
		}
		responseParser.JsonDBError(c,"",err)
		return
	}
	//验证密码
	err = bcrypt.CompareHashAndPassword([]byte(platUser.Password), []byte(parser.Password))
	if err != nil {
		responseParser.JsonDataError(c,"密码错误！")
		return
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(parser.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		responseParser.JsonDBError(c,"",err)
		return
	}

	err = platUser.Update(map[string]interface{}{
		"password":    string(hash),
		//"update_user": user.UserID,
	})
	if err != nil {
		responseParser.JsonDBError(c,"更新密码出错！",err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    codes.OK,
		"message": "修改成功！",
	})
}
