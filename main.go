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
	"log"
	"os"
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

	err := config.AppConf.Load()
	if err != nil {
		return err
	}

	err = senders.Init()
	if err != nil {
		return err
	}

	db.Init()

	err = db.Stor.Db().DropTableIfExists(&service.Service{}).CreateTable(&service.Service{}).Error
	if err != nil {
		return err
	}

	err = db.Stor.Db().DropTableIfExists(&notify.Notify{}).CreateTable(&notify.Notify{}).Error
	if err != nil {
		return err
	}

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
