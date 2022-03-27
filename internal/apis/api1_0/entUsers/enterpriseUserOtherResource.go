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
	"time"
)

type RegisterParser struct {
	UserName     string `form:"UserName" json:"UserName" binding:"required"`
	EnterpriseID string `form:"EnterpriseID" json:"EnterpriseID" binding:"required"`
	Account      string `form:"Account" json:"Account" binding:"required"`
	Password     string `form:"Password" json:"Password" binding:"required"`
	Phone        string `form:"Phone" json:"Phone" binding:""`
	Email        string `form:"Email" json:"Email" binding:""`
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
	entUsersService.EnterpriseID = Parser.EnterpriseID
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
	UserID   string `form:"UserID" json:"UserID" binding:"required"`
	UserName string `form:"UserName" json:"UserName" binding:"required"`
	Phone    string `form:"Phone" json:"Phone" binding:"required"`
	Email    string `form:"Email" json:"Email" binding:"required"`
	IsAdmin  bool   `form:"IsAdmin" json:"IsAdmin" binding:"required"`
}

//获取企业用户信息表
func (*EnterpriseUserApi) GetAllUsersList(c *gin.Context) {
	temp, ok := c.Get("user")
	if !ok {
		responseParser.JsonNotData(c, "用户未登录", nil)
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
	entAdmin := services.EntUserService{}
	entAdmin.UserID = user.UserID
	err := entAdmin.Get()
	if err != nil {
		responseParser.JsonNotData(c, "获取用户企业ID失败", err)
		return
	}
	entUsers, err := entAdmin.GetAll(entAdmin.EnterpriseID)
	if err != nil {
		responseParser.JsonDBError(c, "", err)
		return
	}

	var enterpriseUserList []EnterpriseUserList
	var oneUser EnterpriseUserList
	for _, x := range entUsers {
		oneUser.UserID = x.UserID
		oneUser.UserName = x.UserName
		oneUser.Email = x.Email
		oneUser.Phone = x.Phone
		oneUser.IsAdmin = x.IsAdmin
		enterpriseUserList = append(enterpriseUserList, oneUser)
	}
	responseParser.JsonOK(c, "企业用户列表", enterpriseUserList)
}

type setAdminParser struct {
	UserID string `json:"UserID" form:"UserID" binding:"required"`
}

func (*EnterpriseUserApi) SetAdmin(c *gin.Context) {
	var parser setAdminParser
	err := c.ShouldBind(&parser)
	if err != nil {
		responseParser.JsonParameterIllegal(c, "获取用户id失败", err)
		return
	}

	entUser := services.EntUserService{}
	entUser.UserID = parser.UserID
	if err = entUser.Get(); err != nil {
		responseParser.JsonNotData(c, "该用户信息不存在", err)
		return
	}

	if err = entUser.Update(map[string]interface{}{
		"IsAdmin": 1,
	}); err != nil {
		responseParser.JsonDBError(c, "设置用户为管理员失败", err)
	}

	c.JSON(200, gin.H{
		"code":    codes.OK,
		"message": "设置管理员成功",
	})
}

type removeAdminParser struct {
	UserID       string `json:"UserID" form:"UserID" binding:"required"`
	EnterpriseID string `json:"EnterpriseID" form:"EnterpriseID" binding:"required"`
}

func (*EnterpriseUserApi) RemoveAdmin(c *gin.Context) {
	var parser removeAdminParser
	err := c.ShouldBind(&parser)
	if err != nil {
		responseParser.JsonParameterIllegal(c, "获取信息失败", err)
		return
	}

	temp, _ := c.Get("user")
	user := temp.(ginModels.UserModel)
	if user.IsPlatUser == false {
		entUser := services.EntUserService{}
		entUser.UserID = user.UserID
		err1 := entUser.Get()
		if err1 != nil {
			responseParser.JsonNotData(c, "用户不存在", err1)
			return
		}
		if entUser.EnterpriseID != parser.EnterpriseID {
			responseParser.JsonAccessDenied(c, "用户无法修改其他公司的管理员信息")
			return
		}
		platEnterprise := services.PlatEnterpriseService{}
		platEnterprise.EnterpriseID = parser.EnterpriseID
		err1 = platEnterprise.Get()
		if err1 != nil {
			responseParser.JsonNotData(c,"该企业信息不存在",err1)
			return
		}
		if platEnterprise.AdminID != user.UserID{
			responseParser.JsonAccessDenied(c,"用户没有权限修改管理员信息")
			return
		}
	}

	entUser := services.EntUserService{}
	entUser.UserID = parser.UserID
	err = entUser.Get()
	if err != nil {
		responseParser.JsonNotData(c, "用户不存在", err)
		return
	}
	err = entUser.Update(map[string]interface{}{
		"IsAdmin": false,
	})
	if err != nil {
		responseParser.JsonDBError(c, "修改用户权限失败", err)
		return
	}
	c.JSON(200, gin.H{
		"code":    codes.OK,
		"message": "修改用户权限成功",
	})
}

type getUserLendParser struct {
	BorrowerID	string		`json:"BorrowerID"`
	BorrowTime	time.Time	`json:"BorrowTime"`
	BorrowTerm	int8		`json:"BorrowTerm"`
}

func (*EnterpriseUserApi) GetUserLendInformation(c *gin.Context)  {
	temp, ok := c.Get("user")
	if !ok {
		responseParser.JsonNotData(c,"用户未登录",nil)
		return
	}
	user := temp.(ginModels.UserModel)
	fileLend := services.EnterpriseFileLendService{}
	fileLend.BorrowerID = user.UserID
	data, err := fileLend.GetAllLendInfo()
	if err != nil {
		if err.Error()=="record not found"{
			responseParser.JsonNotData(c,"用户借阅信息不存在",err)
			return
		}
		responseParser.JsonDBError(c,"数据库错误",err)
		return
	}
	parser := []getUserLendParser{}
	for _,item := range data{
		x := getUserLendParser{
			BorrowerID: item.BorrowerID,
			BorrowTime: item.BorrowTime,
			BorrowTerm: item.BorrowTerm,
		}
		parser = append(parser, x)
	}
	responseParser.JsonOK(c,"获取用户借阅信息成功",parser)
}