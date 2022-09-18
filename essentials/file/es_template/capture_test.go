package es_template

import (
	"encoding/json"
	"github.com/tidwall/gjson"
	"github.com/watermint/toolbox/essentials/file/es_filesystem_model"
	"github.com/watermint/toolbox/essentials/model/em_file"
	"testing"
)

func TestCapImpl_Capture(t *testing.T) {
	tree1 := em_file.DemoTree()
	fs1 := es_filesystem_model.NewFileSystem(tree1)

	c := NewCapture(fs1, CaptureOpts{})

	template, err := c.Capture(es_filesystem_model.NewPath("/"))
	if err != nil {
		t.Error(err)
	}
	templateJson, err := json.Marshal(&template)
	if err != nil {
		t.Error(err)
	}
	templateJsonObj := gjson.ParseBytes(templateJson)
	name := templateJsonObj.Get("folders.0.name")
	if name.String() != "a" {
		t.Error(name)
	}
}
