package mo_file

import (
	"encoding/json"
	"github.com/watermint/toolbox/essentials/encoding/es_json"
	"github.com/watermint/toolbox/essentials/file/es_filepath"
	"path/filepath"
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
	X      float64 `path:"x" json:"x"`
	Y      float64 `path:"y" json:"y"`
	Width  float64 `path:"width" json:"width"`
	Height float64 `path:"height" json:"height"`
}

type Node struct {
	Raw                 json.RawMessage
	Id                  string     `path:"id" json:"id"`
	Type                string     `path:"type" json:"type"`
	Name                string     `path:"name" json:"name"`
	AbsoluteBoundingBox *Rectangle `json:"absoluteBoundingBox"`
	Children            []*Node    `json:"children"`
}

type NodeWithPath struct {
	Name []string
	Id   []string
	Node Node
}

func (z NodeWithPath) Path(nameIdSeparator string) string {
	elements := make([]string, len(z.Name))
	for i := range z.Name {
		elements[i] = es_filepath.Escape(z.Name[i]) + nameIdSeparator + es_filepath.Escape(z.Id[i])
	}
	return filepath.Join(elements...)
}

func Pages(doc Document) (ids []string, pages []Node, err error) {
	j, err := es_json.Parse(doc.Raw)
	if err != nil {
		return nil, nil, err
	}
	ids = make([]string, 0)
	pages = make([]Node, 0)
	err = j.FindArrayEach("document.children", func(e es_json.Json) error {
		page := Node{}
		if err := e.Model(&page); err != nil {
			return err
		}
		pages = append(pages, page)
		ids = append(ids, page.Id)
		return nil
	})
	return
}
