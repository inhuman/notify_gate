package main

import (
	"jgit.me/tools/notify_gate/api"
	"jgit.me/tools/notify_gate/pool"
	"jgit.me/tools/notify_gate/db"
	"jgit.me/tools/notify_gate/service"
	"jgit.me/tools/notify_gate/cache"
	"jgit.me/tools/notify_gate/notify"
	"jgit.me/tools/notify_gate/workerpool"
)

func main() {

	db.Init()

	db.Stor.Migrate(service.Service{})
	db.Stor.Migrate(notify.Notify{})

	cache.BuildTokenCache()

	wpool := workerpool.NewPool(5)

	go pool.Saver(wpool)
	go pool.Sender()
	api.Listen()

	wpool.Close()
	wpool.Wait()
}


//TODO: implement tests for api
//TODO: implement tests for tokens cache
//TODO: implement tests for config load
//TODO: implement tests for notify pool
//TODO: implement tests for telegram
//TODO: implement tests for slack
//TODO: implement tests for register service

//TODO: implement reconnect to db if its disconnect