package main

import (
	"github.com/inhuman/notify_gate/api"
	"github.com/inhuman/notify_gate/cache"
	"github.com/inhuman/notify_gate/config"
	"github.com/inhuman/notify_gate/db"
	"github.com/inhuman/notify_gate/notify"
	"github.com/inhuman/notify_gate/pool"
	"github.com/inhuman/notify_gate/senders"
	"github.com/inhuman/notify_gate/service"
	"github.com/inhuman/notify_gate/workerpool"
	"os"
	"log"
)

func main() {

	err := runApp()
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	} else {
		os.Exit(0)
	}

}

func runApp() error {

	if len(os.Args) > 1 {
		err := config.AppConf.Load(os.Args...)
		if err != nil {
			return err
		}
	} else {
		err := config.AppConf.Load()
		if err != nil {
			return err
		}
	}

	err := senders.Init()
	if err != nil {
		return err
	}

	db.Init()

	db.Stor.Migrate(service.Service{})
	db.Stor.Migrate(notify.Notify{})


	cache.BuildServiceTokenCache()
	wpool := workerpool.NewPool(5)

	go pool.Saver(wpool)
	go pool.Sender()
	log.Println("Ready for notifies")

	api.Listen()
	wpool.Close()
	wpool.Wait()

	return nil
}
