package eq_pump

import (
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/queue/eq_bundle"
	"github.com/watermint/toolbox/essentials/queue/eq_pipe"
	"testing"
)

func TestPumpImpl_StartClose(t *testing.T) {
	bundle := eq_bundle.NewSimple(esl.Default(), nil, eq_pipe.NewTransientSimple(esl.Default()))
	pump := New(esl.Default(), bundle)
	pump.Start()
	pump.Close()
}

func TestPumpImpl_StartShutdown(t *testing.T) {
	bundle := eq_bundle.NewSimple(esl.Default(), nil, eq_pipe.NewTransientSimple(esl.Default()))
	pump := New(esl.Default(), bundle)
	pump.Start()
	pump.Shutdown()
}
