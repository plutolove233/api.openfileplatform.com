package api

import (
	"DocumentSystem/commons/codes"
	"DocumentSystem/utils/email"
	"github.com/gin-gonic/gin"
	"math/rand"
	"time"
)

func SMTPSendEmail(c *gin.Context){
	Sender := "开源电子档案管理系统"
	Receiver := c.PostForm("email")
	Subject := "【电子档案系统】您的验证码是："
	code := MakeVerify(10)
	Context := "<html>" +
		"<body>" +
		"您的验证码是"+"<strong>"+code+"</strong>"+
		"，如果不是您的邮件请忽略它"+
		"</body>" +
		"</html>"

	err := email.SMTPSendEmail(Sender,Receiver,Subject,"html",Context)
	if err != nil {
		c.JSON(200,gin.H{
			"code":codes.InternetError,
			"error":err,
			"msg":"邮件发送错误",
		})
		return
	}

	c.JSON(200,gin.H{
		"code":codes.OK,
		"error":"nil",
		"msg":"邮件发送成功",
	})
}

func MakeVerify(len int )string{
	var letters = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, len)
	r:=rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range b {
		b[i] = letters[r.Intn(62)]
	}
	return string(b)
}