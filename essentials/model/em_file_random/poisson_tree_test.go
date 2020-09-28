package em_file_random

import (
	"github.com/watermint/toolbox/essentials/model/em_file"
	"testing"
)

func TestPoissonTreeImpl_Generate(t *testing.T) {
	pt := NewPoissonTree()
	root := pt.Generate(
		NumFiles(100),
	)

	if x := em_file.SumNumFiles(root); x != 100 {
		t.Error(x)
	}
}
