package application

import (
	"erosync/internal/config"
	"erosync/internal/handler"
	"erosync/internal/port"
	"erosync/internal/service"
	"erosync/pkg/validator"
)

type Application struct {
	config  *config.Config
	store   *port.Store
	handler *handler.Handler
}

func New(cfg *config.Config, store *port.Store) *Application {
	validator := validator.New()
	userService := service.NewUserService(store.User)
	otpService := service.NewOtpService(store.Otp)
	mailService := service.NewMailService(
		cfg.MailHost,
		cfg.MailPort,
		cfg.MailUsername,
		cfg.MailPassword,
		cfg.MailFrom,
	)
	jwtService := service.NewJwtService(cfg.JwtSecret)

	return &Application{
		config: cfg,
		store:  store,
		handler: handler.New(
			validator,
			userService,
			otpService,
			mailService,
			jwtService,
		),
	}
}
