package uc_file_size

import (
	"fmt"
	"github.com/google/go-cmp/cmp"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_file"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_file_size"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/essentials/encoding/es_json"
	"reflect"
	"sort"
	"testing"
)

func TestSumImplPathsOfEntry(t *testing.T) {
	s := NewSum(3).(*sumImpl)
	{
		expected := []string{"/", "/a", "/a/b"}
		actual := s.pathsOfEntry("/a/b/c/d.dat")
		sort.Strings(expected)
		sort.Strings(actual)
		if !reflect.DeepEqual(expected, actual) {
			t.Error(cmp.Diff(expected, actual))
		}
	}
	{
		expected := []string{"/", "/a", "/a/b"}
		actual := s.pathsOfEntry("/a/b/c.dat")
		sort.Strings(expected)
		sort.Strings(actual)
		if !reflect.DeepEqual(expected, actual) {
			t.Error(cmp.Diff(expected, actual))
		}
	}
}

func TestSumImpl_Add(t *testing.T) {
	s := NewSum(3)

	s.Each(func(path mo_path.DropboxPath, size mo_file_size.Size) {
		t.Error("should not be here", path, size)
	})

	s.Eval("/", []mo_file.Entry{
		&mo_file.File{Size: 3456, EntryPathDisplay: "/x.dat", EntryTag: "file"},
		&mo_file.Folder{EntryPathDisplay: "/a", EntryTag: "folder"},
	})

	cp := 0
	s.Each(func(path mo_path.DropboxPath, size mo_file_size.Size) {
		switch path.Path() {
		case "":
			cp++
			if size.Size != 3456 || size.CountFolder != 1 || size.CountFile != 1 {
				t.Error(es_json.ToJsonString(size))
			}
		default:
			t.Error(es_json.ToJsonString(size))
		}
	})
	if cp != 1 {
		t.Error(cp)
	}

	s.Eval("/a", []mo_file.Entry{
		&mo_file.File{Size: 2345, EntryPathDisplay: "/a/y.dat", EntryTag: "file"},
		&mo_file.Folder{EntryPathDisplay: "/a/b", EntryTag: "folder"},
	})

	cp = 0
	s.Each(func(path mo_path.DropboxPath, size mo_file_size.Size) {
		switch path.Path() {
		case "/a":
			cp++
			if size.Size != 2345 || size.CountFolder != 1 || size.CountFile != 1 {
				t.Error(es_json.ToJsonString(size))
			}
		case "":
			cp++
			if size.Size != 5801 || size.CountFolder != 2 || size.CountFile != 2 {
				t.Error(es_json.ToJsonString(size))
			}
		default:
			t.Error(es_json.ToJsonString(size))
		}
	})
	if cp != 2 {
		t.Error(cp)
	}

	s.Eval("/a/b", []mo_file.Entry{
		&mo_file.File{Size: 1234, EntryPathDisplay: "/a/b/z.dat", EntryTag: "file"},
	})

	cp = 0
	s.Each(func(path mo_path.DropboxPath, size mo_file_size.Size) {
		switch path.Path() {
		case "/a/b":
			cp++
			if size.Size != 1234 || size.CountFolder != 0 || size.CountFile != 1 {
				t.Error(es_json.ToJsonString(size))
			}
		case "/a":
			cp++
			if size.Size != 3579 || size.CountFolder != 1 || size.CountFile != 2 {
				t.Error(es_json.ToJsonString(size))
			}
		case "":
			cp++
			if size.Size != 7035 || size.CountFolder != 2 || size.CountFile != 3 {
				t.Error(es_json.ToJsonString(size))
			}
		default:
			t.Error(es_json.ToJsonString(size))
		}
	})
	if cp != 3 {
		t.Error(cp)
	}

	entries := make([]mo_file.Entry, 0)
	for i := 1; i <= apiComplexityThreshold; i++ {
		entries = append(entries, &mo_file.File{
			Size:             int64(i),
			EntryPathDisplay: fmt.Sprintf("/a/b/c/w%d.dat", i),
			EntryTag:         "file",
		})
	}
	s.Eval("/a/b/c", entries)
	v := int64((1 + apiComplexityThreshold) * apiComplexityThreshold / 2)

	cp = 0
	s.Each(func(path mo_path.DropboxPath, size mo_file_size.Size) {
		switch path.Path() {
		case "/a/b":
			cp++
			if size.Size != 1234+v || size.CountFolder != 0 || size.CountFile != apiComplexityThreshold+1 {
				t.Error(es_json.ToJsonString(size))
			}
		case "/a":
			cp++
			if size.Size != 3579+v || size.CountFolder != 1 || size.CountFile != apiComplexityThreshold+2 ||
				size.ApiComplexity != apiComplexityThreshold+1+2 {
				t.Error(es_json.ToJsonString(size))
			}
		case "":
			cp++
			if size.Size != 7035+v || size.CountFolder != 2 || size.CountFile != apiComplexityThreshold+3 ||
				size.ApiComplexity != apiComplexityThreshold+2+3 {
				t.Error(es_json.ToJsonString(size))
			}
		default:
			t.Error(es_json.ToJsonString(size))
		}
	})
	if cp != 3 {
		t.Error(cp)
	}
}
