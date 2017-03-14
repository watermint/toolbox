package compare

import (
	"github.com/watermint/toolbox/infra"
	"io/ioutil"
	"os"
	"sync"
	"testing"
)

func TestTraverseLocalFile_Prepare(t *testing.T) {
	base := os.TempDir()
	tmpd, err := ioutil.TempDir(base, "traverse")
	if err != nil {
		t.Error(err)
	}
	opts := &infra.InfraOpts{
		WorkPath: base,
	}
	trav := Traverse{
		LocalBasePath: tmpd,
		InfraOpts:     opts,
	}
	err = trav.Prepare()
	if err != nil {
		t.Error(err)
	}
	defer trav.Close()

	println(tmpd)
	tmpf, err := ioutil.TempFile(tmpd, "localfile")
	if err != nil {
		t.Error(err)
	}
	tmpf.WriteString("Hello")
	tmpf.Close()

	trav.ScanLocal()

	listen := make(chan *LocalFileInfo)
	wg := sync.WaitGroup{}
	go trav.RetrieveLocal(listen, &wg)

	var lfi1, lfi2 *LocalFileInfo
	lfi1 = <-listen
	lfi2 = <-listen

	if lfi1 == nil {
		t.Errorf("Failed to scan or retrieve file: [%s]", tmpf.Name())
	}
	if lfi2 != nil {
		t.Errorf("Unexpected file scanned: [%s]", lfi2.Path)
	}
}
