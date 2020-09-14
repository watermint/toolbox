package es_filesystem_connector

import (
	"github.com/watermint/toolbox/essentials/encoding/es_json"
	"github.com/watermint/toolbox/essentials/file/es_filecompare"
	"github.com/watermint/toolbox/essentials/file/es_filesystem_local"
	"github.com/watermint/toolbox/essentials/file/es_filesystem_model"
	"github.com/watermint/toolbox/essentials/file/es_sync"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/model/em_tree"
	"github.com/watermint/toolbox/essentials/queue/eq_sequence"
	"github.com/watermint/toolbox/quality/infra/qt_file"
	"testing"
)

func TestModelToLocalConn_Copy(t *testing.T) {
	qt_file.TestWithTestFolder(t, "m2l", false, func(path string) {
		tree1 := em_tree.DemoTree()
		fs1 := es_filesystem_model.NewFileSystem(tree1)
		fs2 := es_filesystem_local.NewFileSystem()
		conn := NewModelToLocal(esl.Default(), tree1)

		seq := eq_sequence.New()

		sync := es_sync.New(esl.Default(), seq, fs1, fs2, conn)

		sourcePath := es_filesystem_model.NewPath("/")
		targetPath := es_filesystem_local.NewPath(path)

		errSync := sync.Sync(sourcePath, targetPath)
		if errSync != nil {
			t.Error(errSync)
		}

		comparator := es_filecompare.NewFolderComparator(
			fs1,
			fs2,
			seq,
		)
		missingSrc, missingTgt, fileDiffs, typeDiffs, cmpErr := comparator.CompareAndSummarize(
			sourcePath,
			targetPath,
		)
		if cmpErr != nil {
			t.Error(cmpErr)
		}
		if len(missingSrc) > 0 {
			t.Error(missingSrc)
		}
		if len(missingTgt) > 0 {
			t.Error(missingTgt)
		}
		if len(fileDiffs) > 0 {
			t.Error(es_json.ToJsonString(fileDiffs))
		}
		if len(typeDiffs) > 0 {
			t.Error(typeDiffs)
		}

	})
}
