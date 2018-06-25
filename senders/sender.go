package senders

import (
	"jgit.me/tools/notify_gate/utils"
	"jgit.me/tools/notify_gate/notify"
	"errors"
	"jgit.me/tools/notify_gate/config"
)

const(
	ProviderAvailable = 1
	ProviderUnavailable = 2
	ProvideNotExist = 3
)

var Providers = make(map[string]func(n *notify.Notify) error)

func Init() error {

	utils.ShowDebugMessage("Init send providers")

	if config.AppConf.Telegram != nil {
		InitTelegramClient()
		Providers["TelegramChannel"] = SendToTelegramChat
	} else {
		Providers["TelegramChannel"] = nil
	}

	if config.AppConf.SlackConf != nil {
		Providers["SlackChannel"] = SendToSlackChat
	} else {
		Providers["SlackChannel"] = nil
	}

	atLeastOneProviderAvailable := false

	for _, prov := range Providers {
		if prov != nil {
			atLeastOneProviderAvailable = true
		}
	}

	if !atLeastOneProviderAvailable {
		return errors.New("No send providers available. Exiting..")
	}

	return nil
}

// The function call send provider if it exists on provider map,
// or return error if it doesn't.
// Also error returned if provider can not send the notify
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

func CheckSenderTypeAvailable(n *notify.Notify) int {
	if provider, ok := Providers[n.Type]; ok {
		if provider != nil {
			return ProviderAvailable
		} else {
			return ProviderUnavailable
		}
	}
	return ProvideNotExist
}