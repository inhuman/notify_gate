package senders

import (
	"github.com/inhuman/go-notify/slack"
	"github.com/inhuman/notify_gate/config"
	"github.com/inhuman/notify_gate/notify"
	"github.com/inhuman/notify_gate/utils"
)

func sendToSlackChat(n *notify.Notify) error {

	utils.ShowDebugMessage("Slack sender")

	opts := slack.Options{
		AuthToken: config.AppConf.Senders.Slack.AuthToken,
		Channel:   n.UIDs,
	}

	cl := slack.New(opts)

	err := cl.WithBody(n.Message).Send()
	if err != nil {
		return err
	}

	return nil
}
