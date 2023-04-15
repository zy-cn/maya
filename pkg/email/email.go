// email
package email

import (
	"maya/configs"
	"net/smtp"
	"net/textproto"
	"strconv"
	"time"

	"github.com/jordan-wright/email"
)

type EmailContent struct {
	Subject string
	Content string
	IsHtml  bool
	To      []string
	Attachs []EmailAttachment
}

type EmailAttachment struct {
	Filename    string
	ContentType string
	Content     []byte
}

var (
	pool     *email.Pool
	auth     smtp.Auth
	hostport string
	from     string
	isPool   bool
)

func InitEmail(config *configs.Config) {
	auth = smtp.PlainAuth("", config.SMTPInfo.UserName, config.SMTPInfo.Password, config.SMTPInfo.Host)
	hostport = config.SMTPInfo.Host + ":" + strconv.Itoa(config.SMTPInfo.Port)
	from = config.SMTPInfo.From
	if config.SMTPInfo.PoolSize > 0 {
		var err error
		pool, err = email.NewPool(
			hostport,
			config.SMTPInfo.PoolSize,
			auth,
		)
		if err != nil {
			panic("无法启动email pool")
		}
		isPool = true
	} else {
		isPool = false
	}
}

func SendEmail(emailContent EmailContent) {
	emailer := email.NewEmail()
	emailer.From = from
	emailer.To = emailContent.To

	emailer.Subject = emailContent.Subject
	emailer.Text = []byte(emailContent.Content)
	// emailer.Attachments = nil

	for _, v := range emailContent.Attachs {
		at := &email.Attachment{
			Filename:    v.Filename,
			ContentType: v.ContentType,
			Header:      textproto.MIMEHeader{},
			Content:     v.Content,
		}
		emailer.Attachments = append(emailer.Attachments, at)
	}

	if isPool {
		pool.Send(emailer, 1*time.Second)
	} else {
		emailer.Send(hostport, auth)
	}

}
