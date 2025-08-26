package config

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func CheckEnv() bool {
	file, err := os.Stat(".env")
	if err != nil || file.IsDir() {
		return false
	}

	return true
}

func EnvHelper() {
	var config Config
	fmt.Println("Let's set up the env file.")
	time.Sleep(1 * time.Second)
	fmt.Println("If there is a default value (example: (default: 8080)) in subsequent messages, pressing 'enter' (sending an empty string) will send the default value.")
	time.Sleep(1 * time.Second)

	helperPrint("Enter the port on which the server will start (default: 8080)")
	readWithDefault(&config.Port, "8080")

	helperPrint("Enter your database user (mysql)")
	readWithoutDefault(&config.DbUser)

	helperPrint("Enter your database password (mysql)")
	readWithoutDefault(&config.DbPass)

	helperPrint("Enter your database name (mysql)")
	readWithoutDefault(&config.DbName)

	helperPrint("Enter your database port (default: 3306)")
	readWithDefault(&config.DbPort, "3306")

	helperPrint("Enter your jwt secret key (random chars)")
	readWithoutDefault(&config.JwtSecret)

	helperPrint("Enter your email sender (example: \"yourmail@mail.ru\"")
	readWithoutDefault(&config.Mail)

	helperPrint("Enter your email password (use the app-specific password, NOT your regular email password)")
	readWithoutDefault(&config.MailPassword)

	helperPrint("Enter smtp host (default: smtp.mail.ru)")
	readWithDefault(&config.SmtpHost, "smtp.mail.ru")

	helperPrint("Enter smtp port (default: 587)")
	readWithDefault(&config.SmtpPort, "587")

	helperPrint("Enter your app domain (default: localhost)")
	readWithDefault(&config.AppDomain, "localhost")

	generateEnv(&config)
}

func helperPrint(text string) {
	fmt.Println(text)
	fmt.Print("> ")
}

func readWithDefault(value *string, defaultVal string) {
	reader := bufio.NewReader(os.Stdin)
	tempValue, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err.Error())
	}

	tempValue = strings.TrimSpace(tempValue)
	if tempValue == "" {
		tempValue = defaultVal
	}

	*value = tempValue
}

func readWithoutDefault(value *string) {
	reader := bufio.NewReader(os.Stdin)

	for {
		tempValue, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err.Error())
		}

		tempValue = strings.TrimSpace(tempValue)
		if tempValue == "" {
			fmt.Println("This value cannot be empty, please try again:")
			continue
		}

		*value = tempValue
		break
	}
}

func generateEnv(config *Config) {
	file, err := os.Create(".env")
	if err != nil {
		log.Fatal(err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Println(err)
			return
		}
	}(file)

	content := fmt.Sprintf(`PORT=%s
DB_USER=%s
DB_PASS=%s
DB_NAME=%s
DB_PORT=%s
JWT_SECRET=%s
MAIL=%s
MAIL_PASSWORD=%s
SMTP_HOST=%s
SMTP_PORT=%s
APP_DOMAIN=%s
`, config.Port, config.DbUser, config.DbPass, config.DbName,
		config.DbPort, config.JwtSecret, config.Mail, config.MailPassword,
		config.SmtpHost, config.SmtpPort, config.AppDomain)

	if _, err = file.Write([]byte(content)); err != nil {
		log.Fatal(err)
	}
}
