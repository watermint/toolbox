package nw_http

import (
	"bytes"
	"net/http"
	"testing"
)

func TestCallRpc(t *testing.T) {
	req, err := http.NewRequest("POST", "http://httpbin.org/post", bytes.NewReader([]byte{}))
	if err != nil {
		t.Error("unable to create request", err)
		return
	}
	c := NewClient(&http.Client{})
	res, _, err := c.Call("123", "end/point", req)
	if err != nil {
		t.Error("bad response", err)
	}
	t.Log("code", res.StatusCode)
}
