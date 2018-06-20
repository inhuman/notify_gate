package senders

import (
	"github.com/appscode/go-notify/telegram"
	"jgit.me/tools/notify_gate/config"
	"jgit.me/tools/notify_gate/notify"
	"jgit.me/tools/notify_gate/utils"
)

var TelegramClient = telegram.New(telegram.Options{Token: config.AppConf.Telegram.BotToken})

func SendToTelegramChat(n *notify.Notify) error {

	utils.ShowDebugMessage("Telegram sender")

	err := TelegramClient.
		To("", n.UIDs...).
		WithBody(n.Message).
		Send()

	if err != nil {
		return err
	}

	return nil
}
