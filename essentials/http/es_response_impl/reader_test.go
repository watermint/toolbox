package es_response_impl

import (
	"bytes"
	"github.com/watermint/toolbox/essentials/http/es_client"
	"github.com/watermint/toolbox/quality/infra/qt_file"
	"io/ioutil"
	"os"
	"testing"
)

func TestRead_InvalidBufferSize(t *testing.T) {
	content := `{"contact":[{"name":"John"}, {"name":"David"}]}`
	buf := ioutil.NopCloser(bytes.NewBufferString(content))

	ctx := es_client.NewMock()
	if _, err := read(ctx, buf, 1, 2); err != ErrorInvalidBufferState {
		t.Error(err)
	}
}

func TestRead_Failure(t *testing.T) {
	content := `{"contact":[{"name":"John"}, {"name":"David"}]}`
	bodyPath, err := qt_file.MakeTestFile("read", content)
	if err != nil {
		t.Error(err)
		return
	}
	defer func() {
		os.RemoveAll(bodyPath)
	}()

	bodyFile, err := os.Open(bodyPath)
	if err != nil {
		t.Error(err)
		return
	}
	bodyFile.Close() // close before read

	ctx := es_client.NewMock()
	body, err := Read(ctx, bodyFile)
	if err == nil {
		t.Error(err, body)
	}
}

func TestRead_Success(t *testing.T) {
	content := `{"contact":[{"name":"John"}, {"name":"David"}]}`
	buf := ioutil.NopCloser(bytes.NewBufferString(content))

	ctx := es_client.NewMock()
	body, err := Read(ctx, buf)
	if err != nil {
		t.Error(err)
	}
	if body.IsFile() || body.File() != "" {
		t.Error(body.IsFile(), body.File())
	}
	if body.ContentLength() != int64(len(content)) {
		t.Error(body.ContentLength())
	}
	if string(body.Body()) != content {
		t.Error(body.Body())
	}
}

func TestRead_SuccessChunked(t *testing.T) {
	content := `{"contact":[{"name":"John"}, {"name":"David"}]}`
	buf := ioutil.NopCloser(bytes.NewBufferString(content))

	ctx := es_client.NewMock()
	body, err := read(ctx, buf, 24, 8)
	if err != nil {
		t.Error(err)
	}
	if !body.IsFile() || body.File() == "" {
		t.Error(body.IsFile(), body.File())
	}
	if body.ContentLength() != int64(len(content)) {
		t.Error(body.ContentLength())
	}
	if len(body.Body()) != 0 {
		t.Error(body.Body())
	}
	readContent, err := ioutil.ReadFile(body.File())
	if err != nil {
		t.Error(err)
		return
	}
	if string(readContent) != content {
		t.Error(readContent)
	}
}
