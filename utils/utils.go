package utils

import (
	"fmt"
	"jgit.me/tools/notify_gate/config"
)


// ShowDebugMessage is used for printing debug messages
func ShowDebugMessage(i interface{}) {
	if config.AppConf.Debug {
		fmt.Println(i)
	}
}
