package es_filecompare

import (
	"github.com/watermint/toolbox/essentials/file/es_filesystem"
	"path/filepath"
	"testing"
	"time"
)

func NewMockError(err error) es_filesystem.FileSystemError {
	if err == nil {
		return nil
	}
	return &MockError{
		err: err,
	}
}

type MockError struct {
	err error
}

func (z MockError) IsMockError() bool {
	panic("implement me")
}

func (z MockError) Error() string {
	return z.err.Error()
}

func (z MockError) IsPathNotFound() bool {
	panic("implement me")
}

func (z MockError) IsConflict() bool {
	panic("implement me")
}

func (z MockError) IsNoPermission() bool {
	panic("implement me")
}

func (z MockError) IsInsufficientSpace() bool {
	panic("implement me")
}

func (z MockError) IsDisallowedName() bool {
	panic("implement me")
}

func (z MockError) IsInvalidEntryDataFormat() bool {
	panic("implement me")
}

func NewMockPath(path string) es_filesystem.Path {
	return &MockPath{
		path: path,
	}
}

type MockPath struct {
	path string
}

func (z MockPath) Ancestor() es_filesystem.Path {
	return &MockPath{
		path: filepath.ToSlash(filepath.Dir(z.Path())),
	}
}

func (z MockPath) IsRoot() bool {
	return z.Path() == "/"
}

func (z MockPath) AsData() es_filesystem.PathData {
	return es_filesystem.PathData{
		FileSystemType: "mock",
		EntryPath:      z.path,
		EntryShard:     z.Shard().AsData(),
		Attributes:     map[string]interface{}{},
	}
}

func (z MockPath) Base() string {
	return filepath.Base(z.path)
}

func (z MockPath) Path() string {
	return filepath.ToSlash(z.path)
}

func (z MockPath) Shard() es_filesystem.Shard {
	return es_filesystem.ShardData{
		FileSystemType: "mock",
		ShardId:        "",
		Attributes:     map[string]interface{}{},
	}
}

func (z MockPath) Descendant(pathFragment ...string) es_filesystem.Path {
	fragments := make([]string, 0)
	fragments = append(fragments, z.path)
	fragments = append(fragments, pathFragment...)
	return NewMockPath(filepath.ToSlash(filepath.Join(fragments...)))
}

func (z MockPath) Rel(path es_filesystem.Path) (string, es_filesystem.FileSystemError) {
	rel, err := filepath.Rel(z.path, path.Path())
	if err != nil {
		return "", NewMockError(err)
	}
	return rel, nil
}

type MockEntry struct {
	es_filesystem.EntryData
	Hash string
}

func (z MockEntry) Name() string {
	return z.EntryData.EntryName
}

func (z MockEntry) Path() es_filesystem.Path {
	return NewMockPath(z.EntryData.EntryPath)
}

func (z MockEntry) Size() int64 {
	return z.EntryData.EntrySize
}

func (z MockEntry) ModTime() time.Time {
	return z.EntryData.EntryModTime
}

func (z MockEntry) ContentHash() (string, es_filesystem.FileSystemError) {
	return z.Hash, nil
}

func (z MockEntry) IsFile() bool {
	return z.EntryData.EntryIsFile
}

func (z MockEntry) IsFolder() bool {
	return z.EntryData.EntryIsFolder
}

func (z MockEntry) AsData() es_filesystem.EntryData {
	return z.EntryData
}

