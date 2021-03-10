package eq_queue

import "testing"

func TestExecWithQueue_None(t *testing.T) {
	ExecWithQueue(func(q Definition) {
		// should work
	})
}

func TestExecWithQueue_WithTask(t *testing.T) {
	consumer := func(name string) {
		t.Log(name)
	}

	ExecWithQueue(func(qd Definition) {
		qd.Define("consumer", consumer)
		q := qd.Current().MustGet("consumer")
		q.Enqueue("hello")
		q.Enqueue("world")
	})
}
