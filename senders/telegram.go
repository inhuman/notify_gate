package senders

import (
	go_notify "github.com/inhuman/go-notify"
	"github.com/inhuman/go-notify/telegram"
	"github.com/inhuman/notify_gate/config"
	"github.com/inhuman/notify_gate/notify"
	"github.com/inhuman/notify_gate/utils"
	"log"
)

var telegramClient go_notify.ByChat

func initTelegramClient() {
	telegramClient = telegram.New(telegram.Options{
		Token: config.AppConf.Senders.Telegram.BotToken,
	})
}

func sendToTelegramChat(n *notify.Notify) error {

	utils.ShowDebugMessage("Telegram sender")
	log.Printf("notify: %+v\n", n)

	err := telegramClient.
		To(n.UIDs[0], n.UIDs[1:]...).
		WithBody(n.Message).
		Send()

	if err != nil {
		return err
	}

	return nil
}
