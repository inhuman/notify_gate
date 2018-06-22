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
)

func main() {

	db.Init()
	db.Stor.Db()
	db.Stor.Migrate(service.Service{})
	db.Stor.Migrate(notify.Notify{})
	senders.Init()

	cache.BuildTokenCache()
	wpool := workerpool.NewPool(5)

	go pool.Saver(wpool)
	go pool.Sender()
	api.Listen()

	wpool.Close()
	wpool.Wait()
}
