package queue

import (
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/queue/eq_pipe"
	"testing"
)

func TestQueueImpl_Enqueue(t *testing.T) {
	l := esl.Default()
	processor := func(userId string) {
		l.Info("Greet", esl.String("UserId", userId))
	}
	factory := eq_pipe.NewSimple(l)
	queue := New(l, factory, processor)

	queue.Enqueue("U-001")

	b1 := queue.Batch("B01")
	b1.Enqueue("UB-001")
	b1.Enqueue("UB-002")

	queue.Enqueue("U-002")

	queue.Wait()
}
