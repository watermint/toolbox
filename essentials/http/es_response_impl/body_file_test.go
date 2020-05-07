package es_response_impl

import (
	"github.com/watermint/toolbox/essentials/http/es_context"
	"github.com/watermint/toolbox/essentials/http/es_response"
	"github.com/watermint/toolbox/quality/infra/qt_file"
	"os"
	"testing"
)

func TestBodyFileImpl(t *testing.T) {
	content := "{}"
	tf, err := qt_file.MakeTestFile("test", content)
	if err != nil {
		t.Error(err)
		return
	}
	defer func() {
		os.Remove(tf)
	}()
	ctx := es_context.NewMock()
	bf := newFileBody(ctx, tf, int64(len(content)))
	if bf.File() != tf {
		t.Error(bf.File())
	}
	if !bf.IsFile() {
		t.Error(bf.IsFile())
	}
	if len(bf.Body()) != 0 {
		t.Error(bf.Body())
	}
	if bf.ContentLength() != int64(len(content)) {
		t.Error(bf.ContentLength())
	}
	if tf1, err := bf.AsFile(); tf1 != tf || err != nil {
		t.Error(tf1, err)
	}
	if _, err := bf.AsJson(); err != nil {
		t.Error(err)
	}
}

func TestBodyFileImplFailure(t *testing.T) {
	content := "Hello, World."
	tf, err := qt_file.MakeTestFile("test", content)
	if err != nil {
		t.Error(err)
		return
	}
	defer func() {
		os.Remove(tf)
	}()
	ctx := es_context.NewMock()

	// file too large
	{
		bf := newFileBody(ctx, tf, es_response.MaximumJsonSize+1)
		if _, err := bf.AsJson(); err != es_response.ErrorContentIsTooLarge {
			t.Error(err)
		}
	}

	// invalid json
	{
		bf := newFileBody(ctx, tf, int64(len(content)))
		if _, err := bf.AsJson(); err != es_response.ErrorContentIsNotAJSON {
			t.Error(err)
		}
	}
}
