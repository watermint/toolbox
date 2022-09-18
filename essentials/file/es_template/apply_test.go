package es_template

import (
	"github.com/watermint/toolbox/essentials/file/es_filesystem_model"
	"github.com/watermint/toolbox/essentials/model/em_file"
	"testing"
)

func TestNewApply(t *testing.T) {
	// a [folder]
	// +-- b [folder]
	// +-- c [folder]
	//     +-- d [folder]
	tier3d := em_file.NewFolder("d", []em_file.Node{})
	tier2c := em_file.NewFolder("c", []em_file.Node{tier3d})
	tier2b := em_file.NewFolder("b", []em_file.Node{})
	tier1a := em_file.NewFolder("a", []em_file.Node{tier2b, tier2c})
	tree1 := em_file.NewFolder("", []em_file.Node{tier1a})

	fs1 := es_filesystem_model.NewFileSystem(tree1)
	cp := NewCapture(fs1, CaptureOpts{})
	template, err := cp.Capture(es_filesystem_model.NewPath("/"))
	if err != nil {
		t.Error(err)
	}

	fs2 := es_filesystem_model.NewFileSystem(em_file.NewFolder("root", []em_file.Node{}))
	ap := NewApply(fs2, ApplyOpts{})
	err = ap.Apply(es_filesystem_model.NewPath("/"), template)
	if err != nil {
		t.Error(err)
	}

	entry, fsErr := fs2.Info(es_filesystem_model.NewPath("/a/c/d"))
	if fsErr != nil {
		t.Error(fsErr)
	}
	if entry.Name() != "d" {
		t.Error(entry.AsData())
	}
}
