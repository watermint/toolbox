package eq_sequence

import (
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/queue/eq_pipe"
	"github.com/watermint/toolbox/essentials/queue/eq_queue"
	"go.uber.org/atomic"
	"testing"
)

func TestSeqImpl_DoThen(t *testing.T) {
	l := esl.Default()
	seq := New(
		eq_queue.Logger(l),
		eq_queue.Factory(eq_pipe.NewTransientSimple(l)),
		eq_queue.NumWorker(4),
	)

	stg1HelloCkpt := atomic.NewInt32(0)
	stg1Hello := func(msg string, s Stage) {
		l.Info("Stage1: Hello", esl.String("msg", msg), esl.Int32("ckpt", stg1HelloCkpt.Inc()))

		qg := s.Get("greet")
		qg.Enqueue("[" + msg + "]")
	}
	stg1GreetCkpt := atomic.NewInt32(0)
	stg1Greet := func(msg string) {
		l.Info("Stage1: Greet", esl.String("msg", msg), esl.Int32("ckpt", stg1GreetCkpt.Inc()))
	}

	stg2HelloCkpt := atomic.NewInt32(0)
	stg2Hello := func(msg string) {
		l.Info("Stage2: Hello", esl.String("msg", msg), esl.Int32("ckpt", stg2HelloCkpt.Inc()))
	}
	stg2GreetCkpt := atomic.NewInt32(0)
	stg2Greet := func(msg string, location string) {
		l.Info("Stage2: Greet", esl.String("msg", msg), esl.Int32("ckpt", stg2GreetCkpt.Inc()), esl.String("location", location))
	}

	stg3HelloCkpt := atomic.NewInt32(0)
	stg3Hello := func(msg string) {
		l.Info("Stage3: Hello", esl.String("msg", msg), esl.Int32("ckpt", stg3HelloCkpt.Inc()))
	}

	seq.DoThen(func(s Stage) {
		s.Define("hello", stg1Hello, s)
		s.Define("greet", stg1Greet)

		qh := s.Get("hello")
		qh.Enqueue("Orange")
		qh.Enqueue("Mango")

	}).DoThen(func(s Stage) {
		if x := stg1HelloCkpt.Load(); x != 2 {
			t.Error(x)
		}
		if x := stg1GreetCkpt.Load(); x != 2 {
			t.Error(x)
		}

		s.Define("hello", stg2Hello)
		s.Define("greet", stg2Greet, "Tokyo")

		qh := s.Get("hello")
		qg := s.Get("greet")

		qh.Enqueue("Peanuts")
		qh.Enqueue("Almond")

		qg.Batch("Fruit").Enqueue("Melon")
		qg.Batch("Fruit").Enqueue("Banana")
		qg.Batch("Beverage").Enqueue("Wine")

	}).Do(func(s Stage) {
		if x := stg1HelloCkpt.Load(); x != 2 {
			t.Error(x)
		}
		if x := stg1GreetCkpt.Load(); x != 2 {
			t.Error(x)
		}
		if x := stg2HelloCkpt.Load(); x != 2 {
			t.Error(x)
		}
		if x := stg2GreetCkpt.Load(); x != 3 {
			t.Error(x)
		}

		s.Define("hello", stg3Hello)

		q := s.Get("hello")
		q.Enqueue("Tokyo")
		q.Enqueue("Osaka")
		q.Enqueue("Okinawa")
		q.Enqueue("Hokkaido")
	})

	if x := stg1HelloCkpt.Load(); x != 2 {
		t.Error(x)
	}
	if x := stg1GreetCkpt.Load(); x != 2 {
		t.Error(x)
	}
	if x := stg2HelloCkpt.Load(); x != 2 {
		t.Error(x)
	}
	if x := stg2GreetCkpt.Load(); x != 3 {
		t.Error(x)
	}
	if x := stg3HelloCkpt.Load(); x != 4 {
		t.Error(x)
	}
}
