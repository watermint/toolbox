package fd_file_impl

import (
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/quality/infra/qt_control"
	"testing"
)

func TestNewSpec(t *testing.T) {
	// Test with valid model
	rf := NewRowFeed("test")
	rf.SetModel(&TestModel{})
	rowFeed := rf.(*RowFeed)
	rowFeed.applyModel()
	
	spec := newSpec(rowFeed)
	if spec == nil {
		t.Error("Expected non-nil spec")
	}
	
	specImpl := spec.(*Spec)
	if specImpl.rf != rowFeed {
		t.Error("Expected rf to be set")
	}
	if specImpl.base == "" {
		t.Error("Expected base to be set")
	}
	if len(specImpl.colDesc) != 4 {
		t.Errorf("Expected 4 column descriptions, got %d", len(specImpl.colDesc))
	}
	if len(specImpl.colExample) != 4 {
		t.Errorf("Expected 4 column examples, got %d", len(specImpl.colExample))
	}
	
	// Test panic with nil model
	rf2 := NewRowFeed("test")
	rowFeed2 := rf2.(*RowFeed)
	
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic with nil model")
		}
	}()
	
	newSpec(rowFeed2)
}

func TestSpec_Name(t *testing.T) {
	rf := NewRowFeed("test_feed")
	rf.SetModel(&TestModel{})
	spec := rf.Spec()
	
	if spec.Name() != "test_feed" {
		t.Errorf("Expected name 'test_feed', got '%s'", spec.Name())
	}
}

func TestSpec_Desc(t *testing.T) {
	rf := NewRowFeed("test")
	rf.SetModel(&TestModel{})
	spec := rf.Spec()
	
	desc := spec.Desc()
	if desc == nil {
		t.Error("Expected non-nil description message")
	}
}

func TestSpec_Columns(t *testing.T) {
	rf := NewRowFeed("test")
	rf.SetModel(&TestModel{})
	spec := rf.Spec()
	
	cols := spec.Columns()
	if len(cols) != 4 {
		t.Errorf("Expected 4 columns, got %d", len(cols))
	}
	
	expectedCols := []string{"name", "age", "active", "country"}
	for i, col := range expectedCols {
		if cols[i] != col {
			t.Errorf("Expected column %s at index %d, got %s", col, i, cols[i])
		}
	}
}

func TestSpec_ColumnDesc(t *testing.T) {
	rf := NewRowFeed("test")
	rf.SetModel(&TestModel{})
	spec := rf.Spec()
	
	// Test existing column
	nameDesc := spec.ColumnDesc("name")
	if nameDesc == nil {
		t.Error("Expected non-nil description for 'name' column")
	}
	
	// Test non-existing column
	invalidDesc := spec.ColumnDesc("invalid")
	if invalidDesc != nil {
		t.Error("Expected nil for invalid column")
	}
}

func TestSpec_ColumnExample(t *testing.T) {
	rf := NewRowFeed("test")
	rf.SetModel(&TestModel{})
	spec := rf.Spec()
	
	// Test existing column
	nameExample := spec.ColumnExample("name")
	if nameExample == nil {
		t.Error("Expected non-nil example for 'name' column")
	}
	
	// Test non-existing column
	invalidExample := spec.ColumnExample("invalid")
	if invalidExample != nil {
		t.Error("Expected nil for invalid column")
	}
}

func TestSpec_Doc(t *testing.T) {
	rf := NewRowFeed("test")
	rf.SetModel(&TestModel{})
	spec := rf.Spec()
	
	err := qt_control.WithControl(func(c app_control.Control) error {
		doc := spec.Doc(c.UI())
		
		if doc == nil {
			t.Error("Expected non-nil doc")
		}
		if doc.Name != "test" {
			t.Errorf("Expected doc name 'test', got '%s'", doc.Name)
		}
		// Description might be empty as it uses TextOrEmpty - skip this check
		if len(doc.Columns) != 4 {
			t.Errorf("Expected 4 columns in doc, got %d", len(doc.Columns))
		}
		
		// Check first column
		if doc.Columns[0].Name != "name" {
			t.Errorf("Expected first column name 'name', got '%s'", doc.Columns[0].Name)
		}
		
		return nil
	})
	
	if err != nil {
		t.Error(err)
	}
}

func TestSpec_EmptyModel(t *testing.T) {
	rf := NewRowFeed("test")
	rf.SetModel(&TestModelEmpty{})
	spec := rf.Spec()
	
	cols := spec.Columns()
	if len(cols) != 0 {
		t.Error("Expected no columns for empty model")
	}
	
	err := qt_control.WithControl(func(c app_control.Control) error {
		doc := spec.Doc(c.UI())
		
		if len(doc.Columns) != 0 {
			t.Error("Expected no columns in doc for empty model")
		}
		
		return nil
	})
	
	if err != nil {
		t.Error(err)
	}
}

func TestSpec_AllFieldTypes(t *testing.T) {
	type AllTypesModel struct {
		StringField string
		IntField    int
		BoolField   bool
	}
	
	rf := NewRowFeed("test")
	rf.SetModel(&AllTypesModel{})
	spec := rf.Spec()
	
	cols := spec.Columns()
	if len(cols) != 3 {
		t.Errorf("Expected 3 columns, got %d", len(cols))
	}
	
	// Verify all columns have descriptions and examples
	for _, col := range cols {
		if spec.ColumnDesc(col) == nil {
			t.Errorf("Expected description for column %s", col)
		}
		if spec.ColumnExample(col) == nil {
			t.Errorf("Expected example for column %s", col)
		}
	}
}