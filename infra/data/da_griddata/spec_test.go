package da_griddata

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/watermint/toolbox/infra/doc/dc_recipe"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_ui"
)

type dummyUI struct{ app_ui.UI }

func (d *dummyUI) Text(m app_msg.Message) string { return "desc" }

type dummyRecipe struct{}

func TestNewInputSpec(t *testing.T) {
	spec := NewInputSpec(&dummyRecipe{}, "TestInput")
	assert.Equal(t, "TestInput", spec.Name())
	desc := spec.Desc()
	assert.NotNil(t, desc)
	doc := spec.Doc(&dummyUI{})
	assert.IsType(t, &dc_recipe.DocGridDataInput{}, doc)
	assert.Equal(t, "TestInput", doc.Name)
	assert.Equal(t, "desc", doc.Desc)
}

func TestNewOutputSpec(t *testing.T) {
	spec := NewOutputSpec(&dummyRecipe{}, "TestOutput")
	assert.Equal(t, "TestOutput", spec.Name())
	desc := spec.Desc()
	assert.NotNil(t, desc)
	doc := spec.Doc(&dummyUI{})
	assert.IsType(t, &dc_recipe.DocGridDataOutput{}, doc)
	assert.Equal(t, "TestOutput", doc.Name)
	assert.Equal(t, "desc", doc.Desc)
}
