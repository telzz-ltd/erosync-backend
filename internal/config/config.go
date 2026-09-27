package config

type Config struct {
	AppName     string
	FrontendUrl string
	Port        int

	//Mail
	MailHost     string
	MailPort     int
	MailUsername string
	MailPassword string
	MailFrom     string

	//Database
	DatabaseUrl string

	JwtSecret string
}

func New() *Config {
	return &Config{
		Port:        GetEnv("PORT", 8080),
		AppName:     GetEnv("APP_NAME", "Erosync"),
		FrontendUrl: MustGetEnv[string]("FRONTEND_URL"),

		//Mail
		MailHost:     MustGetEnv[string]("MAIL_HOST"),
		MailPort:     MustGetEnv[int]("MAIL_PORT"),
		MailUsername: GetEnv("MAIL_USERNAME", ""),
		MailPassword: GetEnv("MAIL_PASSWORD", ""),
		MailFrom:     MustGetEnv[string]("MAIL_FROM"),

		//Database
		DatabaseUrl: MustGetEnv[string]("DATABASE_URL"),

		JwtSecret: MustGetEnv[string]("JWT_SECRET"),
	}
}
