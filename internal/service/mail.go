package service

import (
	"context"
	"erosync/internal/domain"
	"html/template"
	"maps"
	"os"
	"path"
	"strings"
	"time"

	"github.com/wneessen/go-mail"
)

type MailService struct {
	config MailConfig
}

type MailConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	From     string
}

func NewMailService(config MailConfig) *MailService {
	return &MailService{config}
}

func (s *MailService) VerifyEmail(user domain.User, code string) error {
	return s.Send(context.Background(), SendMailParam{
		Template:   "verify_email.html",
		Recipients: []string{user.Email},
		Subject:    "Email Verification Code",
		Args: map[string]any{
			"name":    strings.Split(user.Name, " ")[0],
			"code":    code,
			"minutes": 10,
		},
	})
}

func (s *MailService) ResetPassword(user domain.User, code string) error {
	return s.Send(context.Background(), SendMailParam{
		Template:   "reset_password.html",
		Recipients: []string{user.Email},
		Args: map[string]any{
			"name": strings.Split(user.Name, " ")[0],
			"code": code,
		},
	})
}

func (s *MailService) Welcome(user domain.User) error {
	return s.Send(context.Background(), SendMailParam{
		Template:   "welcome.html",
		Recipients: []string{user.Email},
		Args: map[string]any{
			"name": strings.Split(user.Name, " ")[0],
		},
	})
}

type SendMailParam struct {
	Recipients []string
	Args       map[string]any
	Subject    string
	Template   string
}

func (m *MailService) Send(ctx context.Context, param SendMailParam) error {
	msg := mail.NewMsg()
	if err := msg.From(m.config.From); err != nil {
		return err
	}

	if err := msg.To(param.Recipients...); err != nil {
		return err
	}

	msg.Subject(param.Subject)

	param.Args["subject"] = param.Subject
	bodyString, err := m.loadTemplate(param.Template, param.Args)
	if err != nil {
		return err
	}

	msg.SetBodyString(mail.TypeTextHTML, bodyString)

	client, err := mail.NewClient(m.config.Host, mail.WithPort(m.config.Port))
	if err != nil {
		return err
	}

	if strings.TrimSpace(m.config.Username) != "" && strings.TrimSpace(m.config.Password) != "" {
		client.SetSMTPAuth(mail.SMTPAuthPlain)
		client.SetUsername(m.config.Username)
		client.SetPassword(m.config.Password)
	} else {
		client.SetTLSPolicy(mail.NoTLS)
	}

	return client.DialAndSendWithContext(ctx, msg)
}

func (m *MailService) loadTemplate(filename string, args map[string]any) (string, error) {
	cwd, _ := os.Getwd()

	file1 := path.Join(cwd, "templates", "layout.html")
	file2 := path.Join(cwd, "templates", filename)
	tmpl, err := template.ParseFiles(file1, file2)
	if err != nil {
		panic(err)
	}

	var sb strings.Builder

	tmplArgs := map[string]any{
		"appUrl":  "http://localhost:8080",
		"year":    time.Now().Year(),
		"appName": "Erosync",
	}

	maps.Copy(tmplArgs, args)
	tmpl.Execute(&sb, tmplArgs)

	return sb.String(), nil
}
