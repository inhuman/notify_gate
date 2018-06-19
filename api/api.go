package api

import (
	"net/http"
	"log"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"jgit.me/tools/notify_gate/http-errors"
	"jgit.me/tools/notify_gate/config"
	"jgit.me/tools/notify_gate/senders"
)

func Listen() {
	http.HandleFunc("/notify", notifyHandler)
	log.Fatal(http.ListenAndServe(config.AppConf.Port, nil))
}

func notifyHandler(w http.ResponseWriter, r *http.Request) {
	n := &senders.Notify{}
	err := ParseRequest(r, n)
	http_errors.CheckErrorHttp(err, w, 500)

	fmt.Printf("%+v\n", n)

	err = senders.Send(n)
	http_errors.CheckErrorHttp(err, w, 500)
}

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
