package rc_spec

import (
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/quality/infra/qt_control"
	"strings"
	"testing"
)

// Mock recipe for testing
type mockRecipe struct {
	Value string
}

func (m *mockRecipe) Preset() {
	m.Value = "test_value"
}
func (m *mockRecipe) Exec(c app_control.Control) error { return nil }
func (m *mockRecipe) Test(c app_control.Control) error { return nil }

// Mock annotated recipe
type mockAnnotatedRecipe struct {
	mockRecipe
}

func (m *mockAnnotatedRecipe) Seed() rc_recipe.Recipe {
	return &m.mockRecipe
}

func (m *mockAnnotatedRecipe) IsExperimental() bool    { return false }
func (m *mockAnnotatedRecipe) IsIrreversible() bool    { return false }
func (m *mockAnnotatedRecipe) IsTransient() bool       { return false }
func (m *mockAnnotatedRecipe) IsSecret() bool          { return false }
func (m *mockAnnotatedRecipe) IsConsole() bool         { return false }
func (m *mockAnnotatedRecipe) IsLicenseRequired() bool { return false }

func TestNew(t *testing.T) {
	recipe := &mockRecipe{}
	spec := New(recipe)

	if spec == nil {
		t.Error("Expected non-nil spec")
	}
}

func TestNewSelfContained(t *testing.T) {
	// Test with regular recipe
	recipe := &mockRecipe{}
	spec := newSelfContained(recipe)

	if spec == nil {
		t.Error("Expected non-nil spec")
	}

	// Test with annotated recipe
	annotated := &mockAnnotatedRecipe{}
	annotatedSpec := newSelfContained(annotated)

	if annotatedSpec == nil {
		t.Error("Expected non-nil spec for annotated recipe")
	}
}

func TestSpecValueSelfContained_SpecId(t *testing.T) {
	recipe := &mockRecipe{}
	spec := newSelfContained(recipe).(*specValueSelfContained)

	specId := spec.SpecId()
	if specId == "" {
		t.Error("Expected non-empty spec ID")
	}

	if !strings.Contains(specId, "mock") {
		t.Errorf("Expected spec ID to contain recipe name, got %s", specId)
	}
}

func TestSpecValueSelfContained_Path(t *testing.T) {
	recipe := &mockRecipe{}
	spec := newSelfContained(recipe).(*specValueSelfContained)

	_, name := spec.Path()
	if name == "" {
		t.Error("Expected non-empty name")
	}

	// Path could be empty for test recipes
	if name != "mock_recipe" {
		t.Errorf("Expected name to be 'mock_recipe', got %s", name)
	}
}

func TestSpecValueSelfContained_IsLicenseRequired(t *testing.T) {
	// Test regular recipe
	recipe := &mockRecipe{}
	spec := newSelfContained(recipe).(*specValueSelfContained)

	if spec.IsLicenseRequired() {
		t.Error("Expected license not required for regular recipe")
	}

	// Test annotated recipe
	annotated := &mockAnnotatedRecipe{}
	annotatedSpec := newSelfContained(annotated).(*specValueSelfContained)

	if annotatedSpec.IsLicenseRequired() {
		t.Error("Expected license not required for mock annotated recipe")
	}
}

func TestSpecValueSelfContained_IsPruned(t *testing.T) {
	recipe := &mockRecipe{}
	spec := newSelfContained(recipe).(*specValueSelfContained)

	// Mock recipes should not be pruned
	if spec.IsPruned() {
		t.Error("Expected mock recipe not to be pruned")
	}
}

func TestSpecValueSelfContained_MarkSpecChange(t *testing.T) {
	recipe := &mockRecipe{}
	spec := newSelfContained(recipe).(*specValueSelfContained)

	originalSpecChange := spec.IsSpecChange()
	newSpec := spec.MarkSpecChange()

	if newSpec == nil {
		t.Error("Expected non-nil spec from MarkSpecChange")
	}

	// Check if the returned spec has specChange marked
	newSpecTyped := newSpec.(specValueSelfContained)
	if !newSpecTyped.IsSpecChange() {
		t.Error("Expected IsSpecChange to be true in returned spec after MarkSpecChange")
	}

	// Original should remain unchanged (value receiver)
	if spec.IsSpecChange() != originalSpecChange {
		t.Error("Expected original spec to remain unchanged")
	}
}

func TestSpecValueSelfContained_FormerPaths(t *testing.T) {
	recipe := &mockRecipe{}
	spec := newSelfContained(recipe).(*specValueSelfContained)

	formerPaths := spec.FormerPaths()
	if formerPaths == nil {
		t.Error("Expected non-nil former paths slice")
	}
	// Mock recipes typically have no former paths
	if len(formerPaths) != 0 {
		t.Errorf("Expected empty former paths for mock recipe, got %d", len(formerPaths))
	}
}

func TestSpecValueSelfContained_ErrorHandlers(t *testing.T) {
	recipe := &mockRecipe{}
	spec := newSelfContained(recipe).(*specValueSelfContained)

	handlers := spec.ErrorHandlers()
	if handlers == nil {
		t.Error("Expected non-nil error handlers slice")
	}
	// Mock recipes typically have no error handlers
}

func TestSpecValueSelfContained_IsFlags(t *testing.T) {
	recipe := &mockRecipe{}
	spec := newSelfContained(recipe).(*specValueSelfContained)

	// Test all flag methods
	if spec.IsExperimental() {
		t.Error("Expected mock recipe not to be experimental")
	}

	if spec.IsIrreversible() {
		t.Error("Expected mock recipe not to be irreversible")
	}

	if spec.IsTransient() {
		t.Error("Expected mock recipe not to be transient")
	}

	if spec.IsSecret() {
		t.Error("Expected mock recipe not to be secret")
	}

	if spec.IsConsole() {
		t.Error("Expected mock recipe not to be console")
	}

	if spec.IsSpecChange() {
		t.Error("Expected mock recipe not to have spec change initially")
	}
}

func TestSpecValueSelfContained_CaptureRestore(t *testing.T) {
	err := qt_control.WithControl(func(c app_control.Control) error {
		recipe := &mockRecipe{}
		spec := newSelfContained(recipe).(*specValueSelfContained)

		// Test capture
		captured, err := spec.Capture(c)
		if err != nil {
			t.Error(err)
			return err
		}
		if captured == nil {
			t.Error("Expected captured data to be returned")
		}

		return nil
	})
	if err != nil {
		t.Error(err)
	}
}

func TestSpecValueSelfContained_New(t *testing.T) {
	recipe := &mockRecipe{}
	spec := newSelfContained(recipe).(*specValueSelfContained)

	newSpec := spec.New()
	if newSpec == nil {
		t.Error("Expected non-nil spec from New")
	}

	// Should be a different instance
	if newSpec == spec {
		t.Error("Expected New to return a different instance")
	}
}

func TestSpecValueSelfContained_Value(t *testing.T) {
	recipe := &mockRecipe{}
	spec := newSelfContained(recipe).(*specValueSelfContained)

	// Test getting a field value
	value := spec.Value("Value")
	if value == nil {
		t.Error("Expected non-nil value for existing field")
	}

	// Test non-existing field
	nonExistingValue := spec.Value("NonExistingField")
	if nonExistingValue != nil {
		t.Error("Expected nil value for non-existing field")
	}
}