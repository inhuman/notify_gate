package helpers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"jgit.me/tools/notify_gate/cache"
	"jgit.me/tools/notify_gate/config"
	httpErrors "jgit.me/tools/notify_gate/http/errors"
	"net/http"
)

// ParseRequest is used for parsing http.Request into given interface
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

// JSONResponse is used for marshaling interface to json and write it to http.ResponseWriter
func JSONResponse(w http.ResponseWriter, object interface{}) {
	jsn, err := json.Marshal(object)
	httpErrors.CheckError(err)
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(200)
	w.Write([]byte(jsn))
}

// Secured is used for checking auth token and refuse request if token invalid
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
			httpErrors.CheckErrorHTTP(er, w, 401)
			return
		}

		_, ok := cache.GetServiceTokens()[token]

		if ok {
			handler(w, r)
		} else {
			er := errors.New("Unauthorized")
			httpErrors.CheckErrorHTTP(er, w, 401)
			return
		}
	}
}
