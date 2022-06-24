package go_module

import (
	"github.com/watermint/toolbox/catalogue"
	"testing"
)

func TestScan(t *testing.T) {
	_ = catalogue.NewCatalogue()

	b, err := ScanBuild()
	if err != nil {
		t.Error(err)
	}
	if b.GoVersion() == "" {
		t.Error("empty go version")
	}
	if len(b.Modules()) < 1 {
		t.Error("no modules found")
	}
}