func TestCompare_ThreeComparators(t *testing.T) {
	x := &MockEntry{
		EntryData: es_filesystem.EntryData{
			FileSystemType: "mock",
			EntryName:      "test.dat",
			EntryPath:      "/a/test.dat",
			EntrySize:      123,
			EntryModTime:   time.Unix(1500000000, 0),
			EntryIsFile:    true,
			EntryIsFolder:  false,
		},
		Hash: "ABC_0123",
	}
	entrySizeDiff := &MockEntry{
		EntryData: es_filesystem.EntryData{
			FileSystemType: "mock",
			EntryName:      "test.dat",
			EntryPath:      "/a/b/test.dat",
			EntrySize:      1230,
			EntryModTime:   time.Unix(1500000000, 0),
			EntryIsFile:    true,
			EntryIsFolder:  false,
		},
		Hash: "ABC_0123",
	}
	entryTimeDiff := &MockEntry{
		EntryData: es_filesystem.EntryData{
			FileSystemType: "mock",
			EntryName:      "test.dat",
			EntryPath:      "/a/b/test.dat",
			EntrySize:      123,
			EntryModTime:   time.Unix(1600000000, 0),
			EntryIsFile:    true,
			EntryIsFolder:  false,
		},
		Hash: "ABC_0123",
	}
	entryContentDiff := &MockEntry{
		EntryData: es_filesystem.EntryData{
			FileSystemType: "mock",
			EntryName:      "test.dat",
			EntryPath:      "/a/test.dat",
			EntrySize:      123,
			EntryModTime:   time.Unix(1500000000, 0),
			EntryIsFile:    true,
			EntryIsFolder:  false,
		},
		Hash: "XYZ_9876",
	}

	// All comparators
	{
		comparator := New()

		// same entry
		if same, err := comparator.Compare(x, x); !same || err != nil {
			t.Error(same, err)
		}

		// x -> diff
		if same, err := comparator.Compare(x, entrySizeDiff); same || err != nil {
			t.Error(same, err)
		}
		if same, err := comparator.Compare(x, entryTimeDiff); same || err != nil {
			t.Error(same, err)
		}
		if same, err := comparator.Compare(x, entryContentDiff); same || err != nil {
			t.Error(same, err)
		}

		// diff -> x
		if same, err := comparator.Compare(entrySizeDiff, x); same || err != nil {
			t.Error(same, err)
		}
		if same, err := comparator.Compare(entryTimeDiff, x); same || err != nil {
			t.Error(same, err)
		}
		if same, err := comparator.Compare(entryContentDiff, x); same || err != nil {
			t.Error(same, err)
		}
	}

	// Exclude time comparator
	{
		comparator := New(DontCompareTime(true))

		// same entry
		if same, err := comparator.Compare(x, x); !same || err != nil {
			t.Error(same, err)
		}

		// x -> diff
		if same, err := comparator.Compare(x, entrySizeDiff); same || err != nil {
			t.Error(same, err)
		}
		if same, err := comparator.Compare(x, entryTimeDiff); !same || err != nil {
			t.Error(same, err)
		}
		if same, err := comparator.Compare(x, entryContentDiff); same || err != nil {
			t.Error(same, err)
		}

		// diff -> x
		if same, err := comparator.Compare(entrySizeDiff, x); same || err != nil {
			t.Error(same, err)
		}
		if same, err := comparator.Compare(entryTimeDiff, x); !same || err != nil {
			t.Error(same, err)
		}
		if same, err := comparator.Compare(entryContentDiff, x); same || err != nil {
			t.Error(same, err)
		}
	}

	// Exclude content comparator
	{
		comparator := New(DontCompareContent(true))

		// same entry
		if same, err := comparator.Compare(x, x); !same || err != nil {
			t.Error(same, err)
		}

		// x -> diff
		if same, err := comparator.Compare(x, entrySizeDiff); same || err != nil {
			t.Error(same, err)
		}
		if same, err := comparator.Compare(x, entryTimeDiff); same || err != nil {
			t.Error(same, err)
		}
		if same, err := comparator.Compare(x, entryContentDiff); !same || err != nil {
			t.Error(same, err)
		}

		// diff -> x
		if same, err := comparator.Compare(entrySizeDiff, x); same || err != nil {
			t.Error(same, err)
		}
		if same, err := comparator.Compare(entryTimeDiff, x); same || err != nil {
			t.Error(same, err)
		}
		if same, err := comparator.Compare(entryContentDiff, x); !same || err != nil {
			t.Error(same, err)
		}
	}

}
