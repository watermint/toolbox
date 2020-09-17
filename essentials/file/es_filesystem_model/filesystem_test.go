package es_filesystem_model

import (
	"github.com/google/go-cmp/cmp"
	"github.com/watermint/toolbox/essentials/encoding/es_json"
	"github.com/watermint/toolbox/essentials/file/es_filecompare"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/model/em_file"
	"testing"
)

func TestNewFileSystem(t *testing.T) {
	l := esl.Default()
	comparator := es_filecompare.New()
	root := em_file.DemoTree()
	fs := NewFileSystem(root)

	if err := fs.Delete(NewPath("/a/b")); err != nil {
		t.Error(err)
	}
	if err := fs.Delete(NewPath("/a/c")); err != nil {
		t.Error(err)
	}

	entries, err := fs.List(NewPath("/a"))
	if err != nil {
		t.Error(err)
	}
	if len(entries) < 1 {
		t.Error(entries)
	}

	for _, entry := range entries {
		resolvedEntry, err := fs.Info(entry.Path())
		if err != nil {
			t.Error(entry, err)
		}

		if same, err := comparator.Compare(entry, resolvedEntry); !same || err != nil {
			l.Warn("entry", esl.Any("entry", entry.AsData()))
			l.Warn("resolvedEntry", esl.Any("resolvedEntry", resolvedEntry.AsData()))
			l.Warn("diff", esl.String("diff", cmp.Diff(entry.AsData(), resolvedEntry.AsData())))
			t.Error(
				es_json.ToJsonString(entry.AsData()),
				es_json.ToJsonString(resolvedEntry.AsData()),
				same, err)
		}

		if err := fs.Delete(entry.Path()); err != nil {
			t.Error(entry, err)
		}
	}

	if err := fs.CreateFolder(NewPath("/a/hello")); err != nil {
		t.Error(err)
	}
}
