package em_file

import (
	"math/rand"
	"reflect"
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

func TestNewFile(t *testing.T) {
	now := time.Now()
	file := NewFile("test.txt", 1024, now, 42)
	
	if file.Name() != "test.txt" {
		t.Errorf("Expected name 'test.txt', got '%s'", file.Name())
	}
	
	if file.Size() != 1024 {
		t.Errorf("Expected size 1024, got %d", file.Size())
	}
	
	if !file.ModTime().Equal(now) {
		t.Errorf("Expected time %v, got %v", now, file.ModTime())
	}
	
	if file.Type() != FileNode {
		t.Errorf("Expected type FileNode, got %v", file.Type())
	}
}

func TestNewFolder(t *testing.T) {
	file1 := NewFile("file1.txt", 100, time.Now(), 1)
	file2 := NewFile("file2.txt", 200, time.Now(), 2)
	children := []Node{file1, file2}
	
	folder := NewFolder("testfolder", children)
	
	if folder.Name() != "testfolder" {
		t.Errorf("Expected name 'testfolder', got '%s'", folder.Name())
	}
	
	if folder.Type() != FolderNode {
		t.Errorf("Expected type FolderNode, got %v", folder.Type())
	}
	
	descendants := folder.Descendants()
	if len(descendants) != 2 {
		t.Errorf("Expected 2 descendants, got %d", len(descendants))
	}
}

func TestFileNode_Clone(t *testing.T) {
	original := NewFile("original.txt", 512, time.Now(), 123)
	cloned := original.Clone()
	
	if !cloned.Equals(original) {
		t.Error("Cloned file should equal original")
	}
	
	// Verify they are separate instances
	original.Rename("renamed.txt")
	if cloned.Name() == "renamed.txt" {
		t.Error("Cloning should create separate instance")
	}
}

func TestFileNode_Rename(t *testing.T) {
	file := NewFile("old.txt", 100, time.Now(), 1)
	file.Rename("new.txt")
	
	if file.Name() != "new.txt" {
		t.Errorf("Expected name 'new.txt', got '%s'", file.Name())
	}
}

func TestFileNode_UpdateTime(t *testing.T) {
	file := NewFile("test.txt", 100, time.Unix(1000, 0), 1)
	newTime := time.Unix(2000, 0)
	
	file.UpdateTime(newTime)
	
	if !file.ModTime().Equal(newTime) {
		t.Errorf("Expected time %v, got %v", newTime, file.ModTime())
	}
}

func TestFileNode_UpdateContent(t *testing.T) {
	file := NewFile("test.txt", 100, time.Unix(1000, 0), 1)
	originalTime := file.ModTime()
	
	// Wait a bit to ensure time difference
	time.Sleep(time.Millisecond)
	
	file.UpdateContent(42, 200)
	
	if file.Size() != 200 {
		t.Errorf("Expected size 200, got %d", file.Size())
	}
	
	// Time should be updated
	if !file.ModTime().After(originalTime) {
		t.Error("ModTime should be updated after content update")
	}
	
	// Content should be different due to new seed
	expectedContent := make([]byte, 200)
	r := rand.New(rand.NewSource(42))
	r.Read(expectedContent)
	
	actualContent := file.Content()
	if len(actualContent) != 200 {
		t.Errorf("Expected content length 200, got %d", len(actualContent))
	}
}

func TestFileNode_ExtraData(t *testing.T) {
	file := NewFile("test.txt", 100, time.Now(), 42)
	
	extraData := file.ExtraData()
	
	if seed, ok := extraData[ExtraDataContentSeed]; !ok || seed != int64(42) {
		t.Errorf("Expected content seed 42 in extra data, got %v", seed)
	}
}

func TestFileNode_Content(t *testing.T) {
	file := NewFile("test.txt", 100, time.Now(), 42)
	
	content1 := file.Content()
	content2 := file.Content()
	
	// Should be reproducible with same seed
	if len(content1) != 100 {
		t.Errorf("Expected content length 100, got %d", len(content1))
	}
	
	if !reflect.DeepEqual(content1, content2) {
		t.Error("Content should be reproducible with same seed")
	}
	
	// Different seed should produce different content
	file2 := NewFile("test2.txt", 100, time.Now(), 43)
	content3 := file2.Content()
	
	if reflect.DeepEqual(content1, content3) {
		t.Error("Different seeds should produce different content")
	}
}

func TestFileNode_Equals_DetailedTests(t *testing.T) {
	now := time.Now()
	file1 := NewFile("test.txt", 100, now, 42)
	file2 := NewFile("test.txt", 100, now, 42)
	file3 := NewFile("different.txt", 100, now, 42)
	file4 := NewFile("test.txt", 200, now, 42)
	file5 := NewFile("test.txt", 100, now.Add(time.Hour), 42)
	file6 := NewFile("test.txt", 100, now, 43)
	
	// Same files should be equal
	if !file1.Equals(file2) {
		t.Error("Identical files should be equal")
	}
	
	// Different name
	if file1.Equals(file3) {
		t.Error("Files with different names should not be equal")
	}
	
	// Different size
	if file1.Equals(file4) {
		t.Error("Files with different sizes should not be equal")
	}
	
	// Different time
	if file1.Equals(file5) {
		t.Error("Files with different times should not be equal")
	}
	
	// Different content seed
	if file1.Equals(file6) {
		t.Error("Files with different content should not be equal")
	}
	
	// File vs folder
	folder := NewFolder("test", []Node{})
	if file1.Equals(folder) {
		t.Error("File should not equal folder")
	}
}

func TestFolderNode_Rename(t *testing.T) {
	folder := NewFolder("old", []Node{})
	folder.Rename("new")
	
	if folder.Name() != "new" {
		t.Errorf("Expected name 'new', got '%s'", folder.Name())
	}
}

func TestFolderNode_NumFiles(t *testing.T) {
	file1 := NewFile("file1.txt", 100, time.Now(), 1)
	file2 := NewFile("file2.txt", 200, time.Now(), 2)
	subfolder := NewFolder("sub", []Node{})
	
	folder := NewFolder("parent", []Node{file1, file2, subfolder})
	
	if folder.NumFiles() != 2 {
		t.Errorf("Expected 2 files, got %d", folder.NumFiles())
	}
	
	if folder.NumFolders() != 1 {
		t.Errorf("Expected 1 folder, got %d", folder.NumFolders())
	}
}

func TestFolderNode_ExtraData(t *testing.T) {
	folder := NewFolder("test", []Node{})
	
	extraData := folder.ExtraData()
	
	if len(extraData) != 0 {
		t.Errorf("Expected empty extra data for folder, got %v", extraData)
	}
}

func TestFolderNode_DeepEquals(t *testing.T) {
	// Create identical folder structures
	file1a := NewFile("file1.txt", 100, time.Unix(1000, 0), 1)
	file1b := NewFile("file1.txt", 100, time.Unix(1000, 0), 1)
	
	folder1a := NewFolder("folder1", []Node{file1a})
	folder1b := NewFolder("folder1", []Node{file1b})
	
	if !folder1a.DeepEquals(folder1b) {
		t.Error("Identical folder structures should be deep equal")
	}
	
	// Different file content
	file2 := NewFile("file1.txt", 100, time.Unix(1000, 0), 2) // different seed
	folder2 := NewFolder("folder1", []Node{file2})
	
	if folder1a.DeepEquals(folder2) {
		t.Error("Folders with different file content should not be deep equal")
	}
	
	// Different number of children
	file3 := NewFile("file2.txt", 100, time.Unix(1000, 0), 1)
	folder3 := NewFolder("folder1", []Node{file1a, file3})
	
	if folder1a.DeepEquals(folder3) {
		t.Error("Folders with different number of children should not be deep equal")
	}
	
	// Different folder name
	folder4 := NewFolder("folder2", []Node{file1a})
	if folder1a.DeepEquals(folder4) {
		t.Error("Folders with different names should not be deep equal")
	}
	
	// Folder vs file
	if folder1a.DeepEquals(file1a) {
		t.Error("Folder should not deep equal file")
	}
}

func TestFolderNode_Add_CaseInsensitive(t *testing.T) {
	folder := NewFolder("test", []Node{})
	
	file1 := NewFile("Test.txt", 100, time.Now(), 1)
	file2 := NewFile("test.txt", 200, time.Now(), 2)
	
	folder.Add(file1)
	if len(folder.Descendants()) != 1 {
		t.Errorf("Expected 1 child after first add, got %d", len(folder.Descendants()))
	}
	
	// Adding with different case should replace
	folder.Add(file2)
	if len(folder.Descendants()) != 1 {
		t.Errorf("Expected 1 child after case-insensitive replace, got %d", len(folder.Descendants()))
	}
	
	// Should have the second file
	child := folder.Descendants()[0]
	if !child.Equals(file2) {
		t.Error("Expected second file to replace first (case-insensitive)")
	}
}

func TestDemoTree(t *testing.T) {
	root := DemoTree()
	
	if root.Name() != "" {
		t.Errorf("Expected root name to be empty, got '%s'", root.Name())
	}
	
	if root.Type() != FolderNode {
		t.Errorf("Expected root type to be FolderNode, got %v", root.Type())
	}
	
	// Check structure
	aNode := ResolvePath(root, "/a")
	if aNode == nil {
		t.Error("Expected /a to exist")
	}
	
	xNode := ResolvePath(root, "/a/x")
	if xNode == nil || xNode.Type() != FileNode {
		t.Error("Expected /a/x to be a file")
	}
	
	yNode := ResolvePath(root, "/a/y")
	if yNode == nil || yNode.Type() != FileNode {
		t.Error("Expected /a/y to be a file")
	}
	
	bNode := ResolvePath(root, "/a/b")
	if bNode == nil || bNode.Type() != FolderNode {
		t.Error("Expected /a/b to be a folder")
	}
	
	cNode := ResolvePath(root, "/a/c")
	if cNode == nil || cNode.Type() != FolderNode {
		t.Error("Expected /a/c to be a folder")
	}
	
	zNode := ResolvePath(root, "/a/c/z")
	if zNode == nil || zNode.Type() != FileNode {
		t.Error("Expected /a/c/z to be a file")
	}
}
