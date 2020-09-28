package es_size

import (
	"github.com/watermint/toolbox/essentials/ambient/ea_indicator"
	"github.com/watermint/toolbox/essentials/encoding/es_json"
	"github.com/watermint/toolbox/essentials/file/es_filesystem"
	"github.com/watermint/toolbox/essentials/file/es_filesystem_model"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/model/em_file"
	"github.com/watermint/toolbox/essentials/model/em_file_random"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"math/rand"
	"strings"
	"testing"
	"time"
)

func TestFold(t *testing.T) {
	earliest := time.Date(2020, time.September, 15, 01, 02, 03, 0, time.UTC)
	latest := time.Date(2020, time.September, 16, 01, 02, 03, 0, time.UTC)

	entries := []es_filesystem.Entry{
		es_filesystem.NewMockFileEntry("/a/001.txt", 7, earliest, "1111"),
		es_filesystem.NewMockFileEntry("/a/002.txt", 11, earliest.Add(10*time.Minute), "2222"),
		es_filesystem.NewMockFileEntry("/a/003.txt", 13, earliest.Add(20*time.Minute), "3333"),
		es_filesystem.NewMockFileEntry("/a/004.txt", 17, earliest.Add(30*time.Minute), "4444"),
		es_filesystem.NewMockFileEntry("/a/005.txt", 19, earliest.Add(40*time.Minute), "5555"),
		es_filesystem.NewMockFileEntry("/a/006.txt", 23, latest, "5555"),
		es_filesystem.NewMockFolderEntry("/a/b"),
		es_filesystem.NewMockFolderEntry("/a/c"),
		es_filesystem.NewMockFolderEntry("/a/d"),
	}

	// Shuffle entries
	seed := time.Now().UnixNano()
	l := esl.Default().With(esl.Int64("seed", seed))
	l.Info("Test with a seed")
	r := rand.New(rand.NewSource(seed))
	r.Shuffle(len(entries), func(i, j int) {
		entries[i], entries[j] = entries[j], entries[i]
	})
	fs := es_filesystem_model.NewFileSystem(em_file.DemoTree()) // dummy fs

	// Fold
	sum := Fold("/a", fs, entries)

	if sum.Path != "/a" {
		t.Error(es_json.ToJsonString(sum))
	}
	if sum.Size != 90 {
		t.Error(es_json.ToJsonString(sum))
	}
	if sum.NumFile != 6 {
		t.Error(es_json.ToJsonString(sum))
	}
	if sum.NumFolder != 3 {
		t.Error(es_json.ToJsonString(sum))
	}
	if !earliest.Equal(*sum.ModTimeEarliest) {
		t.Error(es_json.ToJsonString(sum))
	}
	if !latest.Equal(*sum.ModTimeLatest) {
		t.Error(es_json.ToJsonString(sum))
	}
}

func TestFoldFolderOnly(t *testing.T) {
	entries := []es_filesystem.Entry{
		es_filesystem.NewMockFolderEntry("/a/b"),
		es_filesystem.NewMockFolderEntry("/a/c"),
		es_filesystem.NewMockFolderEntry("/a/d"),
	}

	// Shuffle entries
	seed := time.Now().UnixNano()
	l := esl.Default().With(esl.Int64("seed", seed))
	l.Info("Test with a seed")
	r := rand.New(rand.NewSource(seed))
	r.Shuffle(len(entries), func(i, j int) {
		entries[i], entries[j] = entries[j], entries[i]
	})

	// Fold
	fs := es_filesystem_model.NewFileSystem(em_file.DemoTree()) // dummy fs
	sum := Fold("/a", fs, entries)

	if sum.Path != "/a" {
		t.Error(es_json.ToJsonString(sum))
	}
	if sum.Size != 0 {
		t.Error(es_json.ToJsonString(sum))
	}
	if sum.NumFile != 0 {
		t.Error(es_json.ToJsonString(sum))
	}
	if sum.NumFolder != 3 {
		t.Error(es_json.ToJsonString(sum))
	}
	if sum.ModTimeLatest != nil {
		t.Error(es_json.ToJsonString(sum))
	}
	if sum.ModTimeEarliest != nil {
		t.Error(es_json.ToJsonString(sum))
	}
}

func TestFoldEmpty(t *testing.T) {
	entries := []es_filesystem.Entry{}

	// Fold
	fs := es_filesystem_model.NewFileSystem(em_file.DemoTree()) // dummy fs
	sum := Fold("/a", fs, entries)

	if sum.Path != "/a" {
		t.Error(es_json.ToJsonString(sum))
	}
	if sum.Size != 0 {
		t.Error(es_json.ToJsonString(sum))
	}
	if sum.NumFile != 0 {
		t.Error(es_json.ToJsonString(sum))
	}
	if sum.NumFolder != 0 {
		t.Error(es_json.ToJsonString(sum))
	}
	if sum.ModTimeLatest != nil {
		t.Error(es_json.ToJsonString(sum))
	}
	if sum.ModTimeEarliest != nil {
		t.Error(es_json.ToJsonString(sum))
	}
}

func TestTraverseImpl_ScanSimple(t *testing.T) {
	ea_indicator.SuppressIndicatorForce()
	qtr_endtoend.TestWithControl(t, func(ctl app_control.Control) {
		model := em_file.DemoTree()
		fs := es_filesystem_model.NewFileSystem(model)
		factory := ctl.NewKvsFactory()
		defer func() {
			factory.Close()
		}()

		reportCount := 0
		err := ScanSingleFileSystem(ctl.Log(), ctl.Sequence(), factory, fs, es_filesystem_model.NewPath("/"), 2, func(s FolderSize) {
			if s.Depth > 2 {
				t.Error(es_json.ToJsonString(s))
			}
			ctl.Log().Info("Sum", esl.Any("sum", s))
			size := em_file.SumFileSize(em_file.ResolvePath(model, s.Path))
			if size != s.Size {
				t.Error(size, s.Size)
			}
			reportCount++
		})
		if err != nil {
			t.Error(err)
		}
		if reportCount != 4 {
			t.Error(reportCount)
		}
	})
}

func TestTraverseImpl_ScanLargeRandom(t *testing.T) {
	seed := time.Now().UnixNano()
	ea_indicator.SuppressIndicatorForce()
	qtr_endtoend.TestWithControl(t, func(ctl app_control.Control) {
		l := ctl.Log().With(esl.Int64("seed", seed))
		l.Info("Generate model with seed")
		model := em_file_random.NewPoissonTree().Generate(
			em_file_random.Seed(seed),
			em_file_random.NumFiles(100),
		)
		fs := es_filesystem_model.NewFileSystem(model)
		factory := ctl.NewKvsFactory()
		defer func() {
			factory.Close()
		}()

		err := ScanSingleFileSystem(ctl.Log(), ctl.Sequence(), factory, fs, es_filesystem_model.NewPath("/"), 2, func(s FolderSize) {
			if s.Depth > 2 {
				t.Error(es_json.ToJsonString(s))
			}
			ctl.Log().Info("Sum", esl.Any("sum", s))
			size := em_file.SumFileSize(em_file.ResolvePath(model, s.Path))
			if size != s.Size {
				t.Error(size, s.Size)
			}
			if x := len(strings.Split(s.Path, "/")) - 1; x != s.Depth && s.Path != "/" {
				t.Error(x, s.Depth)
			}
		})
		if err != nil {
			t.Error(err)
		}
	})
}
