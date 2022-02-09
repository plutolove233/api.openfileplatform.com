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

func JsonOK(c *gin.Context, msg string, data interface{}) {
	if msg == "" {
		msg = "成功!"
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    codes.OK,
		"message": msg,
		"data":    data,
	})
}

func JsonParameterIllegal(c *gin.Context, msg string, err error) {
	if msg == "" {
		msg = "参数非法!"
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    codes.ParameterIllegal,
		"message": msg,
		"err":     err.Error(),
	})
}

func JsonDataError(c *gin.Context, msg string) {
	if msg == "" {
		msg = "数据错误!"
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    codes.DataError,
		"message": msg,
	})
}

func JsonNotData(c *gin.Context, msg string, err error) {
	if msg == "" {
		msg = "无数据!"
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    codes.DataError,
		"message": msg,
		"err":     err.Error(),
	})
}

func JsonInternalError(c *gin.Context, msg string, err error) {
	if msg == "" {
		msg = "系统错误!"
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    codes.InternalError,
		"message": msg,
		"err":     err.Error(),
	})
	return
}

func JsonDBError(c *gin.Context, msg string, err error) {
	if err.Error() == "record not found" {
		if msg == "" {
			msg = "无数据!"
		}
		c.JSON(http.StatusOK, gin.H{
			"code":    codes.NotData,
			"message": msg,
			"err":     err.Error(),
		})
		return
	}
	if msg == "" {
		msg = "数据库错误!"
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    codes.DBError,
		"message": msg,
		"err":     err.Error(),
	})
}
