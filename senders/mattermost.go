package senders

import (
	go_notify "github.com/inhuman/go-notify"
	"github.com/inhuman/go-notify/mattermost"
	"github.com/inhuman/notify_gate/config"
	"github.com/inhuman/notify_gate/notify"
	"github.com/inhuman/notify_gate/utils"
	"log"
)

var mattermostClient go_notify.ByChat

func initMattermostClient() {
	mattermostClient = mattermost.New(mattermost.Options{
		Url:    config.AppConf.Senders.Mattermost.Url,
		HookId: config.AppConf.Senders.Mattermost.HookId,
	})
}

func sendToMattermostChat(n *notify.Notify) error {

	utils.ShowDebugMessage("Mattermost sender")
	log.Printf("notify: %+v\n", n)

	err := mattermostClient.
		To("", n.UIDs...).
		WithBody(n.Message).
		Send()

	if err != nil {
		return err
	}

	return nil
}
