package catalogue

import (
	"github.com/watermint/toolbox/essentials/go/es_module"
	"testing"
)

func TestScan(t *testing.T) {
	_ = NewCatalogue()

	_, err := es_module.ScanBuild()
	if err != nil {
		t.Error(err)
	}
	//if b.GoVersion() == "" {
	//	t.Error("empty go version")
	//}
	//if len(b.Modules()) < 1 {
	//	t.Error("no modules found")
	//}
}
