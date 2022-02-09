// coding: utf-8
// @Author : lryself
// @Date : 2022/1/17 10:43
// @Software: GoLand

package responseParser

import (
	"api.openfileplatform.com/internal/globals/codes"
	"github.com/gin-gonic/gin"
	"net/http"
)

func JsonOK(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code":    codes.OK,
		"message": "成功!",
		"data":    data,
	})
}

func JsonParameterIllegal(c *gin.Context, err error) {
	c.JSON(http.StatusOK, gin.H{
		"code":    codes.ParameterIllegal,
		"message": "参数不合法!",
		"err":     err.Error(),
	})
}

func JsonDataError(c *gin.Context, msg string) {
	if msg != "" {
		c.JSON(http.StatusOK, gin.H{
			"code":    codes.DataError,
			"message": "数据错误: " + msg,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    codes.DataError,
			"message": "数据错误！",
		})
	}
}

func JsonNotData(c *gin.Context, err error) {
	c.JSON(http.StatusOK, gin.H{
		"code":    codes.DataError,
		"message": "无数据！",
		"err":     err.Error(),
	})
}

func JsonInternalError(c *gin.Context, msg string, err error) {
	c.JSON(http.StatusOK, gin.H{
		"code":    codes.InternalError,
		"message": msg,
		"err":     err.Error(),
	})
}

func JsonDBError(c *gin.Context,msg string, err error) {
	if msg != "" {
		c.JSON(http.StatusOK, gin.H{
			"code":    codes.DBError,
			"message": msg,
			"err":     err.Error(),
		})
		return
	}
	if err.Error() == "record not found" {
		c.JSON(http.StatusOK, gin.H{
			"code":    codes.NotData,
			"message": "无数据!",
			"err":     err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    codes.DBError,
		"message": "数据库错误!",
		"err":     err.Error(),
	})
}

func JsonDataExist(c *gin.Context, msg string){
	c.JSON(http.StatusOK,gin.H{
		"code":codes.DataExist,
		"message":msg,
	})
}

func JsonAccessDenied(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, gin.H{
		"code":    codes.AccessDenied,
		"message": msg,
	})
}

func JsonLoginError(c *gin.Context, msg string, err error) {
	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    codes.LoginError,
			"message": msg,
		})
	}else {
		c.JSON(http.StatusOK, gin.H{
			"code":    codes.LoginError,
			"message": msg,
			"err":     err,
		})
	}
}

func JsonUnauthorizedUserId(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, gin.H{
		"code":    codes.UnauthorizedUserId,
		"message": msg,
	})
}

func JsonIncompleteRequest(c *gin.Context,msg string){
	c.JSON(http.StatusOK,gin.H{
		"code":		codes.ParameterIllegal,
		"message":	msg,
	})
}