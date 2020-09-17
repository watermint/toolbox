package em_file

import (
	"github.com/watermint/toolbox/essentials/log/esl"
	"math/rand"
	"testing"
	"time"
)

func TestGenImpl_Generate(t *testing.T) {
	g := NewGenerator()
	opts := []Opt{
		Depth(10),
		NumNodes(10, 10, 100),
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

	Display(esl.Default(), root)
}

func TestGenImpl_Update(t *testing.T) {
	g := NewGenerator()
	root := DemoTree()
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	l := esl.Default()

	Display(esl.Default(), root)

	for i := 0; i < 10; i++ {
		l.Info("Update===========", esl.Int("tries", i))
		g.Update(root, r)
		Display(esl.Default(), root)
	}
}
