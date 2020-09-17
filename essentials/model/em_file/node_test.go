package em_file

import (
	"testing"
	"time"
)

func TestFileNode_Equals(t *testing.T) {
	root := DemoTree()

	// folder to folder : false
	if ResolvePath(root, "/a").Equals(ResolvePath(root, "/a/b")) {
		t.Error("invalid")
	}
	// folder to folder : true
	if !ResolvePath(root, "/a").Equals(ResolvePath(root, "/a")) {
		t.Error("invalid")
	}

	// file to file : false
	if ResolvePath(root, "/a/x").Equals(ResolvePath(root, "/a/y")) {
		t.Error("invalid")
	}
	// file to file : true
	if !ResolvePath(root, "/a/x").Equals(ResolvePath(root, "/a/x")) {
		t.Error("invalid")
	}

	// folder to file
	if ResolvePath(root, "/a").Equals(ResolvePath(root, "/a/x")) {
		t.Error("invalid")
	}
	// file to folder
	if ResolvePath(root, "/a/x").Equals(ResolvePath(root, "/a")) {
		t.Error("invalid")
	}
}

func TestFolderNode_Add(t *testing.T) {
	p := NewFile("p", 48, time.Now(), 48)
	q := NewFile("p", 48, time.Now(), 49)

	root := DemoTree()
	rootFolder := root.(Folder)
	rootFolder.Add(p)

	if x := ResolvePath(root, "/p"); !x.Equals(p) {
		t.Error(x, p)
	}

	// overwrite
	rootFolder.Add(q)

	if x := ResolvePath(root, "/p"); !x.Equals(q) {
		t.Error(x, q)
	}
}

func TestFolderNode_Delete(t *testing.T) {
	root := DemoTree()
	a := ResolvePath(root, "/a").(Folder)

	n := len(a.Descendants())
	if x := a.Delete("x"); !x {
		t.Error(x)
	}
	if x := len(a.Descendants()); x != n-1 {
		t.Error(x, n)
	}

	for _, d := range a.Descendants() {
		if d.Name() == "x" {
			t.Error("deleted entry found")
		}
	}

	// should return false on not found
	if x := a.Delete("x"); x {
		t.Error(x)
	}

	// delete all nodes
	if x := a.Delete("y"); !x {
		t.Error(x)
	}
	if x := a.Delete("b"); !x {
		t.Error(x)
	}
	if x := a.Delete("c"); !x {
		t.Error(x)
	}

	if x := len(a.Descendants()); x != 0 {
		t.Error(x)
	}

	if x := a.Delete("no_existent"); x {
		t.Error(x)
	}
}
