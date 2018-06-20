package main

import (
	"jgit.me/tools/notify_gate/api"
	"jgit.me/tools/notify_gate/pool"
)

func main() {
	go pool.Run()
	api.Listen()
}

//TODO: benchmark send notify through api
//TODO: benchmark write notify to db
//TODO: benchmark read notify to db

//TODO: find out how often possible send  notify to telegram
//TODO: find out how often possible send  notify to slack

//TODO: implement tests for api
//TODO: implement tests for tokens cache
//TODO: implement tests for config load
//TODO: implement tests for notify pool
//TODO: implement tests for telegram
//TODO: implement tests for slack
//TODO: implement tests for register service