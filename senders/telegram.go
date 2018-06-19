package senders

import (
	"github.com/appscode/go-notify/telegram"
	"jgit.me/tools/notify_gate/config"
	"fmt"
)

func SendToTelegramChat(n *Notify) error {

	fmt.Println("Telegram sender")

	opts := telegram.Options{
		Token: config.AppConf.Telegram.BotToken,
		Channel: n.UIDs,
	}

	cl := telegram.New(opts)

	err := cl.WithBody(n.Message).Send()
	if err != nil {
		return err
	}

	return nil
}


