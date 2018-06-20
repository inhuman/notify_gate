package pool

import (
	"jgit.me/tools/notify_gate/senders"
	"jgit.me/tools/notify_gate/utils"
	"time"
)

var NPool = &NotifyPool{}

func init() {
	NPool.ToSend = make(chan *senders.Notify, 1000)
	NPool.Done = make(chan bool)
}

type NotifyPool struct {
	ToSend chan *senders.Notify
	Done   chan bool
}

func (np *NotifyPool) Add(n *senders.Notify) error {
	np.ToSend <- n

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

func Run() {
	utils.ShowDebugMessage("Starting notify pool")

	for {
		select {
		case n, ok := <-NPool.ToSend:
			if ok {
				go senders.Send(n)
			} else {
				utils.ShowDebugMessage("Can not read from notify channel")
			}

		case <-NPool.Done:
			return
		}

		<- time.After(500 * time.Microsecond)
	}
}
