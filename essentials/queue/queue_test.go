package queue

import (
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/queue/eq_pipe"
	"testing"
	"time"
)

func TestQueueImpl_Enqueue(t *testing.T) {
	l := esl.Default()
	processor := func(userId string) {
		l.Info("Greet", esl.String("UserId", userId))
	}
	factory := eq_pipe.NewSimple(l)
	queue := New(l, 2, factory, processor)

	queue.Enqueue("U-001")

	b1 := queue.Batch("B01")
	b1.Enqueue("UB-001")
	b1.Enqueue("UB-002")

	queue.Enqueue("U-002")

	queue.Wait()
}

func TestQueueImpl_Suspend(t *testing.T) {
	waitUnit := 50 * time.Millisecond

	l := esl.Default()
	processor := func(userId string) {
		time.Sleep(waitUnit)
		l.Info("Greet", esl.String("UserId", userId))
	}
	factory := eq_pipe.NewSimple(l)
	queue := New(l, 1, factory, processor)

	queue.Enqueue("U-001")
	queue.Enqueue("U-002")
	queue.Enqueue("U-003")
	queue.Enqueue("U-004")

	session, err := queue.Suspend()
	if err != eq_pipe.ErrorPreserveIsNotSupported {
		t.Error(session, err)
	}
}
