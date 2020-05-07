package catalogue

import (
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"testing"
)

func TestAutoDetectedMessageObjects(t *testing.T) {
	mos := AutoDetectedMessageObjects()
	for _, mo := range mos {
		mo1 := app_msg.Apply(mo)
		ms := app_msg.Messages(mo1)
		for _, m := range ms {
			m.Key()
		}
	}
}
