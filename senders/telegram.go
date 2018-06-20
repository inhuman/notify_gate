package senders

import (
	"github.com/appscode/go-notify/telegram"
	"jgit.me/tools/notify_gate/config"
	"fmt"
	"jgit.me/tools/notify_gate/notify"
)

func SendToTelegramChat(n *notify.Notify) error {

	fmt.Println("Telegram sender")

	opts := telegram.Options{
		Token:   config.AppConf.Telegram.BotToken,
		Channel: n.UIDs,
	}

	cl := telegram.New(opts)

	err := cl.WithBody(n.Message).Send()
	if err != nil {
		return err
	}

	return nil
}
