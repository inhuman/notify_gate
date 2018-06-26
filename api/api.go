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
	"jgit.me/tools/notify_gate/senders"
	"github.com/pkg/errors"
	"github.com/gobuffalo/packr"
	"fmt"
)


// Listen is starting listens http api calls
func Listen() {
	http.HandleFunc("/notify", http_helpers.Secured(notifyHandler))
	http.HandleFunc("/service/register", registerService)
	http.HandleFunc("/service/unregister", http_helpers.Secured(unregisterService))
	http.HandleFunc("/", mainPage)

	log.Fatal(http.ListenAndServe(config.AppConf.Port, nil))
}

func mainPage(w http.ResponseWriter, r *http.Request) {

	box := packr.NewBox("./../templates")
	html := box.String("index.html")
	tmpl := template.New("main")
	view, err := tmpl.Parse(html)

	if err != nil {
		log.Fatal("Can not expand template", err)
		return
	}

	srcs, err := service.GetAll()
	if err != nil {
		fmt.Println(err)
	}

	data := struct {
		Services []service.Service
		Title string
	}{
		Services: srcs,
		Title: config.AppConf.InstanceTitle,
	}

	view.Execute(w, data)
}

func notifyHandler(w http.ResponseWriter, r *http.Request) {
	n := &notify.Notify{}

	respErr := http_helpers.ParseRequest(r, n)
	if http_errors.CheckErrorHttp(respErr, w, 500) {
		return
	}

	var err error

	switch senders.CheckSenderTypeAvailable(n) {
	case senders.ProviderAvailable:
		err = pool.NPool.AddToSave(n)
		if http_errors.CheckErrorHttp(err, w, 500) {
			return
		}

	case senders.ProviderUnavailable:
		err = errors.New("Provider " + n.Type + " not available.")
		if http_errors.CheckErrorHttp(err, w, 406) {
			return
		}

	case senders.ProvideNotExist:
		err = errors.New("Provider " + n.Type + " not exist.")
		if http_errors.CheckErrorHttp(err, w, 404) {
			return
		}
	}
}

func registerService(w http.ResponseWriter, r *http.Request) {
	u := &service.Service{}
	err := http_helpers.ParseRequest(r, u)
	if http_errors.CheckErrorHttp(err, w, http.StatusBadRequest) {
		return
	}
	srvs, err := service.Register(u)

	if !http_errors.CheckErrorHttp(err, w, 409) {
		cache.AddServiceToken(srvs.Name, srvs.Token)
		http_helpers.JsonResponse(w, srvs)
	}
}

func unregisterService(w http.ResponseWriter, r *http.Request) {
	u := &service.Service{
		Token: r.Header.Get("X-AUTH-TOKEN"),
	}

	err := service.Unregister(u)
	http_errors.CheckErrorHttp(err, w, 404)
}
