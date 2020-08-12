package eu_uuid

import (
	"github.com/google/uuid"
	"testing"
)

func TestNewV4(t *testing.T) {
	u := NewV4()
	if u.Version() != 4 {
		t.Error(u.Version())
	}
	if x := u.Variant(); x != uuid.RFC4122 {
		t.Error(x)
	}
}
