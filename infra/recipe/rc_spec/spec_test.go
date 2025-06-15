package rc_spec

import (
	"flag"
	"strings"
	"testing"

	"github.com/watermint/toolbox/essentials/encoding/es_json"
	"github.com/watermint/toolbox/essentials/go/es_lang"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/doc/dc_index"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/quality/infra/qt_control"
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
func (m *mockAnnotatedRecipe) IsDeprecated() bool      { return false }

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

func TestSpecValueSelfContained_Title(t *testing.T) {
	recipe := &mockRecipe{}
	spec := newSelfContained(recipe).(*specValueSelfContained)

	title := spec.Title()
	if title == nil {
		t.Error("Expected non-nil title")
	}
}

func TestSpecValueSelfContained_Desc(t *testing.T) {
	recipe := &mockRecipe{}
	spec := newSelfContained(recipe).(*specValueSelfContained)

	desc := spec.Desc()
	if desc == nil {
		t.Error("Expected non-nil description")
	}
}

func TestSpecValueSelfContained_Name(t *testing.T) {
	recipe := &mockRecipe{}
	spec := newSelfContained(recipe).(*specValueSelfContained)

	name := spec.Name()
	if name == "" {
		t.Error("Expected non-empty name")
	}
}

func TestSpecValueSelfContained_CliPath(t *testing.T) {
	recipe := &mockRecipe{}
	spec := newSelfContained(recipe).(*specValueSelfContained)

	cliPath := spec.CliPath()
	if cliPath == "" {
		t.Error("Expected non-empty CLI path")
	}
}

func TestSpecValueSelfContained_CliArgs(t *testing.T) {
	recipe := &mockRecipe{}
	spec := newSelfContained(recipe).(*specValueSelfContained)

	cliArgs := spec.CliArgs()
	if cliArgs == nil {
		t.Error("Expected non-nil CLI args")
	}
}

func TestSpecValueSelfContained_CliNote(t *testing.T) {
	recipe := &mockRecipe{}
	spec := newSelfContained(recipe).(*specValueSelfContained)

	cliNote := spec.CliNote()
	if cliNote == nil {
		t.Error("Expected non-nil CLI note")
	}
}

func TestSpecValueSelfContained_Reports(t *testing.T) {
	recipe := &mockRecipe{}
	spec := newSelfContained(recipe).(*specValueSelfContained)

	reports := spec.Reports()
	if reports == nil {
		t.Error("Expected non-nil reports")
	}
}

func TestSpecValueSelfContained_Feeds(t *testing.T) {
	recipe := &mockRecipe{}
	spec := newSelfContained(recipe).(*specValueSelfContained)

	feeds := spec.Feeds()
	if feeds == nil {
		t.Error("Expected non-nil feeds")
	}
}

func TestSpecValueSelfContained_GridDataInput(t *testing.T) {
	recipe := &mockRecipe{}
	spec := newSelfContained(recipe).(*specValueSelfContained)

	gridData := spec.GridDataInput()
	if gridData == nil {
		t.Error("Expected non-nil grid data input")
	}
}

func TestSpecValueSelfContained_GridDataOutput(t *testing.T) {
	recipe := &mockRecipe{}
	spec := newSelfContained(recipe).(*specValueSelfContained)

	gridData := spec.GridDataOutput()
	if gridData == nil {
		t.Error("Expected non-nil grid data output")
	}
}

func TestSpecValueSelfContained_TextInput(t *testing.T) {
	recipe := &mockRecipe{}
	spec := newSelfContained(recipe).(*specValueSelfContained)

	textInput := spec.TextInput()
	if textInput == nil {
		t.Error("Expected non-nil text input")
	}
}

func TestSpecValueSelfContained_JsonInput(t *testing.T) {
	recipe := &mockRecipe{}
	spec := newSelfContained(recipe).(*specValueSelfContained)

	jsonInput := spec.JsonInput()
	if jsonInput == nil {
		t.Error("Expected non-nil json input")
	}
}

func TestSpecValueSelfContained_ValueNames(t *testing.T) {
	recipe := &mockRecipe{}
	spec := newSelfContained(recipe).(*specValueSelfContained)

	valueNames := spec.ValueNames()
	if valueNames == nil {
		t.Error("Expected non-nil value names")
	}
}

func TestSpecValueSelfContained_ValueDesc(t *testing.T) {
	recipe := &mockRecipe{}
	spec := newSelfContained(recipe).(*specValueSelfContained)

	valueDesc := spec.ValueDesc("Value")
	if valueDesc == nil {
		t.Error("Expected non-nil value description")
	}
}

func TestSpecValueSelfContained_ValueDefault(t *testing.T) {
	recipe := &mockRecipe{}
	spec := newSelfContained(recipe).(*specValueSelfContained)

	valueDefault := spec.ValueDefault("Value")
	if valueDefault == nil {
		t.Error("Expected non-nil value default")
	}
}

func TestSpecValueSelfContained_ValueCustomDefault(t *testing.T) {
	recipe := &mockRecipe{}
	spec := newSelfContained(recipe).(*specValueSelfContained)

	customDefault := spec.ValueCustomDefault("Value")
	if customDefault == nil {
		t.Error("Expected non-nil custom default")
	}
}

func TestSpecValueSelfContained_Messages(t *testing.T) {
	recipe := &mockRecipe{}
	spec := newSelfContained(recipe).(*specValueSelfContained)

	messages := spec.Messages()
	if messages == nil {
		t.Error("Expected non-nil messages")
	}
}

func TestSpecValueSelfContained_ConnScopeMap(t *testing.T) {
	recipe := &mockRecipe{}
	spec := newSelfContained(recipe).(*specValueSelfContained)

	scopeMap := spec.ConnScopeMap()
	if scopeMap == nil {
		t.Error("Expected non-nil connection scope map")
	}
}

func TestSpecValueSelfContained_SpinDown(t *testing.T) {
	err := qt_control.WithControl(func(c app_control.Control) error {
		recipe := &mockRecipe{}
		spec := newSelfContained(recipe).(*specValueSelfContained)

		err := spec.SpinDown(c)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		t.Error("SpinDown should not error", err)
	}
}

func TestSpecValueSelfContained_ScopeLabels(t *testing.T) {
	recipe := &mockRecipe{}
	spec := newSelfContained(recipe).(*specValueSelfContained)

	labels := spec.ScopeLabels()
	if labels == nil {
		t.Error("Expected non-nil scope labels")
	}
}

// Test annotated recipe methods
type mockFullAnnotatedRecipe struct {
	mockRecipe
}

func (m *mockFullAnnotatedRecipe) Seed() rc_recipe.Recipe {
	return &m.mockRecipe
}

func (m *mockFullAnnotatedRecipe) IsExperimental() bool    { return true }
func (m *mockFullAnnotatedRecipe) IsIrreversible() bool    { return true }
func (m *mockFullAnnotatedRecipe) IsTransient() bool       { return true }
func (m *mockFullAnnotatedRecipe) IsSecret() bool          { return true }
func (m *mockFullAnnotatedRecipe) IsConsole() bool         { return true }
func (m *mockFullAnnotatedRecipe) IsLicenseRequired() bool { return true }
func (m *mockFullAnnotatedRecipe) IsDeprecated() bool      { return false }

func TestSpecValueSelfContained_AnnotatedFlags(t *testing.T) {
	recipe := &mockFullAnnotatedRecipe{}
	spec := newSelfContained(recipe).(*specValueSelfContained)

	if !spec.IsExperimental() {
		t.Error("Expected experimental flag to be true")
	}

	if !spec.IsIrreversible() {
		t.Error("Expected irreversible flag to be true")
	}

	if !spec.IsTransient() {
		t.Error("Expected transient flag to be true")
	}

	if !spec.IsSecret() {
		t.Error("Expected secret flag to be true")
	}

	if !spec.IsConsole() {
		t.Error("Expected console flag to be true")
	}

	if !spec.IsLicenseRequired() {
		t.Error("Expected license required flag to be true")
	}
}

func TestSpecValueSelfContained_Remarks(t *testing.T) {
	t.Run("experimental and irreversible", func(t *testing.T) {
		recipe := &mockFullAnnotatedRecipe{}
		spec := newSelfContained(recipe).(*specValueSelfContained)

		remarks := spec.Remarks()
		if remarks == nil {
			t.Error("Expected non-nil remarks for experimental and irreversible")
		}
	})

	t.Run("irreversible only", func(t *testing.T) {
		recipe := &mockAnnotatedRecipe{}
		recipe.mockRecipe = mockRecipe{}
		spec := newSelfContained(recipe).(*specValueSelfContained)
		spec.annotation = &mockPartialAnnotatedRecipe{irreversible: true}

		remarks := spec.Remarks()
		if remarks == nil {
			t.Error("Expected non-nil remarks for irreversible")
		}
	})

	t.Run("experimental only", func(t *testing.T) {
		recipe := &mockAnnotatedRecipe{}
		recipe.mockRecipe = mockRecipe{}
		spec := newSelfContained(recipe).(*specValueSelfContained)
		spec.annotation = &mockPartialAnnotatedRecipe{experimental: true}

		remarks := spec.Remarks()
		if remarks == nil {
			t.Error("Expected non-nil remarks for experimental")
		}
	})

	t.Run("no special flags", func(t *testing.T) {
		recipe := &mockRecipe{}
		spec := newSelfContained(recipe).(*specValueSelfContained)

		remarks := spec.Remarks()
		if remarks == nil {
			t.Error("Expected non-nil remarks")
		}
	})
}

// Helper for partial annotation testing
type mockPartialAnnotatedRecipe struct {
	mockRecipe
	experimental bool
	irreversible bool
}

func (m *mockPartialAnnotatedRecipe) Seed() rc_recipe.Recipe      { return &m.mockRecipe }
func (m *mockPartialAnnotatedRecipe) IsExperimental() bool        { return m.experimental }
func (m *mockPartialAnnotatedRecipe) IsIrreversible() bool        { return m.irreversible }
func (m *mockPartialAnnotatedRecipe) IsTransient() bool           { return false }
func (m *mockPartialAnnotatedRecipe) IsSecret() bool              { return false }
func (m *mockPartialAnnotatedRecipe) IsConsole() bool             { return false }
func (m *mockPartialAnnotatedRecipe) IsLicenseRequired() bool     { return false }
func (m *mockPartialAnnotatedRecipe) IsDeprecated() bool          { return false }

func TestSpecValueSelfContained_Doc(t *testing.T) {
	recipe := &mockRecipe{}
	spec := newSelfContained(recipe).(*specValueSelfContained)

	err := qt_control.WithControl(func(c app_control.Control) error {
		doc := spec.Doc(c.UI())
		if doc == nil {
			t.Error("Expected non-nil doc")
		}
		if doc.Name != spec.Name() {
			t.Error("Expected doc name to match spec name")
		}
		if doc.Path != spec.CliPath() {
			t.Error("Expected doc path to match CLI path")
		}
		return nil
	})
	if err != nil {
		t.Error(err)
	}
}

func TestSpecValueSelfContained_PrintUsage(t *testing.T) {
	recipe := &mockRecipe{}
	spec := newSelfContained(recipe).(*specValueSelfContained)

	err := qt_control.WithControl(func(c app_control.Control) error {
		// Should not panic
		spec.PrintUsage(c.UI())
		return nil
	})
	if err != nil {
		t.Error(err)
	}
}

func TestSpecValueSelfContained_CliNameRef(t *testing.T) {
	recipe := &mockRecipe{}
	spec := newSelfContained(recipe).(*specValueSelfContained)

	// Test different media types
	err := qt_control.WithControl(func(c app_control.Control) error {
		// Test MediaRepository
		refRepo := spec.CliNameRef(dc_index.MediaRepository, es_lang.English, "docs")
		if refRepo == nil {
			t.Error("Expected non-nil reference for MediaRepository")
		}

		// Test MediaWeb
		refWeb := spec.CliNameRef(dc_index.MediaWeb, es_lang.English, "")
		if refWeb == nil {
			t.Error("Expected non-nil reference for MediaWeb")
		}

		// Test MediaKnowledge
		refKnowledge := spec.CliNameRef(dc_index.MediaKnowledge, es_lang.English, "knowledge")
		if refKnowledge == nil {
			t.Error("Expected non-nil reference for MediaKnowledge")
		}

		return nil
	})
	if err != nil {
		t.Error(err)
	}
}

func TestSpecValueSelfContained_Restore(t *testing.T) {
	recipe := &mockRecipe{}
	spec := newSelfContained(recipe).(*specValueSelfContained)

	err := qt_control.WithControl(func(c app_control.Control) error {
		// Create a simple JSON for restoration
		jsonData := es_json.MustParseString(`{"Value": "restored_value"}`)
		
		restoredRecipe, err := spec.Restore(jsonData, c)
		if err != nil {
			t.Error("Restore should not error", err)
			return err
		}
		if restoredRecipe == nil {
			t.Error("Expected non-nil restored recipe")
		}
		return nil
	})
	if err != nil {
		t.Error(err)
	}
}

func TestSpecValueSelfContained_SetFlags(t *testing.T) {
	recipe := &mockRecipe{}
	spec := newSelfContained(recipe).(*specValueSelfContained)

	err := qt_control.WithControl(func(c app_control.Control) error {
		flags := flag.NewFlagSet("test", flag.ContinueOnError)
		// Should not panic
		spec.SetFlags(flags, c.UI())
		return nil
	})
	if err != nil {
		t.Error(err)
	}
}

// Add necessary imports at the top of the file
func init() {
	// Make sure we import necessary packages
	_ = dc_index.MediaRepository
	_ = es_lang.English
}