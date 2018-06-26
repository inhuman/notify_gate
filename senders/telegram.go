package senders

import (
	go_notify "github.com/appscode/go-notify"
	"github.com/appscode/go-notify/telegram"
	"jgit.me/tools/notify_gate/config"
	"jgit.me/tools/notify_gate/notify"
	"jgit.me/tools/notify_gate/utils"
)

var telegramClient go_notify.ByChat

func initTelegramClient() {
	telegramClient = telegram.New(telegram.Options{
		Token: config.AppConf.Senders.Telegram.BotToken,
	})
}

func sendToTelegramChat(n *notify.Notify) error {

	utils.ShowDebugMessage("Telegram sender")

	err := telegramClient.
		To("", n.UIDs...).
		WithBody(n.Message).
		Send()

	if err != nil {
		return err
	}

	return nil
}
