package utils

import (
	"fmt"
	"jgit.me/tools/notify_gate/config"
)

func ShowDebugMessage(i interface{}) {
	if config.AppConf.Debug {
		fmt.Println(i)
	}
}
