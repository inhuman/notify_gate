package senders

import (
	"errors"
	"fmt"
	"jgit.me/tools/notify_gate/config"
	"jgit.me/tools/notify_gate/notify"
)

// Send provider statuses
const (
	ProviderAvailable   = 1
	ProviderUnavailable = 2
	ProvideNotExist     = 3
)

var providers = make(map[string]func(n *notify.Notify) error)


// Init is used for initialize and build map of senders
func Init() error {

	fmt.Println("Init send providers")

	//TODO: refactor to cycle that iterates throw senders
	if config.AppConf.Senders.Telegram != nil {
		initTelegramClient()
		providers["TelegramChannel"] = sendToTelegramChat
		fmt.Println("Telegram sender initialized")
	} else {
		providers["TelegramChannel"] = nil
	}

	if config.AppConf.Senders.Slack != nil {
		initSlackClient()
		providers["SlackChannel"] = sendToSlackChat
		fmt.Println("Slack sender initialized")
	} else {
		providers["SlackChannel"] = nil
	}

	atLeastOneProviderAvailable := false

	for _, prov := range providers {
		if prov != nil {
			atLeastOneProviderAvailable = true
		}
	}

	if !atLeastOneProviderAvailable {
		return errors.New("no send providers available, exiting")
	}

	return nil
}

// Send is used for call send method of provider if it exists on provider map and equal Notify.Type,
// or return error if it doesn't.
// Also error returned if provider can not send the notify
func Send(n *notify.Notify) error {
	if prov, ok := providers[n.Type]; ok {
		err := prov(n)
		if err != nil {
			return err
		}
	} else {
		return errors.New("no provider for type " + n.Type)
	}

	return nil
}

// CheckSenderTypeAvailable is used for check that provider exist and available for given Notify,
// and return provider status
func CheckSenderTypeAvailable(n *notify.Notify) int {
	if prov, ok := providers[n.Type]; ok {
		if prov != nil {
			return ProviderAvailable
		}
		return ProviderUnavailable
	}
	return ProvideNotExist
}
