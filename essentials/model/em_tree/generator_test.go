package em_tree

import (
	"github.com/watermint/toolbox/essentials/log/esl"
	"path/filepath"
	"testing"
)

func TestGenImpl_Generate(t *testing.T) {
	g := New()
	opts := []Opt{
		Depth(10),
		NumNodes(10, 10000),
		NumDescendant(0, 10000),
		FileSize(0, 2000),
	}
	root := g.Generate(opts...)
	computed := Default().Apply(opts)

	if x := MaxDepth(root); computed.depthRangeMax < x {
		t.Error(x)
	}
	if x := SumFileSize(root); int64(computed.numNodesRangeMax)*computed.fileSizeRangeMax < x {
		t.Error(x)
	}
	if x := SumNumNode(root); x < computed.numNodesRangeMin || computed.numNodesRangeMax < x {
		t.Error(x)
	}

	l := esl.Default()
	var traverse func(path string, node Node)
	traverse = func(path string, node Node) {
		p := filepath.Join(path, node.Name())
		switch n := node.(type) {
		case File:
			l.Info("File",
				esl.String("path", p),
				esl.Int64("size", n.Size()),
				esl.Time("mtime", n.MTime()),
			)
		case Folder:
			l.Info("Folder",
				esl.String("path", p),
			)
			for _, d := range n.Descendants() {
				traverse(p, d)
			}
		}
	}

	traverse("", root)
}
