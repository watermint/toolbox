package es_filesystem_copier

import (
	"github.com/watermint/toolbox/essentials/file/es_filecompare"
	"github.com/watermint/toolbox/essentials/file/es_filesystem_local"
	"github.com/watermint/toolbox/essentials/file/es_sync"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/queue/eq_sequence"
	"github.com/watermint/toolbox/quality/infra/qt_file"
	"testing"
)

func TestLocalToLocalCopier_Copy(t *testing.T) {
	qt_file.TestWithTestFolder(t, "l2l_source", true, func(sourcePath string) {
		qt_file.TestWithTestFolder(t, "l2l_target", false, func(targetPath string) {
			fs1 := es_filesystem_local.NewFileSystem()
			fs2 := es_filesystem_local.NewFileSystem()
			conn := NewLocalToLocal(esl.Default(), fs1, fs2)

			seq := eq_sequence.New()

			sync := es_sync.New(esl.Default(), seq, fs1, fs2, conn)

			errSync := sync.Sync(es_filesystem_local.NewPath(sourcePath), es_filesystem_local.NewPath(targetPath))
			if errSync != nil {
				t.Error(errSync)
			}

			comparator := es_filecompare.NewFolderComparator(
				fs1,
				fs2,
				seq,
			)
			missingSrc, missingTgt, fileDiffs, typeDiffs, cmpErr := comparator.CompareAndSummarize(
				es_filesystem_local.NewPath(sourcePath),
				es_filesystem_local.NewPath(targetPath),
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
				t.Error(fileDiffs)
			}
			if len(typeDiffs) > 0 {
				t.Error(typeDiffs)
			}
		})
	})
}
