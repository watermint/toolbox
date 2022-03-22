package eq_bundle

import (
	"github.com/watermint/toolbox/essentials/log/esl"
	"reflect"
	"testing"
)

func TestNewPetty(t *testing.T) {
	bundle := NewPetty(esl.Default(), nil, 5)

	d1_1 := NewBarrel("m1", "b1", []byte("D01-01"))
	if total := bundle.Size(); total != 0 {
		t.Error(total)
	}

	bundle.Enqueue(d1_1)

	if total := bundle.Size(); total != 1 {
		t.Error(total)
	}

	if d, found := bundle.Fetch(); !found {
		t.Error(found)
	} else if !reflect.DeepEqual(d1_1, d) {
		t.Error(d)
	}

	bundle.Complete(d1_1)

	bundle.Close()
}
