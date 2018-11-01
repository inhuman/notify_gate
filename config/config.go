package config

import (
	"github.com/inhuman/config_merger"
	"github.com/joho/godotenv"
	"log"
)

// AppConf is main app config
var AppConf = &appConfig{}

type appConfig struct {
	DB            DBConf
	Port          string `env:"NG_UI_PORT"`
	Debug         bool   `env:"NG_DEBUG"`
	InstanceTitle string `env:"NG_INSTANCE_TITLE"`
	Senders       struct {
		Telegram telegramConf
		Slack    slackConf
	}
}

type telegramConf struct {
	BotToken string `env:"TELEGRAM_BOT_TOKEN" show_last_symbols:"6"`
}

type slackConf struct {
	AuthToken string `env:"SLACK_AUTH_TOKEN" show_last_symbols:"6"`
}

type DBConf struct {
	Host     string `env:"DB_HOST"`
	Port     string `env:"DB_PORT"`
	User     string `env:"DB_USER"`
	Password string `env:"DB_PASSWORD"`
	Name     string `env:"DB_NAME"`
	Type     string `env:"DB_TYPE" required:"true"`
}

func (c *appConfig) Load(fileNames ...string) error {

	err := godotenv.Overload(fileNames...)
	if err != nil {
		log.Println(".env file not found, trying fetch environment variables")
	}

	configMerger := config_merger.NewMerger(c)

	configMerger.AddSource(&config_merger.EnvSource{
		Variables: []string{
			"NG_UI_PORT",
			"NG_DEBUG",
			"NG_INSTANCE_TITLE",
			"DB_TYPE",
			"DB_HOST",
			"DB_PORT",
			"DB_USER",
			"DB_PASSWORD",
			"DB_NAME",
			"TELEGRAM_BOT_TOKEN",
			"SLACK_AUTH_TOKEN",
		},
	})

	err = configMerger.Run()
	if err != nil {
		return err
	}

	configMerger.PrintConfig()

	return nil
}
