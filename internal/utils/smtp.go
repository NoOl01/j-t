package utils

import (
	"fmt"
	"johny-tuna/internal/config"
	"net/smtp"
)

func SendMessage(email, token string) error {
	from := config.Env.Mail
	password := config.Env.MailPassword
	smtpHost := config.Env.SmtpHost
	smtpPort := config.Env.SmtpPort

	auth := smtp.PlainAuth("", from, password, smtpHost)

	var domain string
	if config.Env.AppDomain == "localhost" {
		domain = fmt.Sprintf("http://localhost:%s", config.Env.Port)
	} else {
		domain = fmt.Sprintf("https://%s", config.Env.AppDomain)
	}
	link := fmt.Sprintf("%s/verify/token/%s", domain, token)

	message := []byte(fmt.Sprintf(
		"From: %s\r\n"+
			"To: %s\r\n"+
			"Subject: Подтверждение почты\r\n"+
			"MIME-Version: 1.0\r\n"+
			"Content-Type: text/plain; charset=\"UTF-8\"\r\n"+
			"\r\n"+
			"Для подтверждения почты и регистрации аккаунта перейдите по ссылке ниже:\r\n%s",
		from, email, link))

	return smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{email}, message)
}
