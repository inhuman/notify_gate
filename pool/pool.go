package pool

import (
	"fmt"
	"jgit.me/tools/notify_gate/db"
	"jgit.me/tools/notify_gate/notify"
	"jgit.me/tools/notify_gate/senders"
	"jgit.me/tools/notify_gate/utils"
	"jgit.me/tools/notify_gate/workerpool"
	"time"
)


var notifyPool = &NotifyPool{}

func init() {
	notifyPool.ToSend = make(chan *notify.Notify, 1024)
	notifyPool.ToSave = make(chan *notify.Notify, 1024)
	notifyPool.ToDelete = make(chan *notify.Notify, 1024)
	notifyPool.Done = make(chan bool)
}

type NotifyPool struct {
	ToSend   chan *notify.Notify
	ToSave   chan *notify.Notify
	ToDelete chan *notify.Notify
	Done     chan bool
}

func AddToSave(n *notify.Notify) error {
	notifyPool.ToSave <- n
	return nil
}

func Saver(wpool *workerpool.Pool) {
	utils.ShowDebugMessage("Starting notify saver")

	for {
		select {
		case n, ok := <-notifyPool.ToSave:
			if ok {
				wpool.Exec(n)
			} else {
				fmt.Println("Can not read from notify channel")
			}
		case <-notifyPool.Done:
			wpool.Close()
			wpool.Wait()
			return
		}
	}
}

func Sender() {
L:
	for {
		select {
		case <-notifyPool.Done:
			break L
		default:
			n := notify.GetNotify()
			if n.ID != 0 {
				err := senders.Send(n)
				if err != nil {
					fmt.Print(err)
				}
				db.Stor.Db().Unscoped().Delete(n)
			}
		}
		<-time.After(1000 * time.Millisecond)
	}
	return
}
