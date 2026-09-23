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
	host     string
	port     int
	username string
	password string
	from     string
}

func NewMailer(host string, port int, user, password, from string) *MailService {
	return &MailService{host, port, user, password, from}
}

func (s *MailService) VerifyEmail(user domain.User, code string) error {
	return s.Send(context.Background(), SendMailParam{
		Template:   "verify_email",
		Recipients: []string{user.Email},
		Args: map[string]any{
			"name": strings.Split(user.Name, " ")[0],
			"code": code,
		},
	})
}

func (s *MailService) ResetPassword(user domain.User, code string) error {
	return s.Send(context.Background(), SendMailParam{
		Template:   "reset_password",
		Recipients: []string{user.Email},
		Args: map[string]any{
			"name": strings.Split(user.Name, " ")[0],
			"code": code,
		},
	})
}

func (s *MailService) Welcome(user domain.User) error {
	return s.Send(context.Background(), SendMailParam{
		Template:   "welcome",
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
	if err := msg.From(m.from); err != nil {
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

	client, err := mail.NewClient(m.host, mail.WithPort(m.port))
	if err != nil {
		return err
	}

	if strings.TrimSpace(m.username) != "" && strings.TrimSpace(m.password) != "" {
		client.SetSMTPAuth(mail.SMTPAuthPlain)
		client.SetUsername(m.username)
		client.SetPassword(m.password)
	}

	return client.DialAndSendWithContext(ctx, msg)
}

func (m *MailService) loadTemplate(filename string, args map[string]any) (string, error) {
	cwd, _ := os.Getwd()

	file1 := path.Join(cwd, "templates", "layout.html")
	file2 := path.Join(cwd, "templates", filename+".html")
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
