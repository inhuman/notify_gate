package pool

import (
	"jgit.me/tools/notify_gate/utils"
	"jgit.me/tools/notify_gate/notify"
	"jgit.me/tools/notify_gate/workerpool"
	"time"
	"jgit.me/tools/notify_gate/senders"
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
	np.ToSave <- n
	return nil
}

func Run(wpool *workerpool.Pool) {
	utils.ShowDebugMessage("Starting notify pool saver")

	for {
		select {
		case n, ok := <-NPool.ToSave:
			if ok {
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

func Send() {
	utils.ShowDebugMessage("Starting notify pool sender")

	for {
		select {
		case n, ok := <-NPool.ToSend:
			if ok {
				senders.Send(n)
				NPool.ToDelete <- n
			} else {
				utils.ShowDebugMessage("Can not read from notify channel")
			}

		case <-NPool.Done:
			return
		}
		<-time.After(500 * time.Millisecond)
	}
}

func Read() {
	for {
		for _, n := range notify.GetNotifies() {
			NPool.ToSend <- n
		}
		<-time.After(500 * time.Millisecond)
	}
}

func Delete()  {
	for {
		select {
			case n, ok := <-NPool.ToDelete:
				if ok {
					n.Delete()
				} else {
					utils.ShowDebugMessage("Can not read from notify channel")
				}
		}
	}
}