package app_vo

import (
	"flag"
	"testing"
)

func TestMakeFlagSet(t *testing.T) {
	type FolderListVO struct {
		Recursive    bool
		NonRecursive bool
		Depth        int
		Filter       string
	}

	vo := &FolderListVO{
		Recursive:    false,
		NonRecursive: true,
		Depth:        2,
		Filter:       "",
	}
	f := flag.NewFlagSet("test", flag.ContinueOnError)

	vc := NewValueContainer(vo)
	vc.MakeFlagSet(f)

	err := f.Parse([]string{"-recursive", "-non-recursive=false", "-depth", "4", "-filter", "haystack"})
	if err != nil {
		t.Error(err)
		return
	}

	recursiveFound := false
	nonRecursiveFound := false
	depthFound := false
	filterFound := false

	f.VisitAll(func(g *flag.Flag) {
		switch g.Name {
		case "recursive":
			if g.DefValue != "false" {
				t.Error("invalid")
			}
			recursiveFound = true

		case "non-recursive":
			if g.DefValue != "true" {
				t.Error("invalid")
			}
			nonRecursiveFound = true

		case "depth":
			if g.DefValue != "2" {
				t.Error("invalid")
			}
			depthFound = true

		case "filter":
			if g.DefValue != "" {
				t.Error("invalid")
			}
			filterFound = true
		}
	})

	if !recursiveFound {
		t.Error("recursive not found")
	}
	if !nonRecursiveFound {
		t.Error("non recursive not found")
	}
	if !depthFound {
		t.Error("depth not found")
	}
	if !filterFound {
		t.Error("filter not found")
	}

	vc.Apply(vo)

	if vo.Filter != "haystack" {
		t.Error("invalid")
	}
	if vo.Depth != 4 {
		t.Error("invalid")
	}
	if !vo.Recursive {
		t.Error("invalid")
	}
	if vo.NonRecursive {
		t.Error("invalid")
	}
}
