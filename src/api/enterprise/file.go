package enterprise

import (
	"api.openfileplatform.com/commons/codes"
	"api.openfileplatform.com/dao"
	"api.openfileplatform.com/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
	"time"
)

func Upload(c *gin.Context){
	//接收文件
	file,err:= c.FormFile("f1")
	if err != nil{
		c.JSON(200,gin.H{
			"code":codes.InternetError,
			"error":err,
			"msg":"表格请求内容出错",
		})
		return
	}

	var logs models.EntUserLog
	logs.UserID = c.MustGet("userID").(int64)
	var res models.PlatUser
	err = dao.DB.Model(models.PlatUser{}).Where("UserID = ?",logs.UserID).Find(&res).Error
	if err != nil {
		c.JSON(200,gin.H{
			"code":codes.UserError,
			"error":err,
			"msg":"用户不存在",
		})
		return
	}

	//保存文件
	log.Println(file.Filename)
	dst := fmt.Sprintf("./save/%s",file.Filename)
	Err := c.SaveUploadedFile(file,dst)
	if Err != nil{
		c.JSON(200,gin.H{
			"code":codes.IOError,
			"error":Err,
			"msg":"文档保存失败",
		})
		return
	}

	c.JSON(200,gin.H{
		"code":200,
		"path":dst,
	})

	//日志记录

	reqIP := c.ClientIP()//获取IP
	if reqIP == "::1" {
		reqIP = "127.0.0.1"
	}

	logs = models.EntUserLog{
		UserID:c.MustGet("userID").(int64),
		UserName:res.UserName,
		Account:res.Account,
		OperationIP:reqIP,
		OperationType:"1",
		OperationContent:fmt.Sprintf("%d",res.Account)+"upload file",
		OperationResult:1,
		OperationStatus:1,
		CreateTime:time.Now(),
	}
	err = dao.DB.Model(&models.EntUserLog{}).Create(&logs).Error
	if err != nil {
		c.JSON(200,gin.H{
			"code":codes.DBError,
			"error":err,
			"msg":"日志记录失败",
		})
		return
	}

	c.JSON(200,gin.H{
		"code":codes.OK,
		"error":"nil",
		"msg":logs,
	})
}

func BorrowFile(c *gin.Context){
	file_number,err := strconv.ParseInt(c.Param("id"),10,64)
	if err != nil {
		c.JSON(200,gin.H{
			"code":codes.InternetError,
			"error":err,
			"msg":"文档编号获取失败",
		})
		return
	}

	var file models.EntFileinfo
	err = dao.DB.Model(&models.EntFileinfo{}).Where("AutoID = ?",file_number).Find(&file).Error
	if err != nil {
		c.JSON(200,gin.H{
			"code":codes.NotData,
			"error":err,
			"msg":"目标文件不存在",
		})
		return
	}

	err = dao.DB.Model(&models.EntFileinfo{}).Where("AutoID = ?",file_number).Update("BorrowTimes",file.BorrowTimes+1).Error
	if err != nil {
		c.JSON(200,gin.H{
			"code":codes.DBError,
			"error":err,
			"msg":"借出失败",
		})
		return
	}

	c.JSON(200,gin.H{
		"code":codes.OK,
		"error":"nil",
		"msg":"借出成功",
	})
}

func ReturnFile(c *gin.Context){
	file_number,err := strconv.ParseInt(c.Param("id"),10,64)
	if err != nil {
		c.JSON(200,gin.H{
			"code":codes.InternetError,
			"error":err,
			"msg":"文档编号获取失败",
		})
		return
	}

	var file models.EntFileinfo
	err = dao.DB.Model(&models.EntFileinfo{}).Where("AutoID = ?",file_number).Find(&file).Error
	if err != nil {
		c.JSON(200,gin.H{
			"code":codes.NotData,
			"error":err,
			"msg":"目标文件不存在",
		})
		return
	}

	err = dao.DB.Model(&models.EntFileinfo{}).Where("AutoID = ?",file_number).Update("BorrowTimes",file.BorrowTimes-1).Error
	if err != nil {
		c.JSON(200,gin.H{
			"code":codes.DBError,
			"error":err,
			"msg":"归还失败",
		})
		return
	}

	c.JSON(200,gin.H{
		"code":codes.OK,
		"error":"nil",
		"msg":"归还成功",
	})
}