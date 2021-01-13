package em_file

import (
	"time"
)

const (
	FileNode   NodeType = "file"
	FolderNode NodeType = "folder"
)

type NodeType string

type Node interface {
	// Name of the node.
	Name() string

	// Node type
	Type() NodeType

	// Compare nodes.
	Equals(x Node) bool

	// Extra data for serialize node state
	ExtraData() map[string]interface{}

	// Rename this node
	Rename(newName string)
}

type File interface {
	Node

	// Size of the file.
	Size() int64

	// Modified time.
	ModTime() time.Time

	// Content of the file.
	Content() []byte

	// Update content
	UpdateContent(newSeed, newSize int64)

	// Update mod time
	UpdateTime(t time.Time)

	// Clone this instance
	Clone() File
}

type Folder interface {
	Node

	// Returns descendant nodes.
	Descendants() []Node

	// Add node under this folder.
	// This will replace existing node by name.
	Add(node Node)

	// Delete node by name. Returns false if the node not found.
	Delete(name string) bool

	// Compare nodes and descendants.
	DeepEquals(x Node) bool

	// Number of files. Does not include files under descendants
	NumFiles() int

	// Number of folders. Does not include folders under descendants
	NumFolders() int
}
