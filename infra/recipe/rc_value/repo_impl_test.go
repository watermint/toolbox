package rc_value

import (
	"flag"
	"github.com/watermint/toolbox/essentials/model/mo_string"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/quality/infra/qt_control"
	"reflect"
	"testing"
)

type TestRecipe struct {
	StringValue    string
	OptionalString mo_string.OptionalString
	IntValue       int
	BoolValue      bool
}

func (z *TestRecipe) Preset() {
	z.StringValue = "default"
	z.IntValue = 42
	z.BoolValue = true
}

func (z *TestRecipe) Exec(c app_control.Control) error {
	return nil
}

func (z *TestRecipe) Test(c app_control.Control) error {
	return nil
}

type InvalidRecipe int

func TestNewRepository(t *testing.T) {
	// Test valid recipe
	recipe := &TestRecipe{}
	repo := NewRepository(recipe)
	if repo == nil {
		t.Error("Expected repository to be created for valid recipe")
	}

	// Test repository implementation
	repoImpl := repo.(*RepositoryImpl)
	if repoImpl.rcp == nil {
		t.Error("Expected recipe to be set")
	}

	expectedName := "github.com/watermint/toolbox/infra/recipe/rc_value.test_recipe"
	if repoImpl.rcpName != expectedName {
		t.Errorf("Expected recipe name %s, got %s", expectedName, repoImpl.rcpName)
	}

	// Verify values were initialized
	if len(repoImpl.values) == 0 {
		t.Error("Expected values to be initialized")
	}

	// Test invalid recipe (non-struct)
	invalid := InvalidRecipe(1)
	invalidRepo := NewRepository(&invalid)
	if invalidRepo != nil {
		t.Error("Expected nil repository for invalid recipe")
	}
}

func TestRepositoryImpl_Current(t *testing.T) {
	recipe := &TestRecipe{}
	repo := NewRepository(recipe)
	repoImpl := repo.(*RepositoryImpl)

	current := repoImpl.Current()
	if current == nil {
		t.Error("Expected current recipe to be returned")
	}

	currentRecipe := current.(*TestRecipe)
	if currentRecipe.StringValue != "default" {
		t.Errorf("Expected default string value, got %s", currentRecipe.StringValue)
	}
}

func TestRepositoryImpl_FieldValue(t *testing.T) {
	recipe := &TestRecipe{}
	repo := NewRepository(recipe)
	repoImpl := repo.(*RepositoryImpl)

	// Test existing field
	stringValue := repoImpl.FieldValue("StringValue")
	if stringValue == nil {
		t.Error("Expected field value to be found")
	}

	// Test non-existing field
	nonExisting := repoImpl.FieldValue("NonExistingField")
	if nonExisting != nil {
		t.Error("Expected nil for non-existing field")
	}
}

func TestRepositoryImpl_FieldNames(t *testing.T) {
	recipe := &TestRecipe{}
	repo := NewRepository(recipe)
	repoImpl := repo.(*RepositoryImpl)

	fieldNames := repoImpl.FieldNames()
	if len(fieldNames) == 0 {
		t.Error("Expected field names to be returned")
	}

	// Field names should be sorted
	for i := 1; i < len(fieldNames); i++ {
		if fieldNames[i-1] > fieldNames[i] {
			t.Error("Expected field names to be sorted")
			break
		}
	}
}

func TestRepositoryImpl_FieldValueText(t *testing.T) {
	recipe := &TestRecipe{}
	repo := NewRepository(recipe)
	repoImpl := repo.(*RepositoryImpl)

	// Test existing field
	textValue := repoImpl.FieldValueText("StringValue")
	if textValue == "" {
		t.Error("Expected field value text to be returned")
	}

	// Test non-existing field
	nonExistingText := repoImpl.FieldValueText("NonExistingField")
	if nonExistingText != "" {
		t.Error("Expected empty string for non-existing field")
	}
}

func TestRepositoryImpl_Messages(t *testing.T) {
	recipe := &TestRecipe{}
	repo := NewRepository(recipe)
	repoImpl := repo.(*RepositoryImpl)

	messages := repoImpl.Messages()
	// Messages can be empty depending on recipe structure
	if messages == nil {
		t.Error("Expected messages slice to be initialized")
	}
}

