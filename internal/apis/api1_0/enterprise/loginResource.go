package enterprise

import (
	"api.openfileplatform.com/internal/globals/codes"
	"api.openfileplatform.com/internal/globals/responseParser"
	"api.openfileplatform.com/internal/models/ginModels/enterprise"
	"api.openfileplatform.com/internal/services"
	"api.openfileplatform.com/internal/utils/jwt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type LoginApiImpl struct {}

type loginByPasswordParser struct {
	Account string	`form:"Account" json:"Account" binding:"required"`
	Password string `form:"Password" json:"Password" binding:"required"`
}

func (*LoginApiImpl) LoginByPassword(c *gin.Context)  {
	var parser loginByPasswordParser
	err := c.ShouldBindJSON(&parser)
	if err != nil {
		responseParser.JsonParameterIllegal(c,err)
		return
	}

	user := enterprise.UserModel{}
	token := ""

	var entUser services.EntUserService
	entUser.Account = parser.Account
	err = entUser.Get()
	if err != nil {
		responseParser.JsonDBError(c,err)
		return
	}

	var pwd []byte
	pwd = []byte(parser.Password)
	err = bcrypt.CompareHashAndPassword([]byte(entUser.Pwd),pwd)

	if err != nil {
		if err.Error() == "crypto/bcrypt: hashedPassword is not the hash of the given password" {
			c.JSON(http.StatusOK, gin.H{
				"code":    codes.AccessDenied,
				"message": "密码错误！",
			})
			return
		}
		responseParser.JsonInternalError(c, err)
		return
	}

	user.UserID = entUser.UserID
	user.UserName = entUser.UserName
	user.IsEntUser = true
	user.Account = entUser.Account
	user.IsAdmin = entUser.IsAdmin
	user.FacePicUrl = entUser.FacePicURL

	var tag bool
	if user.IsAdmin == 0 {
		tag = false
	} else {
		tag = true
	}

	token,err = jwt.MakeToken(user.UserID,user.IsEntUser,tag)
	if err != nil {
		responseParser.JsonInternalError(c,err)
		return
	}
	user.Token = token
	err = entUser.Update(map[string]interface{}{"Token": token})
	if err != nil {
		c.JSON(200,gin.H{
			"code":codes.DBError,
			"message":"token更新失败",
			"error":err.Error(),
		})
		return
	}
	responseParser.JsonOK(c,user)
}

type refushTokenParser struct {
	Token string `form:"Token" json:"Token" binding:"required"`
}

func (*LoginApiImpl) RefushToken(c *gin.Context) {
	var parser refushTokenParser
	err := c.ShouldBindJSON(&parser)
	if err != nil {
		responseParser.JsonParameterIllegal(c,err)
		return
	}

	token := parser.Token
	token,err = jwt.RefreshToken(token)

	if err != nil {
		c.JSON(http.StatusOK,gin.H{
			"code":codes.AccessDenied,
			"message":"token已过期",
			"error":err.Error(),
		})
		return
	}

	responseParser.JsonOK(c,token)
}