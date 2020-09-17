package em_file

import (
	"testing"
	"time"
)

func TestResolvePath(t *testing.T) {
	tier2cZ := NewFile("z", 98, time.Now(), 98)
	tier2c := NewFolder("c", []Node{tier2cZ})
	tier2b := NewFolder("b", []Node{})
	tier1aY := NewFile("y", 101, time.Now(), 101)
	tier1aX := NewFile("x", 123, time.Now(), 123)
	tier1a := NewFolder("a", []Node{tier1aX, tier1aY, tier2b, tier2c})
	root := NewFolder("", []Node{tier1a})

	if x := ResolvePath(root, "/"); !x.Equals(root) {
		t.Error(x)
	}
	if x := ResolvePath(root, "."); !x.Equals(root) {
		t.Error(x)
	}
	if x := ResolvePath(root, "/a"); !x.Equals(tier1a) {
		t.Error(x)
	}
	if x := ResolvePath(root, "/a/x"); !x.Equals(tier1aX) {
		t.Error(x)
	}
	if x := ResolvePath(root, "/a/y"); !x.Equals(tier1aY) {
		t.Error(x)
	}
	if x := ResolvePath(root, "/a/c"); !x.Equals(tier2c) {
		t.Error(x)
	}
	if x := ResolvePath(root, "/a/c/z"); !x.Equals(tier2cZ) {
		t.Error(x)
	}

	if x := ResolvePath(root, "/a/d"); x != nil {
		t.Error(x)
	}
}

func TestCreateFolder(t *testing.T) {
	root := NewFolder("", []Node{})
	if !CreateFolder(root, "/a/b/c") {
		t.Error("failed create folder")
	}

	c := ResolvePath(root, "/a/b/c")
	if c == nil {
		t.Error("folder not found at the path")
	}
}
