package enterprise

import (
	"api.openfileplatform.com/commons/codes"
	"api.openfileplatform.com/dao"
	"api.openfileplatform.com/models"
	"api.openfileplatform.com/utils/authority"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"path"
	"strconv"
	"time"
)

func getTypeID() func(string) int {
	// innerMap is captured in the closure returned below
	innerMap := map[string]int{
		"txt":10,
		"doc":11,
		"docx":111,
		"ppt":12,
		"pptx":121,
		"xls":13,
		"xlsx":131,
		"pdf":14,
	}

	return func(key string) int {
		return innerMap[key]
	}
}

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

	typeID := getTypeID()(path.Ext(dst))
	file_info := models.EntFileinfo{
		FileAddress:dst,
		FileName:file.Filename,
		TypeID:typeID,
		UpTime:time.Now(),
		BorrowTimes:0,
		Status:0,
		UploaderID:logs.UserID,
	}
	err = file_info.Add()
	if err != nil {
		c.JSON(200,gin.H{
			"code":codes.DBError,
			"err":err,
			"msg":"文件信息上传数据库失败",
		})
		return
	}

	c.JSON(200,gin.H{
		"code":200,
		"err":"nil",
		"msg":dst,
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
		OperationContent:fmt.Sprintf("%d",res.Account)+" upload file",
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
}

func BorrowFile(c *gin.Context){
	fileNumber,err := strconv.ParseInt(c.Param("id"),10,64)
	if err != nil {
		c.JSON(200,gin.H{
			"code":codes.InternetError,
			"error":err,
			"msg":"文档编号获取失败",
		})
		return
	}

	var file models.EntFileinfo
	err = dao.DB.Model(&models.EntFileinfo{}).Where("AutoID = ?", fileNumber).Find(&file).Error
	if err != nil {
		c.JSON(200,gin.H{
			"code":codes.NotData,
			"error":err,
			"msg":"目标文件不存在",
		})
		return
	}

	if authority.CheckAuthority(c.MustGet("UserID").(int64), codes.BorrowFilePermission)==false {
		c.JSON(200,gin.H{
			"code":codes.RoleError,
			"error":"permission error",
			"msg":"用户没有权限访问",
		})
		return
	}

	err = dao.DB.Model(&models.EntFileinfo{}).Where("AutoID = ?",fileNumber).Update("Status",1).Error
	if err != nil {
		c.JSON(200,gin.H{
			"code":codes.DBError,
			"error":err,
			"msg":"借出失败",
		})
		return
	}

	err = dao.DB.Model(&models.EntFileinfo{}).Where("AutoID = ?", fileNumber).Update("BorrowTimes",file.BorrowTimes+1).Error
	if err != nil {
		c.JSON(200,gin.H{
			"code":codes.DBError,
			"error":err,
			"msg":"借出次数改动失败",
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

func DeleteFile(c *gin.Context){
	fileID,err := strconv.ParseInt(c.Param("id"),10,64)
	if err != nil {
		c.JSON(200,gin.H{
			"code":codes.InternetError,
			"error":err,
			"msg":"删除文档编号获取失败",
		})
	}

	var aFile models.EntFileinfo
	err = dao.DB.Model(&models.EntFileinfo{}).Where("AutoID = ?",fileID).Find(&aFile).Error
	if err != nil {
		c.JSON(200,gin.H{
			"code":codes.DBError,
			"error":err,
			"msg":"删除文档不存在",
		})
		return
	}

	dst := aFile.FileAddress

	if authority.CheckAuthority(c.MustGet("UserID").(int64),codes.DeleteFilePermission) == false{
		c.JSON(200,gin.H{
			"code":codes.RoleError,
			"error":"permission error",
			"msg":"用户没有权限删除文档",
		})
		return
	}

	err = os.Remove(dst)
	if err != nil{
		c.JSON(200,gin.H{
			"code":codes.IOError,
			"error":err,
			"msg":"文件删除失败",
		})
		return
	}
	dao.DB.Unscoped().Delete(&aFile)

	c.JSON(200,gin.H{
		"code":codes.OK,
		"error":"nil",
		"msg":"文件删除成功",
	})

	var onelog models.EntUserLog
	reqIP := c.ClientIP()//获取IP
	if reqIP == "::1" {
		reqIP = "127.0.0.1"
	}

	onelog = models.EntUserLog{
		UserID:           c.MustGet("UserID").(int64),
		UserName:         c.MustGet("UserName").(string),
		Account:          c.MustGet("Account").(string),
		OperationIP:      reqIP,
		OperationType:    "3",
		OperationContent: fmt.Sprintf(c.MustGet("Account").(string)+" delete file"),
		OperationResult:  1,
		OperationStatus:  1,
		CreateTime:       time.Time{},
	}
	dao.DB.Model(&models.EntUserLog{}).Create(&onelog)
}