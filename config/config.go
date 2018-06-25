package config

import (
	"github.com/joho/godotenv"
	"os"
	"fmt"
	"errors"
	"strings"
)

var AppConf = &appConfig{}

type appConfig struct {
	Postgres   *PostgreConf
	Telegram  *TelegramConf
	SlackConf *SlackConf
	Port      string
	Debug     bool
}

type TelegramConf struct {
	BotToken string
}

type SlackConf struct {
	AuthToken string
}

type PostgreConf struct {
	Host     string
	Port     string
	User     string
	Password string
	DbName   string
}


func (c *appConfig) Load(fileNames ...string) error {

	err := godotenv.Overload(fileNames...)
	if err != nil {
		fmt.Println(".env file not found, trying fetch environment variables")
	}

	if e, ok := os.LookupEnv("NG_DEBUG"); ok {
		fmt.Printf("Notify gate debug mode: %s\n", e)
		c.Debug = e == "true"
	} else {
		fmt.Println("Notify gate debug mode (default): false")
		c.Debug = false
	}

	if e, ok := os.LookupEnv("NG_PORT"); ok {
		fmt.Printf("Notify gate port: %s\n", e)
		c.Port = ":" + e
	} else {
		fmt.Println("Notify gate port (default): 80")
		c.Port = ":80"
	}

	c.Postgres, err = loadPostgreConfig()
	if err != nil {
		return err
	}

	c.Telegram = loadTelegramConfig()

	c.SlackConf = loadSlackConfig()


	return nil
}

func loadPostgreConfig() (*PostgreConf, error) {
	Postgre := &PostgreConf{}
	if e, ok := os.LookupEnv("POSTGRES_HOST"); ok {
		fmt.Printf("Setup Postgre host: %s\n", e)
		Postgre.Host = e
	} else {
		return nil, errors.New("POSTGRES_HOST is invalid")
	}

	if e, ok := os.LookupEnv("POSTGRES_PORT"); ok {
		fmt.Printf("Setup Postgre port: %s\n", e)
		Postgre.Port = e
	} else {
		fmt.Println("Setup default Postgre port: 3306")
		Postgre.Port = "3306"
	}

	if e, ok := os.LookupEnv("POSTGRES_DB_NAME"); ok {
		fmt.Printf("Setup Postgre db: %s\n", e)
		Postgre.DbName = e
	}

	if e, ok := os.LookupEnv("POSTGRES_USER"); ok {
		fmt.Printf("Setup Postgre user: %s\n", e)
		Postgre.User = e
	} else {
		return nil, errors.New("POSTGRES_USER is invalid")
	}

	if e, ok := os.LookupEnv("POSTGRES_PASSWORD"); ok {
		fmt.Printf("Setup Postgre password: %s\n", maskString(e, 0))
		Postgre.Password = e
	} else {
		return nil, errors.New("POSTGRES_PASSWORD is invalid")
	}

	return Postgre, nil
}

func loadTelegramConfig() *TelegramConf {
	telegramConf := &TelegramConf{}

	if e, ok := os.LookupEnv("TELEGRAM_BOT_TOKEN"); ok {
		fmt.Printf("Setup telegram bot token: %s\n", maskString(e, 6))
		telegramConf.BotToken = e
	} else {
		fmt.Println("TELEGRAM_BOT_TOKEN not found, telegram notification disabling...")
		return nil
	}

	return telegramConf
}

func loadSlackConfig() *SlackConf {
	slackConf := &SlackConf{}

	if e, ok := os.LookupEnv("SLACK_AUTH_TOKEN"); ok {
		fmt.Printf("Setup slack auth token: %s\n", maskString(e, 6))
		slackConf.AuthToken = e
	} else {
		fmt.Println("SLACK_AUTH_TOKEN not found, slack notification disabling...")
		return nil
	}
	return slackConf
}

func maskString(s string, showLastSymbols int) string {
	if len(s) <= showLastSymbols {
		return s
	}
	return strings.Repeat("*", len(s)-showLastSymbols) + s[len(s)-showLastSymbols:]
}
