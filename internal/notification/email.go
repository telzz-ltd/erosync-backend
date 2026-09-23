package notification

import (
	"context"
	"html/template"
	"maps"
	"os"
	"path"
	"strings"
	"time"

	"github.com/wneessen/go-mail"
)

const (
	VerifyEmail   = "verify_email"
	ResetPassword = "reset_password"
	Welcome       = "welcome"
)

type Mailer struct {
	host     string
	port     int
	username string
	password string
	from     string
}

func NewMailer(host string, port int, user, password, from string) *Mailer {
	return &Mailer{host, port, user, password, from}
}

type SendEmailParams struct {
	Recipients []string
	Args       map[string]any
	Subject    string
	Template   string
}

func (m *Mailer) Send(ctx context.Context, param SendEmailParams) error {
	msg := mail.NewMsg()
	if err := msg.From(m.from); err != nil {
		return err
	}

	if err := msg.To(param.Recipients...); err != nil {
		return err
	}

	msg.Subject(param.Subject)
	msg.SetBodyString(mail.TypeTextHTML, "")

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

func (m *Mailer) loadTemplate(filename string, args map[string]any) (string, error) {
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
