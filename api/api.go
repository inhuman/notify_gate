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
	"jgit.me/tools/notify_gate/user"
	"jgit.me/tools/notify_gate/db"
	"jgit.me/tools/notify_gate/cache"
	"github.com/pkg/errors"
	"html/template"
)

func Listen() {

	db.Init()
	usr := user.Service{}
	db.Stor.Migrate(usr)

	cache.BuildTokenCache()

	http.HandleFunc("/notify", Secured(notifyHandler))
	http.HandleFunc("/register", register)
	http.HandleFunc("/", mainPage)

	http.HandleFunc("/get_registered", getAll)
	log.Fatal(http.ListenAndServe(config.AppConf.Port, nil))
}

func mainPage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.New("main")
	tmpl, err := template.New("index.html").ParseFiles("index.html")
	if err != nil {
		log.Fatal("Can not expand template", err)
		return
	}

	tmpl.Execute(w, nil)
}

func notifyHandler(w http.ResponseWriter, r *http.Request) {
	n := &senders.Notify{}
	err := ParseRequest(r, n)
	http_errors.CheckErrorHttp(err, w, 500)

	fmt.Printf("%+v\n", n)

	err = senders.Send(n)
	http_errors.CheckErrorHttp(err, w, 500)
}

func register(w http.ResponseWriter, r *http.Request) {
	u := &user.Service{}
	err := ParseRequest(r, u)
	if http_errors.CheckErrorHttp(err, w, http.StatusBadRequest) {
		return
	}

	res, err := user.Register(u)
	if !http_errors.CheckErrorHttp(err, w, 409) {
		JsonResponse(w, res)
	}
}

func getAll(w http.ResponseWriter, r *http.Request) {

	users, err := user.GetAll()
	http_errors.CheckErrorHttp(err, w, 500)
	JsonResponse(w, users)
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

		//if len(cache.TokensCache) == 0 {
		//	cache.BuildTokenCache()
		//}

		if len(token) == 0 {
			er := errors.New("Unauthorized")
			http_errors.CheckErrorHttp(er, w, 401)
			return
		}

		_, ok := cache.TokensCache[token]

		if ok {
			handler(w, r)
		} else {
			er := errors.New("Unauthorized")
			http_errors.CheckErrorHttp(er, w, 401)
			return
		}
	}
}