func TestRepositoryImpl_Conns(t *testing.T) {
	recipe := &TestRecipe{}
	repo := NewRepository(recipe)
	repoImpl := repo.(*RepositoryImpl)

	conns := repoImpl.Conns()
	if conns == nil {
		t.Error("Expected connections map to be initialized")
	}
}

func TestRepositoryImpl_GridDataInputSpecs(t *testing.T) {
	recipe := &TestRecipe{}
	repo := NewRepository(recipe)
	repoImpl := repo.(*RepositoryImpl)

	specs := repoImpl.GridDataInputSpecs()
	if specs == nil {
		t.Error("Expected grid data input specs map to be initialized")
	}
}

func TestRepositoryImpl_GridDataOutputSpecs(t *testing.T) {
	recipe := &TestRecipe{}
	repo := NewRepository(recipe)
	repoImpl := repo.(*RepositoryImpl)

	specs := repoImpl.GridDataOutputSpecs()
	if specs == nil {
		t.Error("Expected grid data output specs map to be initialized")
	}
}

func TestRepositoryImpl_TextInputSpecs(t *testing.T) {
	recipe := &TestRecipe{}
	repo := NewRepository(recipe)
	repoImpl := repo.(*RepositoryImpl)

	specs := repoImpl.TextInputSpecs()
	if specs == nil {
		t.Error("Expected text input specs map to be initialized")
	}
}

func TestRepositoryImpl_ApplyCustom(t *testing.T) {
	recipe := &TestRecipe{}
	repo := NewRepository(recipe)
	repoImpl := repo.(*RepositoryImpl)

	// ApplyCustom should not panic
	repoImpl.ApplyCustom()
}

func TestValueOfType(t *testing.T) {
	recipe := &TestRecipe{}
	stringType := reflect.TypeOf("")
	
	// Test with valid type
	value := valueOfType(recipe, stringType, recipe, "test")
	if value == nil {
		t.Error("Expected value to be found for string type")
	}

	// Test with invalid type
	invalidType := reflect.TypeOf(complex(1, 2))
	invalidValue := valueOfType(recipe, invalidType, recipe, "test")
	if invalidValue != nil {
		t.Error("Expected nil for unsupported type")
	}
}

func TestRepositoryImpl_WithFlags(t *testing.T) {
	err := qt_control.WithControl(func(c app_control.Control) error {
		recipe := &TestRecipe{}
		repo := NewRepository(recipe)

		// Test applying flags
		flg := flag.NewFlagSet("test", flag.ContinueOnError)
		repo.ApplyFlags(flg, c.UI())

		// Parse some flags
		if err := flg.Parse([]string{"-string-value", "modified"}); err != nil {
			t.Error(err)
			return err
		}

		// Apply parsed values
		applied := repo.Apply()
		appliedRecipe := applied.(*TestRecipe)
		if appliedRecipe.StringValue != "modified" {
			t.Errorf("Expected modified string value, got %s", appliedRecipe.StringValue)
		}

		return nil
	})
	if err != nil {
		t.Error(err)
	}
}

func TestRepositoryImpl_SpinUpDown(t *testing.T) {
	err := qt_control.WithControl(func(c app_control.Control) error {
		recipe := &TestRecipe{}
		repo := NewRepository(recipe)

		// Test spin up
		spunUp, err := repo.SpinUp(c)
		if err != nil {
			t.Error(err)
			return err
		}
		if spunUp == nil {
			t.Error("Expected spun up recipe to be returned")
		}

		// Test spin down
		if err := repo.SpinDown(c); err != nil {
			t.Error(err)
			return err
		}

		return nil
	})
	if err != nil {
		t.Error(err)
	}
}

func TestErrorConstants(t *testing.T) {
	if ErrorMissingRequiredOption == nil {
		t.Error("ErrorMissingRequiredOption should be defined")
	}

	if ErrorInvalidValue == nil {
		t.Error("ErrorInvalidValue should be defined")
	}

	if ErrorMissingRequiredOption.Error() == "" {
		t.Error("ErrorMissingRequiredOption should have error message")
	}

	if ErrorInvalidValue.Error() == "" {
		t.Error("ErrorInvalidValue should have error message")
	}
}

