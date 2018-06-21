package pool

import (
	"testing"
	"jgit.me/tools/notify_gate/notify"
	"jgit.me/tools/notify_gate/db"
	"jgit.me/tools/notify_gate/config"
	"jgit.me/tools/notify_gate/workerpool"
	"time"
	"jgit.me/tools/notify_gate/senders"
	"fmt"
	"strconv"
)

func BenchmarkNotifyPool_Add(b *testing.B) {

	//err := config.AppConf.Load()
	//
	//if err != nil {
	//	b.Log(err)
	//}
	//
	//db.Init()
	//
	//go Saver()
	//
	//senders.Providers["test"] = func(n *notify.Notify) error {
	//	time.Sleep(1000 * time.Millisecond)
	//	return nil
	//}
	//
	//n := &notify.Notify{
	//	Type:    "test",
	//	Message: "test message",
	//}
	//
	//for i := 0; i < 100000; i++ {
	//	NPool.AddToSave(n)
	//}
}

func TestNotifyPool_Add(t *testing.T) {

	err := config.AppConf.Load()

	if err != nil {
		t.Log(err)
	}

	db.Init()
	db.Stor.Migrate(notify.Notify{})


	wpool := workerpool.NewPool(5)

	go Saver(wpool)
	go Sender()


	senders.Providers["test"] = func(n *notify.Notify) error {
		time.Sleep(500 * time.Millisecond)
		fmt.Println("sent message " + n.Message)
		return nil
	}

	for i := 0; i < 10; i++ {
		n := &notify.Notify{
			Type:    "test",
			Message: "test message " + strconv.Itoa(i),
		}
		NPool.AddToSave(n)
	}

	for {
		n := notify.GetNotify()
		if n.ID == 0 {
			<- time.After(1 * time.Second)
			NPool.Done <- true
		}
	}
}
