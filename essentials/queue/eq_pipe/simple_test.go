package eq_pipe

import (
	"github.com/watermint/toolbox/essentials/log/esl"
	"reflect"
	"testing"
)

func TestSimpleImpl_Size(t *testing.T) {
	factory := NewTransientSimple(esl.Default())
	pipe := factory.New("root")
	if x := pipe.Size(); x != 0 {
		t.Error(x)
	}
	pipe.Enqueue([]byte("B01"))
	if x := pipe.Size(); x != 1 {
		t.Error(x)
	}
	pipe.Enqueue([]byte("B02"))
	if x := pipe.Size(); x != 2 {
		t.Error(x)
	}
}

func TestSimpleImpl_Delete(t *testing.T) {
	factory := NewTransientSimple(esl.Default())
	pipe := factory.New("root")
	if x := pipe.Size(); x != 0 {
		t.Error(x)
	}
	pipe.Enqueue([]byte("B01"))
	if x := pipe.Size(); x != 1 {
		t.Error(x)
	}
	pipe.Delete([]byte("B01"))
	if x := pipe.Size(); x != 0 {
		t.Error(x)
	}
}

func TestSimpleImpl_Dequeue(t *testing.T) {
	factory := NewTransientSimple(esl.Default())
	pipe := factory.New("root")
	if x := pipe.Size(); x != 0 {
		t.Error(x)
	}
	pipe.Enqueue([]byte("B01"))
	if x := pipe.Size(); x != 1 {
		t.Error(x)
	}
	d := pipe.Dequeue()
	if !reflect.DeepEqual(d, []byte("B01")) {
		t.Error(d)
	}
	if x := pipe.Size(); x != 0 {
		t.Error(x)
	}
}
