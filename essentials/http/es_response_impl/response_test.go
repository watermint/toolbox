package es_response_impl

import (
	"io"
	"net/http"
	"strconv"
	"strings"
	"testing"

	"github.com/watermint/toolbox/essentials/http/es_client"
	"github.com/watermint/toolbox/essentials/http/es_response"
)

func TestResImpl(t *testing.T) {
	content := `{"message":"Hello"}`
	ctx := es_client.NewMock()
	res := &http.Response{
		StatusCode: 200,
		Header: http.Header{
			"Content-Length": []string{strconv.FormatInt(int64(len(content)), 10)},
			"X-Toolbox":      []string{"true"},
		},
		Body:          io.NopCloser(strings.NewReader(content)),
		ContentLength: int64(len(content)),
	}

	r := New(ctx, res)

	if r.Success().BodyString() != content {
		t.Error(r.Success().BodyString())
	}
	if r.Code() != 200 {
		t.Error(r.Code())
	}
	if r.CodeCategory() != es_response.Code2xxSuccess {
		t.Error(r.CodeCategory())
	}
	if x := r.Header("x-toolbox"); x != "true" {
		t.Error(x)
	}
	if len(r.Headers()) != 2 {
		t.Error(r.Headers())
	}
}
