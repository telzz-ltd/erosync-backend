package notification

import (
	"context"
	"erosync/internal/model"
	"erosync/internal/shared/config"
	"erosync/templates"
	"os"
	"strings"

	"github.com/wneessen/go-mail"
)

type SendEmailParam struct {
	Subject    string
	Recipients []string
	Content    string
}

func SendEmail(param SendEmailParam) error {
	from := "Erosync Limited <support@erosyncng.com>"
	server := config.GetString("MAIL_HOST", "127.0.0.1")
	port := config.GetInt("MAIL_PORT", 1025)

	m := mail.NewMsg()
	if err := m.From(from); err != nil {
		return err
	}

	if err := m.To(param.Recipients...); err != nil {
		return err
	}

	m.Subject(param.Subject)
	m.SetBodyString(mail.TypeTextHTML, param.Content)

	c, err := mail.NewClient(server, mail.WithPort(port))
	if err != nil {
		return err
	}

	if os.Getenv("APP_ENV") != "production" {
		c.SetTLSPolicy(mail.NoTLS)
	}

	return c.DialAndSend(m)
}

func SendEmailVerificationMail(user model.User, code string) error {
	var c strings.Builder
	err := templates.VerificationCode(user, code).Render(context.Background(), &c)
	if err != nil {
		return err
	}

	return SendEmail(SendEmailParam{
		Subject:    "Email Verification Code",
		Recipients: []string{user.Email},
		Content:    c.String(),
	})
}
