package senders

import (
	go_notify "github.com/appscode/go-notify"
	"github.com/appscode/go-notify/slack"
	"jgit.me/tools/notify_gate/config"
	"jgit.me/tools/notify_gate/notify"
	"jgit.me/tools/notify_gate/utils"
)

var slackClient go_notify.ByChat

func initSlackClient() {
	slackClient = slack.New(slack.Options{
		AuthToken: config.AppConf.Senders.Slack.AuthToken,
	})
}

func sendToSlackChat(n *notify.Notify) error {

	utils.ShowDebugMessage("Slack sender")

	err := slackClient.
		To("", n.UIDs...).
		WithBody(n.Message).
		Send()

	if err != nil {
		return err
	}

	return nil
}
