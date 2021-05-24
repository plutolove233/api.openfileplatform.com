package ent

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func Upload(c *gin.Context){
	file,err:= c.FormFile("f1")
	if err != nil{
		c.JSON(200,gin.H{
			"code":401,
			"error":err,
		})
		return
	}
	log.Println(file.Filename)
	dst := fmt.Sprintf("./save/%s",file.Filename)
	Err := c.SaveUploadedFile(file,dst)
	if Err != nil{
		c.JSON(200,gin.H{
			"code":408,
			"error":Err,
		})
		return
	}
	c.JSON(200,gin.H{
		"code":200,
		"path":dst,
	})
}

