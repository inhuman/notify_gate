package pool

import (
	"jgit.me/tools/notify_gate/utils"
	"jgit.me/tools/notify_gate/notify"
	"jgit.me/tools/notify_gate/workerpool"
	"time"
	"fmt"
	"jgit.me/tools/notify_gate/senders"
	"jgit.me/tools/notify_gate/db"
)

var NPool = &NotifyPool{}

func init() {
	NPool.ToSend = make(chan *notify.Notify, 1024)
	NPool.ToSave = make(chan *notify.Notify, 1024)
	NPool.ToDelete = make(chan *notify.Notify, 1024)
	NPool.Done = make(chan bool)
}

type NotifyPool struct {
	ToSend   chan *notify.Notify
	ToSave   chan *notify.Notify
	ToDelete chan *notify.Notify
	Done     chan bool
}

func (np *NotifyPool) AddToSave(n *notify.Notify) error {
	np.ToSave <- n
	return nil
}

func (np *NotifyPool) AddToSend(n *notify.Notify) error {
	np.ToSend <- n
	return nil
}

func Saver(wpool *workerpool.Pool) {
	utils.ShowDebugMessage("Starting notify saver")

	for {
		select {
		case n, ok := <-NPool.ToSave:
			if ok {
				wpool.Exec(n)
			} else {
				fmt.Println("Can not read from notify channel")
			}
		case <-NPool.Done:
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
		case <-NPool.Done:
			break L
		default:
			n := notify.GetNotify()
			if n.ID != 0 {
				senders.Send(n)
				db.Stor.Db().Unscoped().Delete(n)
			}
		}
		<-time.After(1000 * time.Millisecond)
	}
	return
}
