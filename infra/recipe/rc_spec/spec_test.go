package rc_spec

import (
	"testing"
	"github.com/watermint/toolbox/infra/control/app_control"
)

// Mock recipe for testing
type mockRecipe struct{}

func (m *mockRecipe) Preset() {}
func (m *mockRecipe) Exec(c app_control.Control) error { return nil }
func (m *mockRecipe) Test(c app_control.Control) error { return nil }

func TestNew(t *testing.T) {
	recipe := &mockRecipe{}
	spec := New(recipe)
	
	if spec == nil {
		t.Error("Expected non-nil spec")
	}
}