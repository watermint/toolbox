package eq_worker

import (
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/queue/eq_bundle"
	"github.com/watermint/toolbox/essentials/queue/eq_pipe"
	"github.com/watermint/toolbox/essentials/queue/eq_registry"
	"testing"
	"time"
)

func TestWorkerImpl(t *testing.T) {
	bundle := eq_bundle.NewSimple(esl.Default(), nil, eq_pipe.NewTransientSimple(esl.Default()))
	processor := func(v string) {
		l := esl.Default()
		l.Info("Process", esl.String("v", v))
	}

	c := make(chan eq_bundle.Barrel)
	reg := eq_registry.New(bundle)
	reg.Define("m", processor)
	mould, found := reg.Get("m")
	if !found {
		t.Error(found)
		return
	}
	worker := New(esl.Default(), reg, c)
	worker.Startup(1)
	worker.Startup(2)

	l := esl.Default()
	go func() {
		defer func() {
			if err := recover(); err != nil {
				l.Debug("In case of timing issue; The channel already closed", esl.Any("err", err))
			}
		}()
		for {
			d, found := bundle.Fetch()
			if found {
				c <- d
			} else {
				time.Sleep(10 * time.Millisecond)
			}
		}
	}()

	batch1 := mould.Batch("Batch1")
	batch1.Pour("Hello")
	batch1.Pour("World")

	batch2 := mould.Batch("Batch2")
	batch2.Pour("こんにちは")
	batch2.Pour("世界")

	time.Sleep(40 * time.Millisecond)

	batch3 := mould.Batch("Batch3")
	batch3.Pour("Hallo")
	batch3.Pour("Welt")

	close(c)

	worker.Wait()
}
