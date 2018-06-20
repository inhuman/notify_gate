package senders

import (
	"jgit.me/tools/notify_gate/utils"
	"jgit.me/tools/notify_gate/notify"
	"errors"
)

var Providers = make(map[string]func(n *notify.Notify) error)

func init() {
	utils.ShowDebugMessage("Init send providers")
	Providers["TelegramChannel"] = SendToTelegramChat
	Providers["SlackChannel"] = SendToSlackChat
}

// The function call send provider if it exists on provider map,
// or return error if it doesn't.
// Also error returned if provider can not send the send
func Send(n *notify.Notify) error {
	if provider, ok := Providers[n.Type]; ok {
		err := provider(n)
		if err != nil {
			return err
		}
	} else {
		return errors.New("no provider for type " + n.Type)
	}

	return nil
}
