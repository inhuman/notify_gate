package utils

import (
	"github.com/inhuman/notify_gate/config"
	"log"
)

// ShowDebugMessage is used for printing debug messages
func ShowDebugMessage(i interface{}) {
	if config.AppConf.Debug {
		log.Println(i)
	}
}

// CheckError is used for log error if not nil
func CheckError(err error) {
	if err != nil {
		log.Println("error:", err)
	}
}

func CheckErrorMessage(message string, err error) {
	if err != nil {
		log.Println(message, err)
	}
}
