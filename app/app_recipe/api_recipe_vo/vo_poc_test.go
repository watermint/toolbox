package api_recipe_vo

import (
	"flag"
	"testing"
)

func TestMakeFlagSet(t *testing.T) {
	type FolderListVO struct {
		Recursive    bool
		NonRecursive bool
	}

	vo := &FolderListVO{
		Recursive:    false,
		NonRecursive: true,
	}
	f := flag.NewFlagSet("test", flag.ContinueOnError)

	recursiveFound := false
	nonRecursiveFound := false

	MakeFlagSet(f, vo)

	err := f.Parse([]string{"-recursive", "-non-recursive"})
	if err != nil {
		t.Error(err)
		return
	}

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
		}
	})

	if !recursiveFound {
		t.Error("recursive not found")
	}
	if !nonRecursiveFound {
		t.Error("non recursive not found")
	}
}
