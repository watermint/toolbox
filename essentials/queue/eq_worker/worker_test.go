package eq_worker

import (
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/queue/eq_bundle"
	"github.com/watermint/toolbox/essentials/queue/eq_mould"
	"github.com/watermint/toolbox/essentials/queue/eq_pipe"
	"testing"
	"time"
)

func TestWorkerImpl(t *testing.T) {
	bundle := eq_bundle.NewSimple(esl.Default(), eq_pipe.NewSimple())
	processor := func(v string) {
		l := esl.Default()
		l.Info("Process", esl.String("v", v))
	}
	mould := eq_mould.New(bundle, processor)

	c := make(chan eq_bundle.Data)
	worker := New(esl.Default(), mould, c)
	worker.Startup()
	worker.Startup()

	go func() {
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
