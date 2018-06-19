package senders

import (
	"github.com/appscode/go-notify/slack"
	"jgit.me/tools/notify_gate/config"
	"jgit.me/tools/notify_gate/utils"
)

func SendToSlackChat(n *Notify) error {

	utils.ShowDebugMessage("Slack sender")

	opts := slack.Options{
		AuthToken: config.AppConf.SlackConf.AuthToken,
		Channel:   n.UIDs,
	}

	cl := slack.New(opts)

	err := cl.WithBody(n.Message).Send()
	if err != nil {
		return err
	}

	return nil
}
