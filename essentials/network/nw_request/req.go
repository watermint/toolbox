package nw_request

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"github.com/watermint/toolbox/essentials/api/api_client"
	"github.com/watermint/toolbox/essentials/api/api_request"
	"github.com/watermint/toolbox/essentials/network/nw_client"
	"net/http"
	"sort"
	"strings"
)

type Req struct {
	RequestMethod  string            `json:"method"`
	RequestUrl     string            `json:"url"`
	RequestParam   string            `json:"param,omitempty"`
	RequestHeaders map[string]string `json:"headers"`
	ContentLength  int64             `json:"content_length"`
	RequestHash    string            `json:"hash"`
	Peer           string            `json:"peer"`
}

func (z *Req) Apply(ctx api_client.Client, rb nw_client.RequestBuilder, req *http.Request) {
	url := req.URL.String()
	param := rb.Param()

	if ruf, ok := rb.(nw_client.RequestUrlFilter); ok {
		url = ruf.FilterUrl(url)
	}
	z.Peer = ctx.Name()
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

	z.RequestHash = HashSeed{
		Url:      z.RequestUrl,
		Method:   z.RequestMethod,
		Param:    z.RequestParam,
		Length:   z.ContentLength,
		Header:   z.RequestHeaders,
		PeerName: ctx.Name(),
	}.Hash()
}

type HashSeed struct {
	Url      string            `json:"u"`
	Method   string            `json:"m"`
	Param    string            `json:"p"`
	Length   int64             `json:"l"`
	Header   map[string]string `json:"h"`
	PeerName string            `json:"n"`
}

func (z HashSeed) Hash() string {
	seed := "n" + z.PeerName +
		"u" + z.Url +
		"m" + z.Method +
		"p" + z.Param +
		"l" + fmt.Sprintf("%x", z.Length)

	headers := make([]string, 0)
	for k, v := range z.Header {
		switch k {
		case api_request.ReqHeaderAuthorization,
			api_request.ReqHeaderUserAgent:
			// skip
		default:
			headers = append(headers, "h"+k+":"+v)
		}
	}

	sort.Strings(headers)
	seed += strings.Join(headers, "")
	h := sha256.Sum256([]byte(seed))
	return base64.RawStdEncoding.EncodeToString(h[:])
}