func TestValueTypes(t *testing.T) {
	if len(ValueTypes) == 0 {
		t.Error("ValueTypes should not be empty")
	}

	// Test that all value types are valid
	for i, vt := range ValueTypes {
		if vt == nil {
			t.Errorf("ValueTypes[%d] should not be nil", i)
		}
	}
}

type PresetRecipe struct {
	Value string
}

func (z *PresetRecipe) Preset() {
	z.Value = "preset_value"
}

func (z *PresetRecipe) Exec(c app_control.Control) error {
	return nil
}

func (z *PresetRecipe) Test(c app_control.Control) error {
	return nil
}

func TestRepositoryImpl_PresetCalled(t *testing.T) {
	recipe := &PresetRecipe{}
	repo := NewRepository(recipe)
	repoImpl := repo.(*RepositoryImpl)
	
	current := repoImpl.Current().(*PresetRecipe)
	if current.Value != "preset_value" {
		t.Errorf("Expected Preset() to be called, value should be 'preset_value', got %s", current.Value)
	}
}

func TestRepositoryImpl_Feeds(t *testing.T) {
	recipe := &TestRecipe{}
	repo := NewRepository(recipe)
	repoImpl := repo.(*RepositoryImpl)

	feeds := repoImpl.Feeds()
	if feeds == nil {
		t.Error("Expected feeds map to be initialized")
	}
}

func TestRepositoryImpl_FeedSpecs(t *testing.T) {
	recipe := &TestRecipe{}
	repo := NewRepository(recipe)
	repoImpl := repo.(*RepositoryImpl)

	feedSpecs := repoImpl.FeedSpecs()
	if feedSpecs == nil {
		t.Error("Expected feed specs map to be initialized")
	}
}

func TestRepositoryImpl_Reports(t *testing.T) {
	recipe := &TestRecipe{}
	repo := NewRepository(recipe)
	repoImpl := repo.(*RepositoryImpl)

	reports := repoImpl.Reports()
	if reports == nil {
		t.Error("Expected reports map to be initialized")
	}
}

func TestRepositoryImpl_ReportSpecs(t *testing.T) {
	recipe := &TestRecipe{}
	repo := NewRepository(recipe)
	repoImpl := repo.(*RepositoryImpl)

	reportSpecs := repoImpl.ReportSpecs()
	if reportSpecs == nil {
		t.Error("Expected report specs map to be initialized")
	}
}

func TestRepositoryImpl_JsonInputSpecs(t *testing.T) {
	recipe := &TestRecipe{}
	repo := NewRepository(recipe)
	repoImpl := repo.(*RepositoryImpl)

	specs := repoImpl.JsonInputSpecs()
	if specs == nil {
		t.Error("Expected JSON input specs map to be initialized")
	}
}

func TestRepositoryImpl_FieldDesc(t *testing.T) {
	recipe := &TestRecipe{}
	repo := NewRepository(recipe)
	repoImpl := repo.(*RepositoryImpl)

	desc := repoImpl.FieldDesc("StringValue")
	if desc == nil {
		t.Error("Expected field description to be returned")
	}
}

func TestRepositoryImpl_FieldCustomDefault(t *testing.T) {
	recipe := &TestRecipe{}
	repo := NewRepository(recipe)
	repoImpl := repo.(*RepositoryImpl)

	customDefault := repoImpl.FieldCustomDefault("StringValue")
	if customDefault == nil {
		t.Error("Expected field custom default to be returned")
	}
}

func TestRepositoryImpl_Debug(t *testing.T) {
	recipe := &TestRecipe{}
	repo := NewRepository(recipe)
	repoImpl := repo.(*RepositoryImpl)

	debug := repoImpl.Debug()
	if debug == nil {
		t.Error("Expected debug map to be initialized")
	}
}

func TestRepositoryImpl_CaptureRestore(t *testing.T) {
	err := qt_control.WithControl(func(c app_control.Control) error {
		recipe := &TestRecipe{}
		repo := NewRepository(recipe)
		repoImpl := repo.(*RepositoryImpl)

		// Test capture
		captured, err := repoImpl.Capture(c)
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