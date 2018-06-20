package senders

import (
	"github.com/appscode/go-notify/slack"
	"jgit.me/tools/notify_gate/config"
	"jgit.me/tools/notify_gate/utils"
	"jgit.me/tools/notify_gate/notify"
)

func SendToSlackChat(n *notify.Notify) error {

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
