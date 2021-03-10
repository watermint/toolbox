package es_filesystem_copier

import (
	"github.com/watermint/toolbox/essentials/file/es_filesystem"
	"github.com/watermint/toolbox/essentials/file/es_filesystem_model"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/model/em_file"
	"testing"
)

func TestModelToModelConn_Copy(t *testing.T) {
	tree1 := em_file.DemoTree()
	tree2 := em_file.NewFolder("", []em_file.Node{})

	con := NewModelToModel(esl.Default(), tree1, tree2)

	z := em_file.ResolvePath(tree1, "/a/c/z")
	ze := es_filesystem_model.NewEntry("/a/c/z", z)

	con.Copy(
		ze,
		es_filesystem_model.NewPath("/a/c/z"),
		func(pair es_filesystem.CopyPair, copied es_filesystem.Entry) {
		},
		func(pair es_filesystem.CopyPair, err es_filesystem.FileSystemError) {
			t.Error(err)
		},
	)

	z2 := em_file.ResolvePath(tree2, "/a/c/z")
	if !z.Equals(z2) {
		t.Error(z2)
	}
}
