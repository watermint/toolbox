package es_filesystem_connector

import (
	"github.com/watermint/toolbox/essentials/file/es_filesystem_model"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/model/em_tree"
	"testing"
)

func TestModelToModelConn_Copy(t *testing.T) {
	tree1 := em_tree.DemoTree()
	tree2 := em_tree.NewFolder("", []em_tree.Node{})

	con := NewModelToModel(esl.Default(), tree1, tree2)

	z := em_tree.ResolvePath(tree1, "/a/c/z")
	ze := es_filesystem_model.NewEntry("/a/c/z", z)

	err := con.Copy(ze, es_filesystem_model.NewPath("/a/c/z"))
	if err != nil {
		t.Error(err)
	}

	z2 := em_tree.ResolvePath(tree2, "/a/c/z")
	if !z.Equals(z2) {
		t.Error(z2)
	}
}
