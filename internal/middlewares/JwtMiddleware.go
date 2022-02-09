// coding: utf-8
// @Author : lryself
// @Date : 2021/4/8 11:36
// @Software: GoLand

package middlewares

import (
	"api.openfileplatform.com/internal/globals/codes"
	"api.openfileplatform.com/internal/models/ginModels"
	"api.openfileplatform.com/internal/utils/jwt"
	"api.openfileplatform.com/internal/utils/logs"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

var log = logs.GetLogger()

func TokenRequire() gin.HandlerFunc {
	return func(c *gin.Context) {
		//fullPath := c.FullPath()
		//r,err := regexp.MatchString("/api([a-z/]*/ping|/login|[a-z/]*/register)$", fullPath)
		//if err != nil {
		//	c.JSON(http.StatusOK, gin.H{
		//	    "code": codes.InternalError,
		//	    "message": "正则表达式错误！",
		//	    "err": err,
		//	})
		//	return
		//}
		//if r {
		//	c.Next()
		//	return
		//}

		//token验证
		token := c.Request.Header.Get("Token")
		jwtChaim, err := jwt.VerifyToken(token)
		if err != nil {
			log.Errorln(err)
			c.JSON(http.StatusOK, gin.H{
				"code":    codes.AccessDenied,
				"message": "您的Token已过期！",
			})
			c.Abort()
			return
		}

		//tokenID := jwtChaim.TokenID
		//验证是否与session中tokenID相同
		//session := sessions.Default(c)
		//temp := session.GetUserList("tokenID")
		//if temp == nil {
		//	c.JSON(http.StatusOK, gin.H{
		//		"code":    codes.AccessDenied,
		//		"message": "您的Token非法！",
		//	})
		//	c.Abort()
		//	return
		//}
		//tempTokenID := temp.(string)
		//if tempTokenID != tokenID {
		//	c.JSON(http.StatusOK, gin.H{
		//		"code":    codes.AccessDenied,
		//		"message": "您的Token非法！",
		//	})
		//	c.Abort()
		//	return
		//}

		//从数据库读取token信息
		//redisManager, ctx := database.GetRedisManager()
		//result, err := redisManager.GetUserList(ctx, "Token_"+tokenID).Result()
		//if err != nil {
		//	log.Errorln(err)
		//	c.JSON(http.StatusOK, gin.H{
		//		"code":    codes.AccessDenied,
		//		"message": "您的Token已失效！",
		//	})
		//	c.Abort()
		//	return
		//}

		//刷新token有效期
		//err = redisManager.Expire(ctx, "Token_"+tokenID, time.Duration(viper.GetInt("system.RedisExpireTime"))*time.Second).Err()
		//if err != nil {
		//	c.JSON(http.StatusOK, gin.H{
		//		"code":    codes.InternalError,
		//		"message": "刷新token错误！",
		//		"err":     err,
		//	})
		//	c.Abort()
		//	return
		//}
		//c.Set("TokenID", tokenID)
		//加载用户信息到上下文
		temp, err := json.Marshal(jwtChaim)

		var User ginModels.UserModel
		err = json.Unmarshal(temp, &User)
		if err != nil {
			log.Errorln(err)
			c.JSON(http.StatusOK, gin.H{
				"code":    codes.InternalError,
				"message": "用户信息读取错误！",
			})
			c.Abort()
			return
		}
		c.Set("user", User)
		c.Next()
	}
}
