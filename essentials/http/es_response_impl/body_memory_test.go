package es_response_impl

import (
	"github.com/watermint/toolbox/essentials/http/es_client"
	"github.com/watermint/toolbox/essentials/http/es_response"
	"github.com/watermint/toolbox/quality/infra/qt_file"
	"io/ioutil"
	"os"
	"testing"
)

func TestBodyMemoryImpl_Success(t *testing.T) {
	content := []byte("{}")
	tf, err := qt_file.MakeTestFile("test", string(content))
	if err != nil {
		t.Error(err)
		return
	}
	defer func() {
		os.Remove(tf)
	}()
	ctx := es_client.NewMock()
	bm := newMemoryBody(ctx, content)
	if bm.IsFile() {
		t.Error(bm.IsFile())
	}
	if bm.ContentLength() != int64(len(content)) {
		t.Error(bm.ContentLength())
	}
	if bm.File() != "" {
		t.Error(bm.File())
	}
	if _, err := bm.AsJson(); err != nil {
		t.Error(err)
	}
	if f, err := bm.AsFile(); err != nil {
		t.Error(err)
		if c2, err := ioutil.ReadFile(f); err != nil {
			t.Error(err)
			if string(content) != string(c2) {
				t.Error(c2)
			}
		}
	}
}

func TestBodyMemoryImpl_Failure(t *testing.T) {
	content := []byte("hello")
	tf, err := qt_file.MakeTestFile("test", string(content))
	if err != nil {
		t.Error(err)
		return
	}
	defer func() {
		os.Remove(tf)
	}()
	ctx := es_client.NewMock()
	bm := newMemoryBody(ctx, content)
	if _, err := bm.AsJson(); err != es_response.ErrorContentIsNotAJSON {
		t.Error(err)
	}
}
