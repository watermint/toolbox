package es_sync

import (
	"github.com/watermint/toolbox/essentials/encoding/es_json"
	"github.com/watermint/toolbox/essentials/file/es_filecompare"
	"github.com/watermint/toolbox/essentials/file/es_filesystem_connector"
	"github.com/watermint/toolbox/essentials/file/es_filesystem_model"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/model/em_tree"
	"github.com/watermint/toolbox/essentials/queue/eq_sequence"
	"testing"
)

func TestSyncImpl_Sync(t *testing.T) {
	tree1 := em_tree.DemoTree()
	tree2 := em_tree.NewFolder("root", []em_tree.Node{})

	fs1 := es_filesystem_model.NewFileSystem(tree1)
	fs2 := es_filesystem_model.NewFileSystem(tree2)

	seq := eq_sequence.New()
	conn := es_filesystem_connector.NewModelToModel(esl.Default(), tree1, tree2)

	syncer := New(
		esl.Default(),
		seq,
		fs1,
		fs2,
		conn,
	)
	err := syncer.Sync(es_filesystem_model.NewPath("/"), es_filesystem_model.NewPath("/"))
	if err != nil {
		t.Error(err)
	}

	folderCmp := es_filecompare.NewFolderComparator(fs1, fs2, seq)
	missingSources, missingTargets, fileDiffs, typeDiffs, err := folderCmp.CompareAndSummarize(es_filesystem_model.NewPath("/"), es_filesystem_model.NewPath("/"))
	if err != nil {
		t.Error(err)
	}
	if len(missingSources) > 0 {
		t.Error(missingSources)
	}
	if len(missingTargets) > 0 {
		t.Error(missingTargets)
	}
	if len(typeDiffs) > 0 {
		t.Error(typeDiffs)
	}
	if len(fileDiffs) > 0 {
		t.Error(es_json.ToJsonString(fileDiffs))
	}
}
