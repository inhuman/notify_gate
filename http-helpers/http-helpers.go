package http_helpers

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"jgit.me/tools/notify_gate/cache"
	"jgit.me/tools/notify_gate/http-errors"
	"errors"
	"jgit.me/tools/notify_gate/config"
)

func ParseRequest(r *http.Request, object interface{}) error {
	jsonBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(jsonBytes, &object)

	if err != nil {
		return err
	}
	return nil
}

func JsonResponse(w http.ResponseWriter, object interface{}) {
	jsn, err := json.Marshal(object)
	http_errors.CheckError(err)
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(200)
	w.Write([]byte(jsn))
}

func Secured(handler func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("X-AUTH-TOKEN")

		//TODO: rebuild cache if unavailable ?
		//if len(cache.tokensCached) == 0 {
		//	cache.BuildServiceTokenCache()
		//}

		if config.AppConf.Debug && (token == "test_token") {
			handler(w, r)
			return
		}

		if len(token) == 0 {
			er := errors.New("Unauthorized")
			http_errors.CheckErrorHttp(er, w, 401)
			return
		}

		_, ok := cache.GetServiceTokens()[token]

		if ok {
			handler(w, r)
		} else {
			er := errors.New("Unauthorized")
			http_errors.CheckErrorHttp(er, w, 401)
			return
		}
	}
}
