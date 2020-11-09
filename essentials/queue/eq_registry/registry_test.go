package eq_registry

import (
	"github.com/watermint/toolbox/essentials/ambient/ea_indicator"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/queue/eq_bundle"
	"github.com/watermint/toolbox/essentials/queue/eq_mould"
	"github.com/watermint/toolbox/essentials/queue/eq_pipe"
	"github.com/watermint/toolbox/essentials/queue/eq_progress"
	"testing"
)

func TestRegImpl_Define(t *testing.T) {
	logger := esl.Default()
	handler := eq_progress.NewProgress(ea_indicator.Global())
	storage := eq_bundle.NewSimple(logger, eq_bundle.FetchRandom, handler, eq_pipe.NewTransientSimple(logger))
	reg := New(storage, nil)
	if _, found := reg.Get("no_existent"); found {
		t.Error(t)
	}
	processed := 0
	processor := func(userId string) {
		logger.Info("Processing", esl.String("userId", userId))
		processed++
	}
	reg.Define("alpha", processor, eq_mould.Opts{})
	if mould, found := reg.Get("alpha"); !found {
		t.Error(found)
	} else {
		mould.Pour("U001")
		if barrel, found := storage.Fetch(); !found {
			t.Error(t)
		} else {
			mould.Process(barrel)
			if processed < 1 {
				t.Error(processed)
			}
		}
	}
}
