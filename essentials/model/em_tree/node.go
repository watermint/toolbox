package em_tree

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
}

type File interface {
	Node

	// Size of the file.
	Size() int64

	// Modified time.
	ModTime() time.Time

	// Content of the file.
	Content() []byte
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
}
