package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	Port         string
	DbUser       string
	DbPass       string
	DbName       string
	DbPort       string
	JwtSecret    string
	Mail         string
	MailPassword string
	SmtpHost     string
	SmtpPort     string
	AppDomain    string
}

var BuildExist bool
var Env *Config

func LoadEnv() {
	Env = &Config{}

	created := CheckEnv()
	if !created {
		EnvHelper()
	}

	if err := godotenv.Load(); err != nil {
		fmt.Println(err.Error())
	}

	Env.Port = os.Getenv("PORT")
	Env.DbUser = os.Getenv("DB_USER")
	Env.DbPass = os.Getenv("DB_PASS")
	Env.DbName = os.Getenv("DB_NAME")
	Env.DbPort = os.Getenv("DB_PORT")
	Env.JwtSecret = os.Getenv("JWT_SECRET")
	Env.Mail = os.Getenv("MAIL")
	Env.MailPassword = os.Getenv("MAIL_PASSWORD")
	Env.SmtpHost = os.Getenv("SMTP_HOST")
	Env.SmtpPort = os.Getenv("SMTP_PORT")
	Env.AppDomain = os.Getenv("APP_DOMAIN")
}
