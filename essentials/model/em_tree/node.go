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
	Name() string
	Type() NodeType
}

type File interface {
	Node
	Size() int64
	MTime() time.Time
}

type Folder interface {
	Node
	Descendants() []Node
}
