package cache

import (
	"jgit.me/tools/notify_gate/utils"
	"jgit.me/tools/notify_gate/user"
	"fmt"
)

type Tokens map[string]string

var TokensCache = make(Tokens)

func BuildTokenCache() {
	utils.ShowDebugMessage("Building tokens cache")

	users, err := user.GetAll()
	if err!= nil {
		fmt.Println("Build tokens cache error: " + err.Error())
	}

	for _, usr := range users {
		TokensCache[usr.Token] = usr.Name
	}
}