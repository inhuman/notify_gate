package main

import (
	"jgit.me/tools/notify_gate/api"
	"jgit.me/tools/notify_gate/pool"
	"jgit.me/tools/notify_gate/db"
	"jgit.me/tools/notify_gate/service"
	"jgit.me/tools/notify_gate/cache"
	"jgit.me/tools/notify_gate/notify"
)

func main() {

	db.Init()

	db.Stor.Migrate(service.Service{})
	db.Stor.Migrate(notify.Notify{})

	cache.BuildTokenCache()

	go pool.Run()
	api.Listen()
}

//TODO: benchmark send notify through api
//TODO: benchmark write notify to db
//TODO: benchmark read notify to db


//TODO: implement tests for api
//TODO: implement tests for tokens cache
//TODO: implement tests for config load
//TODO: implement tests for notify pool
//TODO: implement tests for telegram
//TODO: implement tests for slack
//TODO: implement tests for register service