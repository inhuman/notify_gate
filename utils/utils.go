package utils

import (
	"jgit.me/tools/alerter/config"
	"fmt"
	"strings"
)

func ShowDebugMessage(i interface{}) {
	if config.AppConf.Debug {
		fmt.Println(i)
	}
}

func MaskString(s string, showLastSymbols int) string {
	if len(s) <= showLastSymbols {
		return s
	}
	return strings.Repeat("*", len(s)-showLastSymbols) + s[len(s)-showLastSymbols:]
}
