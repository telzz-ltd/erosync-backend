package application

import (
	"erosync/internal/config"
	"erosync/internal/handler"
	"erosync/internal/port"
	"erosync/internal/service"
	"erosync/pkg/validator"
)

type Application struct {
	Config  *config.Config
	Store   *port.Store
	Handler *handler.Handler

	//services
	Users *service.UserService
	Otps  *service.OtpService
	Mail  *service.MailService
	Jwt   *service.JwtService

	Validator *validator.Validator
}

func New(cfg *config.Config, store *port.Store) *Application {
	app := &Application{
		Config: cfg,
		Store:  store,

		Validator: validator.New(),
		Users:     service.NewUserService(store.User),
		Otps:      service.NewOtpService(store.Otp),
		Mail: service.NewMailService(service.MailConfig{
			Host:     cfg.MailHost,
			Port:     cfg.MailPort,
			Username: cfg.MailUsername,
			Password: cfg.MailPassword,
			From:     cfg.MailFrom,
		}),
		Jwt: service.NewJwtService(cfg.JwtSecret),
	}

	app.Handler = handler.New(
		app.Validator,
		app.Users,
		app.Otps,
		app.Mail,
		app.Jwt,
	)

	return app
}
