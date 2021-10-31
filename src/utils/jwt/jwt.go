//token生成

package jwt

import (
	"api.openfileplatform.com/models"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type MyClaims struct{
	userID int64
	jwt.StandardClaims
}

const TokenExpireDuration = time.Hour * 2

var MyKey = []byte("MyNameIsShyHao")

func GetToken(user models.EntUser)(string,error){
	c := MyClaims{
		userID:user.UserID,
			//EnterpriseID:user.EnterpriseID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(), // 过期时间
			Issuer:    "PlutoLove233",                               // 签发人
		},
	}
	token:=jwt.NewWithClaims(jwt.SigningMethodHS256,c)
	return token.SignedString(MyKey)
}

func ParseToken(tokenString string)(*MyClaims,error){
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return MyKey, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

//func JWTAuthMiddleware() func(c *gin.Context) {
//	return func(c *gin.Context) {
//		// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
//		// 这里假设Token放在Header的Authorization中，并使用Bearer开头
//		// 这里的具体实现方式要依据你的实际业务情况决定
//		authHeader := c.Request.Header.Get("Authorization")
//		if authHeader == "" {
//			c.JSON(http.StatusOK, gin.H{
//				"code": 2003,
//				"msg":  "请求头中auth为空",
//			})
//			c.Abort()
//			return
//		}
//		// 按空格分割
//		parts := strings.SplitN(authHeader, " ", 2)
//		if !(len(parts) == 2 && parts[0] == "Bearer") {
//			c.JSON(http.StatusOK, gin.H{
//				"code": 2004,
//				"msg":  "请求头中auth格式有误",
//			})
//			c.Abort()
//			return
//		}
//		// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
//		mc, err := ParseToken(parts[1])
//		if err != nil {
//			c.JSON(http.StatusOK, gin.H{
//				"code": 2005,
//				"msg":  "无效的Token",
//			})
//			c.Abort()
//			return
//		}
//		// 将当前请求的user信息保存到请求的上下文c上
//		c.Set("UserName", mc.UserName)
//		c.Set("UserID",mc.UserID)
//		c.Set("EnterpriseID",mc.EnterpriseID)
//		c.Set("Account",mc.Account)
//		c.Next() // 后续的处理函数可以用过c.Get("username")来获取当前请求的用户信息
//	}
//}