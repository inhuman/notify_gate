package senders

import (
	"jgit.me/tools/notify_gate/utils"
)

type Notify struct {
	Type    string   `json:"type"`
	Message string   `json:"message"`
	UIDs    []string `json:"uids"`
	Token   string
}

func Send(n *Notify) error {

	utils.ShowDebugMessage("Sender")

	switch n.Type {
	case "TelegramChannel":
		err := SendToTelegramChat(n)
		if err != nil {
			return err
		}

	case "SlackChannel":
		err := SendToSlackChat(n)
		if err != nil {
			return err
		}
	}

	return nil
}
