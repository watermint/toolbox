package response

import (
	"github.com/watermint/toolbox/infra/api/api_context"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"testing"
)

func TestResImpl(t *testing.T) {
	content := `{"message":"Hello"}`
	ctx := api_context.NewMock()
	res := &http.Response{
		StatusCode: 200,
		Header: http.Header{
			"Content-Length": []string{strconv.FormatInt(int64(len(content)), 10)},
			"X-Toolbox":      []string{"true"},
		},
		Body:          ioutil.NopCloser(strings.NewReader(content)),
		ContentLength: int64(len(content)),
	}

	r := New(ctx, res)

	if r.Body().BodyString() != content {
		t.Error(r.Body().BodyString())
	}
	if r.Code() != 200 {
		t.Error(r.Code())
	}
	if r.CodeCategory() != Code2xxSuccess {
		t.Error(r.CodeCategory())
	}
	if x := r.Header("x-toolbox"); x != "true" {
		t.Error(x)
	}
	if len(r.Headers()) != 2 {
		t.Error(r.Headers())
	}
}
