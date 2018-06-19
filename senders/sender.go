package senders

import "fmt"

type Notify struct {
	Type    string   `json:"type"`
	Message string   `json:"message"`
	UIDs     []string `json:"uids"`
}

func Send(n *Notify) error {

	fmt.Println("Sender")

	switch n.Type {
	case "TelegramChannel":
		err := SendToTelegramChat(n)
		if err != nil {
			return err
		}

	case "SlackChannel":
		err := SendToTelegramChat(n)
		if err != nil {
			return err
		}
	}

	return nil
}
