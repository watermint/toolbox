package nw_request

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"github.com/watermint/toolbox/essentials/network/nw_client"
	"github.com/watermint/toolbox/infra/api/api_request"
	"net/http"
	"strings"
)

type Req struct {
	RequestMethod  string            `json:"method"`
	RequestUrl     string            `json:"url"`
	RequestParam   string            `json:"param,omitempty"`
	RequestHeaders map[string]string `json:"headers"`
	ContentLength  int64             `json:"content_length"`
	RequestHash    string            `json:"hash"`
}

func (z *Req) Apply(rb nw_client.RequestBuilder, req *http.Request) {
	url := req.URL.String()
	param := rb.Param()
	z.RequestHash = HashSeed{
		Url:    url,
		Method: req.Method,
		Param:  param,
		Length: req.ContentLength,
		Header: z.RequestHeaders,
	}.Hash()

	if ruf, ok := rb.(nw_client.RequestUrlFilter); ok {
		url = ruf.FilterUrl(url)
	}
	z.RequestMethod = req.Method
	z.RequestUrl = url
	z.RequestParam = param
	z.RequestHeaders = make(map[string]string)
	z.ContentLength = req.ContentLength
	for k, v := range req.Header {
		v0 := v[0]
		// Anonymize token
		if k == api_request.ReqHeaderAuthorization {
			vv := strings.Split(v0, " ")
			z.RequestHeaders[k] = vv[0] + " <secret>"
		} else {
			z.RequestHeaders[k] = v0
		}
	}
}

type HashSeed struct {
	Url    string            `json:"u"`
	Method string            `json:"m"`
	Param  string            `json:"p"`
	Length int64             `json:"l"`
	Header map[string]string `json:"h"`
}

func (z HashSeed) Hash() string {
	seed := "u" + z.Url +
		"m" + z.Method +
		"p" + z.Param +
		"l" + fmt.Sprintf("%x", z.Length)
	for k, v := range z.Header {
		seed += "h" + k + ":" + v
	}
	h := sha256.Sum256([]byte(seed))
	return base64.RawStdEncoding.EncodeToString(h[:])
}
