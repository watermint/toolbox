package mo_file

import (
	"encoding/json"
	"github.com/watermint/toolbox/essentials/file/es_filepath"
	"strings"
)

type File struct {
	Raw          json.RawMessage
	Key          string `json:"key" path:"key"`
	Name         string `json:"name" path:"name"`
	ThumbnailUrl string `json:"thumbnailUrl" path:"thumbnailUrl"`
	LastModified string `json:"lastModified" path:"lastModified"`
}

type Document struct {
	Raw          json.RawMessage
	Name         string `path:"name" json:"name"`
	Role         string `path:"role" json:"role"`
	LastModified string `path:"lastModified" json:"lastModified"`
	EditorType   string `path:"editorType" json:"editorType"`
	ThumbnailUrl string `path:"thumbnailUrl" json:"thumbnailUrl"`
	Version      string `path:"version" json:"version"`
	Document     *Node  `json:"document"`
}

func (z Document) FindById(id string) (node Node, found bool) {
	for _, node := range z.Nodes() {
		if node.Id == id {
			return node, true
		}
	}
	return Node{}, false
}

func (z Document) Nodes() (nodes []Node) {
	nodes = make([]Node, 0)
	var findNode func(node Node)
	findNode = func(node Node) {
		nodes = append(nodes, node)
		for _, n := range node.Children {
			findNode(*n)
		}
	}
	findNode(*z.Document)
	return nodes
}

func (z Document) NodesWithPathByType(nodeType string) (nodes []NodeWithPath) {
	nodes = make([]NodeWithPath, 0)
	for _, n := range z.NodesWithPath() {
		if n.Node.Type == nodeType {
			nodes = append(nodes, n)
		}
	}
	return
}

func (z Document) NodesWithPath() (nodes []NodeWithPath) {
	nodes = make([]NodeWithPath, 0)
	var findNode func(name, id []string, node Node)
	findNode = func(name, id []string, node Node) {
		name = append(name, node.Name)
		id = append(id, node.Id)
		nn := make([]string, len(name))
		ni := make([]string, len(id))
		copy(nn[:], name[:])
		copy(ni[:], id[:])

		nodes = append(nodes, NodeWithPath{
			Name: name,
			Id:   id,
			Node: node,
		})
		for _, n := range node.Children {
			cn := make([]string, len(name))
			ci := make([]string, len(id))
			copy(cn[:], name[:])
			copy(ci[:], id[:])

			findNode(cn, ci, *n)
		}
	}
	findNode(
		[]string{},
		[]string{},
		*z.Document,
	)
	return nodes
}

type Rectangle struct {
	X      float64 `json:"x"`
	Y      float64 `json:"y"`
	Width  float64 `json:"width"`
	Height float64 `json:"height"`
}

type Node struct {
	Raw                 json.RawMessage
	Id                  string     `json:"id"`
	Type                string     `json:"type"`
	Name                string     `json:"name"`
	AbsoluteBoundingBox *Rectangle `json:"absoluteBoundingBox"`
	Children            []*Node    `json:"children"`
}

type NodeWithPath struct {
	Name []string `json:"name"`
	Id   []string `json:"id"`
	Node Node     `json:"node"`
}

func (z NodeWithPath) Path(nodeSeparator, nameIdSeparator string) string {
	elements := make([]string, len(z.Name))
	for i := range z.Name {
		elements[i] = es_filepath.Escape(z.Name[i]) + nameIdSeparator + es_filepath.Escape(z.Id[i])
	}
	return strings.Join(elements, nodeSeparator)
}
