package service

import (
	"context"
	"html/template"
	"maps"
	"os"
	"path"
	"strings"
	"time"

	"github.com/telzz/erosync-api/internal/domain"
	"github.com/wneessen/go-mail"
)

type MailConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	MailFrom string
	AppName  string
	AppUrl   string
}

type MailService struct {
	cfg MailConfig
}

func NewMailService(cfg MailConfig) *MailService {
	return &MailService{cfg}
}

type sendMailParam struct {
	Subject      string
	Recipients   []string
	Template     string
	TemplateArgs map[string]any
}

func (s *MailService) SendWelcomeMail(ctx context.Context, user domain.User) error {
	return s.sendMail(ctx, sendMailParam{
		Subject:    "Welcome to " + s.cfg.AppName,
		Recipients: []string{user.Email},
		Template:   "welcome.html",
		TemplateArgs: map[string]any{
			"name": strings.Split(user.Name, " ")[0],
		},
	})
}

func (s *MailService) SendVerificationCode(ctx context.Context, user domain.User, code string, expMin int) error {
	return s.sendMail(ctx, sendMailParam{
		Subject:    "Verify your account",
		Recipients: []string{user.Email},
		Template:   "verify_email.html",
		TemplateArgs: map[string]any{
			"name":    strings.Split(user.Name, " ")[0],
			"code":    code,
			"minutes": expMin,
		},
	})
}

func (s *MailService) sendMail(ctx context.Context, param sendMailParam) error {
	msg := mail.NewMsg()

	if err := msg.From(s.cfg.MailFrom); err != nil {
		return err
	}

	if err := msg.To(param.Recipients...); err != nil {
		return err
	}

	msg.Subject(param.Subject)

	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	var files = []string{
		path.Join(cwd, "templates", "layout.html"),
		path.Join(cwd, "templates", param.Template),
	}

	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		return err
	}

	args := map[string]any{
		"subject": param.Subject,
		"year":    time.Now().Year(),
		"appName": s.cfg.AppName,
		"appUrl":  s.cfg.AppUrl,
	}
	maps.Copy(args, param.TemplateArgs)

	if err := msg.SetBodyHTMLTemplate(tmpl, args); err != nil {
		return err
	}

	c, err := mail.NewClient(s.cfg.Host, mail.WithPort(s.cfg.Port))
	if err != nil {
		return err
	}

	if s.cfg.Username != "" || s.cfg.Password != "" {
		c.SetUsername(s.cfg.Username)
		c.SetPassword(s.cfg.Password)
	} else {
		c.SetTLSPolicy(mail.NoTLS)
	}

	// ctx, cancel := context.WithTimeout(ctx, 20*time.Second)
	// defer cancel()

	return c.DialAndSend(msg)
}
