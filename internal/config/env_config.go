package config

import (
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	Port   string
	DbUser string
	DbPass string
	DbName string
	DbPort string
}

var Env *Config

func LoadEnv() {
	Env = &Config{}
	isCreated, err := checkOrCreateEnv()
	if err != nil {
		panic(err)
	}

	if isCreated {
		fmt.Println("env config created")
		fmt.Printf("HINT:" +
			"PORT: Application port (for example, 8080)\n" +
			"DB_USER: Database user (for example, 'root')\n" +
			"DB_PASS: Database password (for example, '1234')\n" +
			"DB_NAME: Database schema name (for example, 'johny')\n" +
			"DB_PORT: Database port (for mysql, the default is 3306)\n")
		os.Exit(0)
	}

	if err := godotenv.Load(); err != nil {
		fmt.Println(err.Error())
	}

	Env.Port = os.Getenv("PORT")
	Env.DbUser = os.Getenv("DB_USER")
	Env.DbPass = os.Getenv("DB_PASS")
	Env.DbName = os.Getenv("DB_NAME")
	Env.DbPort = os.Getenv("DB_PORT")
}

func checkOrCreateEnv() (bool, error) {
	fileExist, err := os.Stat(".env")
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return true, createEnvTemplate()
		}
		return false, err
	}

	if fileExist.IsDir() {
		return true, createEnvTemplate()
	}
	return false, nil
}

func createEnvTemplate() error {
	file, err := os.Create(".env")
	if err != nil {
		return err
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println(err)
			return
		}
	}(file)

	envText := `PORT=
DB_USER=
DB_PASS=
DB_NAME=
DB_PORT=`

	_, err = file.Write([]byte(envText))
	if err != nil {
		return err
	}
	return nil
}
