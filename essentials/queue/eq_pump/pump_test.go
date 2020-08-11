package eq_pump

import (
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/queue/eq_bundle"
	"github.com/watermint/toolbox/essentials/queue/eq_pipe"
	"testing"
)

func TestPumpImpl_Start(t *testing.T) {
	bundle := eq_bundle.NewSimple(esl.Default(), eq_pipe.NewSimple())
	pump := New(esl.Default(), bundle)
	pump.Start()
	pump.Close()
}
