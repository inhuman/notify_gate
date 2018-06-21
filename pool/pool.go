package pool

import (
	"jgit.me/tools/notify_gate/utils"
	"jgit.me/tools/notify_gate/notify"
	"jgit.me/tools/notify_gate/workerpool"
	"time"
	"jgit.me/tools/notify_gate/senders"
	"jgit.me/tools/notify_gate/db"
	"fmt"
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

	// TODO: test write to closed channel
	//defer func() error {
	//	if r := recover(); r != nil {
	//		if err, ok := r.(error); ok {
	//			return err
	//		}
	//	}
	//	return nil
	//}()
	//
	return nil
}

func (np *NotifyPool) AddToSend(n *notify.Notify) error {
	np.ToSend <- n
	return nil
}

func Saver(wpool *workerpool.Pool) {
	utils.ShowDebugMessage("Starting notify pool saver")

	for {
		select {
		case n, ok := <-NPool.ToSave:
			if ok {
				fmt.Println("Saver.ToSave", n)
				wpool.Exec(n)
			} else {
				utils.ShowDebugMessage("Can not read from notify channel")
			}

		case <-NPool.Done:
			wpool.Close()
			wpool.Wait()
			return
		}
	}
}

func Sender() {
	for {

		n := notify.GetNotify()

		if n.ID != 0 {
			senders.Send(n)
			db.Stor.Db().Unscoped().Delete(n)
		}

		<-time.After(1000 * time.Millisecond)
	}
}
