package cache

import (
	"fmt"
	"jgit.me/tools/notify_gate/service"
	"jgit.me/tools/notify_gate/utils"
)

var tokensCached = make(map[string]string)

// BuildServiceTokenCache is used for build service tokens cache from db.
// Tokens cache used for authorize services
func BuildServiceTokenCache() {
	utils.ShowDebugMessage("Building tokens cache")

	srvs, err := service.GetAll()
	if err != nil {
		fmt.Println("Build service tokens cache error: " + err.Error())
	}

	for _, usr := range srvs {
		tokensCached[usr.Token] = usr.Name
	}
}

// GetServiceTokens is used for receive service tokens cache
func GetServiceTokens() map[string]string {
	return tokensCached
}

// AddServiceToken is used to add service token to cache
func AddServiceToken(serviceName, token string) {
	tokensCached[token] = serviceName
}

// InvalidateServiceTokens is used to invalidate service tokens cache
func InvalidateServiceTokens() {
	tokensCached = make(map[string]string)
}
