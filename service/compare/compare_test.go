package compare

import (
	"io/ioutil"
	"testing"
)

func TestContentHash(t *testing.T) {
	tmpf, err := ioutil.TempFile("", "hash_test")
	if err != nil {
		t.Error("Unable to create temp file", err)
	}

	h, err := ContentHash(tmpf.Name())
	if err != nil {
		t.Error("Unable to read file", err)
	}
	if h != HASH_FOR_EMPTY {
		t.Errorf("Hash not matched expected[%s] actual[%s]", HASH_FOR_EMPTY, h)
	}

	if _, err = tmpf.WriteString("Hello"); err != nil {
		t.Error("Unabble to write content", err)
	}
	hashHello := "70bc18bef5ae66b72d1995f8db90a583a60d77b4066e4653f1cead613025861c"
	h, err = ContentHash(tmpf.Name())
	if err != nil {
		t.Error("Unable to read file", err)
	}
	if h != hashHello {
		t.Errorf("Hash not matched expected[%s] actual[%s]", hashHello, h)
	}
	expectedChunked := "4fdd596c2141f63b577aa7be1e73bb78fbd3966a1d8b46d56650e26b07acea16"
	for i := 0; i < BLOCK_SIZE/2; i++ {
		_, err = tmpf.WriteString("COMPARE")
	}
	h, err = ContentHash(tmpf.Name())
	if err != nil {
		t.Error("Unable to read file", err)
	}
	if h != expectedChunked {
		t.Errorf("Hash not matched expected[%s] actual[%s]", expectedChunked, h)
	}
}
