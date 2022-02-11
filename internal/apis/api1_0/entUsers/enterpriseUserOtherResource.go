package entUsers

import (
	"api.openfileplatform.com/internal/globals/codes"
	"api.openfileplatform.com/internal/globals/responseParser"
	"api.openfileplatform.com/internal/globals/snowflake"
	"api.openfileplatform.com/internal/models/ginModels"
	"api.openfileplatform.com/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type RegisterParser struct {
	UserName string `form:"UserName" json:"UserName" binding:""`
	Account  string `form:"Account" json:"Account" binding:"required"`
	Password string `form:"Password" json:"Password" binding:"required"`
	Phone    string `form:"Phone" json:"Phone" binding:""`
	Email    string `form:"Email" json:"Email" binding:""`
}

func (*EnterpriseUserApi) Register(c *gin.Context) {
	var Parser RegisterParser
	var err error
	//解析参数
	err = c.ShouldBind(&Parser)
	if err != nil {
		responseParser.JsonParameterIllegal(c, "", err)
		return
	}

	var entUsersService services.EntUserService

	// 检验此注册方式是否已经注册过
	entUsersService.Account = Parser.Account
	err = entUsersService.Get()
	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    codes.DataExist,
			"message": "账号已被注册！",
		})
		return
	} else if err.Error() != "record not found" {
		responseParser.JsonDBError(c, "", err)
	}
	// 未注册过则注册此登录方式
	hash, err := bcrypt.GenerateFromPassword([]byte(Parser.Password), bcrypt.DefaultCost)
	if err != nil {
		responseParser.JsonInternalError(c, "密码加密错误", err)
		return
	}
	entUsersService.UserName = Parser.UserName
	entUsersService.Account = Parser.Account
	entUsersService.Password = string(hash)
	entUsersService.UserID = snowflake.GetSnowflakeID()
	entUsersService.Phone = Parser.Phone
	entUsersService.Email = Parser.Email
	err1 := entUsersService.Add()
	if err1 != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    codes.DBError,
			"message": "数据库错误！",
			"err":     err1,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    codes.OK,
		"message": "注册成功！",
	})
}

type changePasswordParser struct {
	UserID      string `form:"UserID" json:"UserID" binding:"required"`
	Password    string `form:"Password" json:"Password" binding:"required"`
	NewPassword string `form:"NewPassword" json:"NewPassword" binding:"required"`
}

func (*EnterpriseUserApi) ChangePassword(c *gin.Context) {
	var Parser changePasswordParser
	var err error
	//解析参数
	err = c.ShouldBind(&Parser)
	if err != nil {
		responseParser.JsonParameterIllegal(c, "", err)
		return
	}

	//查询账号信息
	temp, ok := c.Get("user")
	if !ok {
		responseParser.JsonNotData(c, "用户未登录", nil)
		return
	}

	user := temp.(ginModels.UserModel)
	userID := Parser.UserID

	if !user.VerifyAdminRole() {
		if user.UserID != userID {
			c.JSON(http.StatusOK, gin.H{
				"code":    codes.UnauthorizedUserId,
				"message": "只能修改自己的密码！",
			})
			return
		}
	}

	var entUser services.EntUserService
	entUser.UserID = userID
	err = entUser.Get()
	if err != nil {
		responseParser.JsonDBError(c, "", err)
		return
	}
	//验证密码
	err = bcrypt.CompareHashAndPassword([]byte(entUser.Password), []byte(Parser.Password))
	if err != nil {
		responseParser.JsonDataError(c, "密码错误！", err)
		return
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(Parser.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		responseParser.JsonDBError(c, "", err)
		return
	}

	err = entUser.Update(map[string]interface{}{
		"password": string(hash),
	})
	if err != nil {
		responseParser.JsonDBError(c, "密码出错", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    codes.OK,
		"message": "修改成功！",
	})
}

type refreshPasswordParser struct {
	UserID string `form:"UserID" json:"UserID" binding:"required"`
}

func (*EnterpriseUserApi) RefreshPassword(c *gin.Context) {
	var Parser refreshPasswordParser
	var err error
	//解析参数
	err = c.ShouldBind(&Parser)
	if err != nil {
		responseParser.JsonParameterIllegal(c, "", err)
		return
	}

	//查询账号信息
	temp, ok := c.Get("user")
	if !ok {
		responseParser.JsonNotData(c, "用户未登录", nil)
		return
	}

	user := temp.(ginModels.UserModel)
	userID := Parser.UserID

	if !user.VerifyAdminRole() {
		if user.UserID != userID {
			c.JSON(http.StatusOK, gin.H{
				"code":    codes.UnauthorizedUserId,
				"message": "只能修改自己的密码！",
			})
			return
		}
	}

	var entUser services.EntUserService
	entUser.UserID = userID
	err = entUser.Get()
	if err != nil {
		responseParser.JsonDBError(c, "", err)
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(viper.GetString("user.defaultPassword")), bcrypt.DefaultCost)
	if err != nil {
		responseParser.JsonDBError(c, "", err)
		return
	}

	err = entUser.Update(map[string]interface{}{
		"password": string(hash),
	})
	if err != nil {
		responseParser.JsonDBError(c, "密码出错", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    codes.OK,
		"message": "密码重置成功！",
	})
}

type EnterpriseUserList struct {
	UserID 		string	`form:"UserID" json:"UserID" binding:"required"`
	UserName 	string	`form:"UserName" json:"UserName" binding:"required"`
	Phone		string	`form:"Phone" json:"Phone" binding:"required"`
	Email 		string	`form:"Email" json:"Email" binding:"required"`
}
//获取企业用户信息表
func (*EnterpriseUserApi) GetAllUsersList(c *gin.Context) {
	 temp,ok := c.Get("user")
	if !ok {
		responseParser.JsonNotData(c,"用户未登录",nil)
		return
	}

	user := temp.(ginModels.UserModel)
	if !user.IsAdmin {
		c.JSON(http.StatusOK, gin.H{
			"code":    codes.UnauthorizedUserId,
			"message": "只能由管理员查看用户列表",
		})
		return
	}
}