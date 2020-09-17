package es_filecompare

import (
	"github.com/google/go-cmp/cmp"
	"github.com/watermint/toolbox/essentials/encoding/es_json"
	"github.com/watermint/toolbox/essentials/file/es_filesystem_model"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/model/em_file"
	"github.com/watermint/toolbox/essentials/queue/eq_sequence"
	"testing"
	"time"
)

func TestFolderComparator_Compare(t *testing.T) {
	seq := eq_sequence.New()

	{
		srcRoot := em_file.DemoTree()
		srcFs := es_filesystem_model.NewFileSystem(srcRoot)
		tgtRoot := em_file.DemoTree()
		tgtFs := es_filesystem_model.NewFileSystem(tgtRoot)

		fc := NewFolderComparator(srcFs, tgtFs, seq)

		missingSources, missingTargets, fileDiffs, typeDiffs, err := fc.CompareAndSummarize(es_filesystem_model.NewPath("/"), es_filesystem_model.NewPath("/"))
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
			for _, d := range fileDiffs {
				println("Diff", esl.String("diff", cmp.Diff(d.SourceData, d.TargetData)))
			}
		}
	}

	{
		srcRoot := em_file.DemoTree()
		tgtRoot := em_file.DemoTree()

		srcRoot.Add(em_file.NewFolder("q", []em_file.Node{}))
		srcRoot.Add(em_file.NewFolder("r", []em_file.Node{}))
		tgtRoot.Add(em_file.NewFile("r", 55, time.Now(), 55))
		em_file.ResolvePath(tgtRoot, "/a/b").(em_file.Folder).Add(em_file.NewFile(
			"v",
			83,
			time.Now(),
			83,
		))
		em_file.ResolvePath(srcRoot, "/a").(em_file.Folder).Add(em_file.NewFile(
			"x",
			83,
			time.Now(),
			83,
		))

		srcFs := es_filesystem_model.NewFileSystem(srcRoot)
		tgtFs := es_filesystem_model.NewFileSystem(tgtRoot)
		fc := NewFolderComparator(srcFs, tgtFs, seq)

		missingSources, missingTargets, fileDiffs, typeDiffs, err := fc.CompareAndSummarize(es_filesystem_model.NewPath("/"), es_filesystem_model.NewPath("/"))
		if err != nil {
			t.Error(err)
		}
		if len(missingSources) != 1 {
			t.Error(missingSources)
		}
		if len(missingTargets) != 1 {
			t.Error(missingTargets)
		}
		if len(typeDiffs) != 1 {
			t.Error(typeDiffs)
		}
		if len(fileDiffs) != 1 {
			t.Error(es_json.ToJsonString(fileDiffs))
		}
	}

}
