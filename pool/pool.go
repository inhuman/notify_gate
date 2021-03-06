package pool

import (
	"github.com/inhuman/notify_gate/notify"
	"github.com/inhuman/notify_gate/senders"
	"github.com/inhuman/notify_gate/utils"
	"github.com/inhuman/notify_gate/workerpool"
	"log"
	"time"
)

var nPool = &notifyPool{}

func init() {
	nPool.ToSend = make(chan *notify.Notify, 1024)
	nPool.ToSave = make(chan *notify.Notify, 1024)
	nPool.ToDelete = make(chan *notify.Notify, 1024)
	nPool.Done = make(chan bool)
}

type notifyPool struct {
	ToSend   chan *notify.Notify
	ToSave   chan *notify.Notify
	ToDelete chan *notify.Notify
	Done     chan bool
}

// AddToSave is used for adding notify into queue to save
func AddToSave(n *notify.Notify) error {
	nPool.ToSave <- n
	return nil
}

// Saver is used for process notifyPool.ToSave channel
func Saver(wpool *workerpool.Pool) {
	for {
		select {
		case n, ok := <-nPool.ToSave:
			if ok {
				wpool.Exec(n)
			} else {
				log.Println("Can not read from notify channel")
			}
		case <-nPool.Done:
			wpool.Close()
			wpool.Wait()
			return
		}
	}
}

// Sender is used for read notify from db and send them
func Sender() {
L:
	for {
		select {
		case <-nPool.Done:
			break L
		default:
			n := notify.GetNotify()
			if n.ID != 0 {
				err := senders.Send(n)
				utils.CheckErrorMessage(n.Type+" sending error:", err)
				n.Delete()
			}
		}
		<-time.After(1000 * time.Millisecond)
	}
	return
}
