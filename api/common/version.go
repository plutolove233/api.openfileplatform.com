package common

import "github.com/gin-gonic/gin"

func Version(c *gin.Context){
	c.JSON(200,gin.H{
		"code":200,
		"message":"OK",
		"data":gin.H{
			"version":"v1.0",
		},
	})
}