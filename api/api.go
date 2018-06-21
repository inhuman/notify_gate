package api

import (
	"net/http"
	"log"
	"jgit.me/tools/notify_gate/http-errors"
	"jgit.me/tools/notify_gate/config"
	"jgit.me/tools/notify_gate/service"
	"jgit.me/tools/notify_gate/cache"
	"html/template"
	"jgit.me/tools/notify_gate/http-helpers"
	"jgit.me/tools/notify_gate/pool"
	"jgit.me/tools/notify_gate/notify"
)

func Listen() {

	http.HandleFunc("/notify", http_helpers.Secured(notifyHandler))
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
	n := &notify.Notify{}

	err := http_helpers.ParseRequest(r, n)
	http_errors.CheckErrorHttp(err, w, 500)

	err = pool.NPool.AddToSave(n)
	http_errors.CheckErrorHttp(err, w, 500)
}

func register(w http.ResponseWriter, r *http.Request) {
	u := &service.Service{}
	err := http_helpers.ParseRequest(r, u)
	if http_errors.CheckErrorHttp(err, w, http.StatusBadRequest) {
		return
	}

	res, err := service.Register(u)

	cache.TokensCache[res.Token] = res.Name

	if !http_errors.CheckErrorHttp(err, w, 409) {
		http_helpers.JsonResponse(w, res)
	}
}

func getAll(w http.ResponseWriter, r *http.Request) {

	users, err := service.GetAll()
	http_errors.CheckErrorHttp(err, w, 500)
	http_helpers.JsonResponse(w, users)
}
