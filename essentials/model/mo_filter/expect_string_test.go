package mo_filter

import (
	"testing"
)

type ExpectStringHello struct {
	Message string
}

func (z ExpectStringHello) String() string {
	return z.Message
}

func TestExpectString(t *testing.T) {
	es := ExpectString("hello", func(s string) bool {
		return s == "hello"
	})
	if !es {
		t.Error(es)
	}

	es = ExpectString(ExpectStringHello{Message: "hello"}, func(s string) bool {
		return s == "hello"
	})
	if !es {
		t.Error(es)
	}

	es = ExpectString(123, func(s string) bool {
		return s == "123"
	})
	if !es {
		t.Error(es)
	}
}
