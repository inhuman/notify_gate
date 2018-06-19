package utils

import (
	"jgit.me/tools/alerter/config"
	"fmt"
)

func ShowDebugMessage(i interface{}) {

	if config.AppConf.Debug {
		fmt.Println(i)
	}

}