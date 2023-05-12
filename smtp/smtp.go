package smtp

import (
	"bytes"
	"fmt"
	"logs-monitoring/utils"
	"mime/quotedprintable"
	"net/smtp"
)

var log = utils.InitLogger()

func SendNotification(level, msg string) {
	const (
		SmtpServer = "smtp.mail.ru"

		SmtpPort = 587

		FromEmail = "lalalalalala98912@mail.ru"

		ToEmail = "agfgf11@mail.ru"

		Password = "UDfbFs4mV9fnh9BPqnCr"

		Subject = "Logs Notification"

		MimeVersion = "1.0"

		ContentType = "text/plain; charset=\"utf-8\""

		ContentTransferEncoding = "quoted-printable"
	)

	headers := make(map[string]string)
	headers["Level"] = level
	headers["Description"] = msg
	headers["From"] = FromEmail
	headers["To"] = ToEmail
	headers["Subject"] = Subject
	headers["MIME-Version"] = MimeVersion
	headers["Content-Type"] = ContentType
	headers["Content-Transfer-Encoding"] = ContentTransferEncoding

	message := ""
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}

	message += "\r\n" + quotedprintableEncode(message)

	auth := smtp.PlainAuth("", FromEmail, Password, SmtpServer)
	err := smtp.SendMail(fmt.Sprintf("%s:%d", SmtpServer, SmtpPort), auth, FromEmail,
		[]string{ToEmail}, bytes.NewBufferString(message).Bytes())

	if err != nil {
		log.Fatal("Error send", err.Error())
	}

}

func quotedprintableEncode(str string) string {
	var buf bytes.Buffer
	qpw := quotedprintable.NewWriter(&buf)
	qpw.Write([]byte(str))
	qpw.Close()
	return buf.String()
}
