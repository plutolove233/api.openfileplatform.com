package email

import (
	"encoding/base64"
	"fmt"
	"net/smtp"
	"strings"
)

const (
	SMTPHost = "smtp.qq.com"
	SMTPPort = ":587"
	SMTPUser = "1580916438@qq.com"
	SMTPPass = "hbzsqbbfspmjffii"
)

func SMTPSendEmail(userNikeName, to, subject, format, body string) error {
	auth := smtp.PlainAuth("", SMTPUser, SMTPPass, SMTPHost)

	bs64 := base64.NewEncoding("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/")
	header := make(map[string]string)
	header["From"] = userNikeName + "<" + SMTPUser + ">"
	header["To"] = to
	header["Subject"] = fmt.Sprintf("=?UTF-8?B?%s?=", bs64.EncodeToString([]byte(subject)))
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = "text/" + format + "; charset=UTF-8"
	header["Content-Transfer-Encoding"] = "base64"

	data := ""
	for k, v := range header {
		data += k + ": " + v + "\r\n"
	}
	data += "\r\n" + bs64.EncodeToString([]byte(body))
	sendTo := strings.Split(to, ";")

	err := smtp.SendMail(SMTPHost+SMTPPort, auth, SMTPUser, sendTo, []byte(data))
	return err
}