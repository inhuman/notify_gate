package utils

import (
	"jgit.me/tools/notify_gate/config"
	"fmt"
)

func ShowDebugMessage(i interface{}) {
	if config.AppConf.Debug {
		fmt.Println(i)
	}
}

