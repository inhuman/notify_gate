package config

import (
	"github.com/joho/godotenv"
	"os"
	"fmt"
	"errors"
)

var AppConf = &appConfig{}

type appConfig struct {
	Mysql    *MysqlConf
	Telegram *TelegramConf
	Port     string
	Debug    bool
}

type TelegramConf struct {
	BotToken string
}

type MysqlConf struct {
	Host     string
	Port     string
	User     string
	Password string
	DbName   string
}

func init() {
	AppConf.Load()
}

func (c *appConfig) Load(fileNames ...string) error {

	err := godotenv.Overload(fileNames...)
	if err != nil {
		return nil
	}

	if e, ok := os.LookupEnv("ALERTER_DEBUG"); ok {
		fmt.Printf("Alerter debug mode: %s\n", e)
		c.Debug = e == "true"
	} else {
		fmt.Println("Alerter debug mode (default): false")
		c.Debug = false
	}

	if e, ok := os.LookupEnv("NOTIFY_GATE_PORT"); ok {
		fmt.Printf("Notify gate port: %s\n", e)
		c.Port = ":" + e
	} else {
		fmt.Println("Notify gate port (default): 80")
		c.Port = "80"
	}


	c.Mysql, err = loadMysqlConfig()
	if err != nil {
		return err
	}

	c.Telegram, err = loadTelegramConfig()

	return nil
}

func loadMysqlConfig() (*MysqlConf, error) {
	mysql := &MysqlConf{}
	if e, ok := os.LookupEnv("MYSQL_HOST"); ok {
		fmt.Printf("Setup mysql host: %s\n", e)
		mysql.Host = e
	} else {
		return nil, errors.New("MYSQL_HOST is invalid")
	}

	if e, ok := os.LookupEnv("MYSQL_PORT"); ok {
		fmt.Printf("Setup mysql port: %s\n", e)
		mysql.Port = e
	} else {
		fmt.Println("Setup default mysql port: 3306")
		mysql.Port = "3306"
	}

	if e, ok := os.LookupEnv("MYSQL_DB_NAME"); ok {
		fmt.Printf("Setup mysql db: %s\n", e)
		mysql.DbName = e
	}

	if e, ok := os.LookupEnv("MYSQL_USER"); ok {
		fmt.Printf("Setup mysql user: %s\n", e)
		mysql.User = e
	} else {
		return nil, errors.New("MYSQL_USER is invalid")
	}

	if e, ok := os.LookupEnv("MYSQL_PASSWORD"); ok {
		fmt.Println("Setup mysql password")
		mysql.Password = e
	} else {
		return nil, errors.New("MYSQL_PASSWORD is invalid")
	}

	return mysql, nil
}

func loadTelegramConfig() (*TelegramConf, error) {
	telegramConf := &TelegramConf{}

	if e, ok := os.LookupEnv("TELEGRAM_BOT_TOKEN"); ok {
		fmt.Printf("Setup telegram bot token host: %s\n", e)
		telegramConf.BotToken = e
	} else {
		return nil, errors.New("TELEGRAM_BOT_TOKEN is invalid")
	}

	return telegramConf, nil
}
