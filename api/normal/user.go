package normal

import (
	"DocumentSystem/dao"
	"DocumentSystem/models"
	"DocumentSystem/myJWT"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var user models.NormalUsers
	c.ShouldBind(&user)
	token, _ := myJWT.GetToken(user.Name)
	if user.Login(){
		fmt.Println("you login successed")
		c.JSON(200,gin.H{
			"code":200,
			"data":user,
			"token":token,
		})
	}else{
		fmt.Println("you login failed")
		c.JSON(200,gin.H{
			"code":400,
			"msg":"账号或密码错误",
		})
	}
}

func Register(c *gin.Context){
	var user models.NormalUsers
	c.ShouldBind(&user)
	err,ok := user.Register()
	if err!=nil{
		c.JSON(200,gin.H{
			"code":400,
			"error":err,
		})
		return
	}
	if ok{
		c.JSON(200,user)
	}else{
		c.JSON(200, gin.H{
			"code":400,
			"error":"the nickname has been existed",
			"msg":user,
		})
	}
}

func View(c *gin.Context) {
	var user []models.NormalUsers
	err := dao.DB.Find(&user).Error
	if err!=nil{
		c.JSON(200,gin.H{
			"code":405,
			"msg":user,
			"error":err,
		})
		return
	}
	c.JSON(200,gin.H{
		"code":200,
		"msg":user,
	})
}

func Upload(c *gin.Context){

}