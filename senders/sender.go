package senders

import (
	"errors"
	"github.com/inhuman/notify_gate/config"
	"github.com/inhuman/notify_gate/notify"
	"log"
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

	log.Println("Init send providers")

	initTelegramSender()
	initSlackSender()

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

// initTelegramSender is used for initialize telegram sender
func initTelegramSender() {
	if config.AppConf.Senders.Telegram != nil {
		initTelegramClient()
		providers["TelegramChannel"] = sendToTelegramChat
		log.Println("Telegram sender initialized")
	} else {
		providers["TelegramChannel"] = nil
	}
}

// initSlackSender is used for initialize slack sender
func initSlackSender() {
	if config.AppConf.Senders.Slack != nil {
		providers["SlackChannel"] = sendToSlackChat
		log.Println("Slack sender initialized")
	} else {
		providers["SlackChannel"] = nil
	}
}

func AddSender(name string, f func(n *notify.Notify) error) error {
	providers[name] = f

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
