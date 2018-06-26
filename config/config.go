package config

import (
	"errors"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strings"
)

// AppConf is main app config
var AppConf = &appConfig{}

type appConfig struct {
	Postgres      *postgreConf
	Port          string
	Debug         bool
	InstanceTitle string
	Senders       struct {
		Telegram *telegramConf
		Slack    *slackConf
	}
}

type telegramConf struct {
	BotToken string
}

type slackConf struct {
	AuthToken string
}

type postgreConf struct {
	Host     string
	Port     string
	User     string
	Password string
	DbName   string
}

func (c *appConfig) Load(fileNames ...string) error {

	err := godotenv.Overload(fileNames...)
	if err != nil {
		log.Println(".env file not found, trying fetch environment variables")
	}

	if e, ok := os.LookupEnv("NG_DEBUG"); ok {
		log.Printf("Notify gate debug mode: %s\n", e)
		c.Debug = e == "true"
	} else {
		log.Println("Notify gate debug mode (default): false")
		c.Debug = false
	}

	if e, ok := os.LookupEnv("NG_INSTANCE_TITLE"); ok {
		log.Printf("Notify gate instance title: %s\n", e)
		c.InstanceTitle = e
	} else {
		log.Println("Notify gate instance title (default): <none>")
		c.InstanceTitle = ""
	}

	if e, ok := os.LookupEnv("NG_UI_PORT"); ok {
		log.Printf("Notify gate port: %s\n", e)
		c.Port = ":" + e
	} else {
		log.Println("Notify gate port (default): 80")
		c.Port = ":80"
	}

	c.Postgres, err = loadPostgreConfig()
	if err != nil {
		return err
	}

	c.Senders.Telegram = loadTelegramConfig()

	c.Senders.Slack = loadSlackConfig()

	return nil
}

func loadPostgreConfig() (*postgreConf, error) {
	Postgre := &postgreConf{}
	if e, ok := os.LookupEnv("POSTGRES_HOST"); ok {
		log.Printf("Setup Postgre host: %s\n", e)
		Postgre.Host = e
	} else {
		return nil, errors.New("POSTGRES_HOST not found")
	}

	if e, ok := os.LookupEnv("POSTGRES_PORT"); ok {
		log.Printf("Setup Postgre port: %s\n", e)
		Postgre.Port = e
	} else {
		log.Println("Setup default Postgres port: 5432")
		Postgre.Port = "5432"
	}

	if e, ok := os.LookupEnv("POSTGRES_DB"); ok {
		log.Printf("Setup Postgres db: %s\n", e)
		Postgre.DbName = e
	}

	if e, ok := os.LookupEnv("POSTGRES_USER"); ok {
		log.Printf("Setup Postgres user: %s\n", e)
		Postgre.User = e
	} else {
		return nil, errors.New("POSTGRES_USER not found")
	}

	if e, ok := os.LookupEnv("POSTGRES_PASSWORD"); ok {
		log.Printf("Setup Postgre password: %s\n", maskString(e, 0))
		Postgre.Password = e
	} else {
		return nil, errors.New("POSTGRES_PASSWORD not found")
	}

	return Postgre, nil
}

func loadTelegramConfig() *telegramConf {
	telegramConf := &telegramConf{}

	if e, ok := os.LookupEnv("TELEGRAM_BOT_TOKEN"); ok {
		log.Printf("Setup telegram bot token: %s\n", maskString(e, 6))
		telegramConf.BotToken = e
	} else {
		log.Println("TELEGRAM_BOT_TOKEN not found, telegram notification disabling...")
		return nil
	}

	return telegramConf
}

func loadSlackConfig() *slackConf {
	slackConf := &slackConf{}

	if e, ok := os.LookupEnv("SLACK_AUTH_TOKEN"); ok {
		log.Printf("Setup slack auth token: %s\n", maskString(e, 6))
		slackConf.AuthToken = e
	} else {
		log.Println("SLACK_AUTH_TOKEN not found, slack notification disabling...")
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
