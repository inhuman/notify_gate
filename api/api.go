package api

import (
	"fmt"
	"github.com/gobuffalo/packr"
	"github.com/inhuman/notify_gate/config"
	httpErrors "github.com/inhuman/notify_gate/http/errors"
	httpHelpers "github.com/inhuman/notify_gate/http/helpers"
	"github.com/inhuman/notify_gate/notify"
	"github.com/inhuman/notify_gate/pool"
	"github.com/inhuman/notify_gate/senders"
	"github.com/inhuman/notify_gate/service"
	"github.com/pkg/errors"
	"html/template"
	"log"
	"net/http"
)

// Listen is starting listens http api calls
func Listen() {
	http.HandleFunc("/notify", notifyHandler)
	http.HandleFunc("/", mainPage)

	log.Fatal(http.ListenAndServe(":"+config.AppConf.Port, nil))
}

func mainPage(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Main page")
	box := packr.NewBox("./../templates")

	fmt.Println(box)

	html := box.String("index.html")
	tmpl := template.New("main")
	view, err := tmpl.Parse(html)

	if err != nil {
		log.Fatal("Can not expand template", err)
		return
	}

	srcs, err := service.GetAll()
	if err != nil {
		log.Println(err)
	}

	data := struct {
		Services []service.Service
		Title    string
	}{
		Services: srcs,
		Title:    config.AppConf.InstanceTitle,
	}

	view.Execute(w, data)
}

func notifyHandler(w http.ResponseWriter, r *http.Request) {
	n := &notify.Notify{}

	respErr := httpHelpers.ParseRequest(r, n)
	if httpErrors.CheckErrorHTTP(respErr, w, 500) {
		return
	}

	var err error

	switch senders.CheckSenderTypeAvailable(n) {
	case senders.ProviderAvailable:
		err = pool.AddToSave(n)
		if httpErrors.CheckErrorHTTP(err, w, 500) {
			return
		}

	case senders.ProviderUnavailable:
		err = errors.New("Provider " + n.Type + " not available.")
		if httpErrors.CheckErrorHTTP(err, w, 406) {
			return
		}

	case senders.ProvideNotExist:
		err = errors.New("Provider " + n.Type + " not exist.")
		if httpErrors.CheckErrorHTTP(err, w, 404) {
			return
		}
	}
}
