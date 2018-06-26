package senders

import (
	go_notify "github.com/appscode/go-notify"
	"github.com/appscode/go-notify/telegram"
	"jgit.me/tools/notify_gate/config"
	"jgit.me/tools/notify_gate/notify"
	"jgit.me/tools/notify_gate/utils"
)

var TelegramClient go_notify.ByChat

func InitTelegramClient() {
	TelegramClient = telegram.New(telegram.Options{
		Token: config.AppConf.Telegram.BotToken,
	})
}

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
