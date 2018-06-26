package main

import (
	"jgit.me/tools/notify_gate/api"
	"jgit.me/tools/notify_gate/pool"
	"jgit.me/tools/notify_gate/db"
	"jgit.me/tools/notify_gate/service"
	"jgit.me/tools/notify_gate/cache"
	"jgit.me/tools/notify_gate/notify"
	"jgit.me/tools/notify_gate/workerpool"
	"jgit.me/tools/notify_gate/senders"
	"jgit.me/tools/notify_gate/config"
	"os"
	"fmt"
)

func main() {

	err := runApp()
	if err != nil {
		fmt.Println(err.Error())
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
	api.Listen()

	wpool.Close()
	wpool.Wait()

	return nil
}
