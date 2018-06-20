package pool

import (
	"testing"
	"jgit.me/tools/notify_gate/notify"
	"jgit.me/tools/notify_gate/senders"
	"time"
)

func BenchmarkNotifyPool_Add(b *testing.B) {

	go Run()

	senders.Providers["test"] = func(n *notify.Notify) error {
		time.Sleep(1000 * time.Millisecond)
		return nil
	}

	n := &notify.Notify{
		Type:    "test",
		Message: "test message",
	}

	for i := 0; i < b.N; i++ {
		NPool.Add(n)
	}
}
