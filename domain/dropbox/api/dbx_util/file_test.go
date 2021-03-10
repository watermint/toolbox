package dbx_util

import (
	"github.com/watermint/toolbox/quality/infra/qt_endtoend"
	"io/ioutil"
	"testing"
)

func TestContentHash(t *testing.T) {
	if qt_endtoend.IsSkipEndToEndTest() {
		t.Skipped()
		return
	}
	tmpf, err := ioutil.TempFile("", "hash_test")
	if err != nil {
		t.Error("Unable to create temp file", err)
	}

	h, err := FileContentHash(tmpf.Name())
	if err != nil {
		t.Error("Unable to read file", err)
	}
	if h != contentHashZeroHash {
		t.Errorf("Hash not matched expected[%s] actual[%s]", contentHashZeroHash, h)
	}

	if _, err = tmpf.WriteString("Hello"); err != nil {
		t.Error("Unabble to write content", err)
	}
	hashHello := "70bc18bef5ae66b72d1995f8db90a583a60d77b4066e4653f1cead613025861c"
	h, err = FileContentHash(tmpf.Name())
	if err != nil {
		t.Error("Unable to read file", err)
	}
	if h != hashHello {
		t.Errorf("Hash not matched expected[%s] actual[%s]", hashHello, h)
	}
	expectedChunked := "4fdd596c2141f63b577aa7be1e73bb78fbd3966a1d8b46d56650e26b07acea16"
	for i := 0; i < contentHashBlockSize/2; i++ {
		_, err = tmpf.WriteString("COMPARE")
	}
	h, err = FileContentHash(tmpf.Name())
	if err != nil {
		t.Error("Unable to read file", err)
	}
	if h != expectedChunked {
		t.Errorf("Hash not matched expected[%s] actual[%s]", expectedChunked, h)
	}
}
