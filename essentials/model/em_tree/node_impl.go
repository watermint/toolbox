package em_tree

import (
	"math/rand"
	"reflect"
	"strings"
	"sync"
	"time"
)

const (
	ExtraDataContentSeed = "content_seed"
)

func NewFile(name string, size int64, mtime time.Time, contentSeed int64) File {
	return &fileNode{
		name:        name,
		size:        size,
		mtime:       mtime,
		contentSeed: contentSeed,
	}
}

func NewFolder(name string, descendant []Node) Folder {
	return &folderNode{
		name:        name,
		descendants: descendant,
		folderMutex: sync.Mutex{},
	}
}

type fileNode struct {
	name        string
	size        int64
	mtime       time.Time
	contentSeed int64
}

func (z *fileNode) Rename(newName string) {
	z.name = newName
}

func (z *fileNode) UpdateContent(newSeed, newSize int64) {
	z.contentSeed = newSeed
	z.size = newSize
	z.mtime = time.Now()
}

func (z fileNode) ExtraData() map[string]interface{} {
	return map[string]interface{}{
		ExtraDataContentSeed: z.contentSeed,
	}
}

func (z fileNode) Equals(x Node) bool {
	switch n := x.(type) {
	case File:
		return z.Name() == n.Name() &&
			z.Size() == n.Size() &&
			z.ModTime().Equal(n.ModTime()) &&
			reflect.DeepEqual(z.Content(), n.Content())
	default:
		return false
	}
}

func (z fileNode) Content() []byte {
	content := make([]byte, z.size)
	r := rand.New(rand.NewSource(z.contentSeed))
	r.Read(content)
	return content
}

func (z fileNode) Name() string {
	return z.name
}

func (z fileNode) Type() NodeType {
	return FileNode
}

func (z fileNode) Size() int64 {
	return z.size
}

func (z fileNode) ModTime() time.Time {
	return z.mtime
}

type folderNode struct {
	name        string
	descendants []Node
	folderMutex sync.Mutex
}

func (z *folderNode) Rename(newName string) {
	z.name = newName
}

func (z *folderNode) ExtraData() map[string]interface{} {
	return map[string]interface{}{}
}
func (z *folderNode) Equals(x Node) bool {
	switch f := x.(type) {
	case Folder:
		return z.Name() == f.Name()

	default:
		return false
	}
}

func (z *folderNode) DeepEquals(x Node) bool {
	switch f := x.(type) {
	case Folder:
		if z.Name() != x.Name() {
			return false
		}
		sd := make(map[string]Node)
		xd := make(map[string]Node)
		for _, d := range z.Descendants() {
			sd[d.Name()] = d
		}
		for _, d := range f.Descendants() {
			xd[d.Name()] = d
		}

		if len(sd) != len(xd) {
			return false
		}

		// sd -> xd
		for n, d := range sd {
			if v, ok := xd[n]; ok {
				if !v.Equals(d) {
					return false
				}
			} else {
				return false
			}
		}

		// xd -> sd
		for n, d := range xd {
			if v, ok := sd[n]; ok {
				if !v.Equals(d) {
					return false
				}
			} else {
				return false
			}
		}
		return true

	default:
		return false
	}
}

func (z *folderNode) Add(node Node) {
	z.folderMutex.Lock()
	defer z.folderMutex.Unlock()

	for i := 0; i < len(z.descendants); i++ {
		if strings.ToLower(z.descendants[i].Name()) == strings.ToLower(node.Name()) {
			z.descendants[i] = node
			return
		}
	}
	z.descendants = append(z.descendants, node)
}

func (z *folderNode) Delete(name string) bool {
	z.folderMutex.Lock()
	defer z.folderMutex.Unlock()

	for i := 0; i < len(z.descendants); i++ {
		if z.descendants[i].Name() == name {
			newDescendants := make([]Node, 0)
			newDescendants = append(newDescendants, z.descendants[:i]...)
			newDescendants = append(newDescendants, z.descendants[i+1:]...)
			z.descendants = newDescendants
			return true
		}
	}
	return false
}

func (z *folderNode) Name() string {
	return z.name
}

func (z *folderNode) Type() NodeType {
	return FolderNode
}

func (z *folderNode) Descendants() []Node {
	return z.descendants
}

// Returns demo tree like below
// ```
// root
//  |-- a (folder)
//      |-- x (file)
//      |-- y (file)
//      |-- b (empty folder)
//      |-- c (folder)
//          |-- z (file)
// ```
func DemoTree() Folder {
	tier2cZ := NewFile("z", 98, time.Unix(1599540262, 0), 98)
	tier2c := NewFolder("c", []Node{tier2cZ})
	tier2b := NewFolder("b", []Node{})
	tier1aY := NewFile("y", 101, time.Unix(1599540000, 0), 101)
	tier1aX := NewFile("x", 123, time.Unix(1599500000, 0), 123)
	tier1a := NewFolder("a", []Node{tier1aX, tier1aY, tier2b, tier2c})
	return NewFolder("", []Node{tier1a})
}
