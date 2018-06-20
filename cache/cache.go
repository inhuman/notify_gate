package cache

import (
	"jgit.me/tools/notify_gate/utils"
	"jgit.me/tools/notify_gate/service"
	"fmt"
)

type Tokens map[string]string

var TokensCache = make(Tokens)

func BuildTokenCache() {
	utils.ShowDebugMessage("Building tokens cache")

	users, err := service.GetAll()
	if err!= nil {
		fmt.Println("Build tokens cache error: " + err.Error())
	}

	for _, usr := range users {
		TokensCache[usr.Token] = usr.Name
	}
}