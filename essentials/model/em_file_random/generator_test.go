package em_file_random

import (
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/model/em_file"
	"math/rand"
	"testing"
	"time"
)

func TestGenImpl_Generate(t *testing.T) {
	g := NewPoissonTree()
	opts := []Opt{
		Depth(10),
		NumFiles(100),
		NumDescendant(0, 10000),
		FileSize(0, 2000),
	}
	root := g.Generate(opts...)
	computed := Default().Apply(opts)

	if x := em_file.MaxDepth(root); computed.depthRangeMax < x {
		t.Error(x)
	}
	if x := em_file.SumFileSize(root); int64(computed.numFiles)*computed.fileSizeRangeMax < x {
		t.Error(x)
	}
	if x := em_file.SumNumFiles(root); x != computed.numFiles {
		t.Error(x)
	}

	em_file.Display(esl.Default(), root)
}

func TestGenImpl_Update(t *testing.T) {
	g := NewPoissonTree()
	root := em_file.DemoTree()
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	l := esl.Default()

	em_file.Display(esl.Default(), root)

	for i := 0; i < 10; i++ {
		l.Info("Update===========", esl.Int("tries", i))
		g.Update(root, r)
		em_file.Display(esl.Default(), root)
	}
}
