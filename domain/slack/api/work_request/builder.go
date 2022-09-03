package work_request

import (
	"encoding/json"
	"github.com/google/go-querystring/query"
	"github.com/watermint/toolbox/essentials/api/api_auth"
	api_request2 "github.com/watermint/toolbox/essentials/api/api_request"
	"github.com/watermint/toolbox/essentials/io/es_rewinder"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/network/nw_client"
	"github.com/watermint/toolbox/infra/app"
	"github.com/watermint/toolbox/infra/control/app_control"
	"net/http"
	"regexp"
)

func New(ctl app_control.Control, entity api_auth.OAuthEntity) Builder {
	return &builderImpl{
		ctl:    ctl,
		entity: entity,
	}
}

type Builder interface {
	api_request2.Builder
	With(method, url string, data api_request2.RequestData) Builder
}

type builderImpl struct {
	ctl    app_control.Control
	entity api_auth.OAuthEntity
	method string
	url    string
	data   api_request2.RequestData
}

func (z builderImpl) WithData(data api_request2.RequestDatum) api_request2.Builder {
	z.data = z.data.WithDatum(data)
	return z
}

var (
	slackTokenParam   = regexp.MustCompile(`token=([\w-]+)`)
	slackTokenReplace = "token=<SECRET>"
)

func (z builderImpl) FilterUrl(url string) string {
	if slackTokenParam.MatchString(url) {
		return slackTokenParam.ReplaceAllString(url, slackTokenReplace)
	}
	return url
}

func (z builderImpl) Log() esl.Logger {
	l := z.ctl.Log()
	if z.method != "" {
		l = l.With(esl.String("method", z.method))
	}
	if z.url != "" {
		l = l.With(esl.String("url", z.url))
	}
	if !z.entity.IsNoAuth() {
		l = l.With(esl.Strings("scopes", z.entity.Scopes))
	}
	return l
}

func (z builderImpl) ClientHash() string {
	return nw_client.ClientHash(z.entity.HashSeed(), []string{
		"m", z.method,
		"u", z.url,
	})
}

func (z builderImpl) reqContent() es_rewinder.ReadRewinder {
	if z.data.Content() != nil {
		return z.data.Content()
	}
	return es_rewinder.NewReadRewinderOnMemory([]byte(z.Param()))
}

func (z builderImpl) reqHeaders() (headers map[string]string) {
	headers = make(map[string]string)
	headers[api_request2.ReqHeaderUserAgent] = app.UserAgent()
	headers[api_request2.ReqHeaderContentType] = "application/json; charset=UTF-8"
	headers[api_request2.ReqHeaderAccept] = "application/json"
	//if z.token != nil {
	//	headers[api_request.ReqHeaderAuthorization] = "Bearer " + z.token.Token().AccessToken
	//} else {
	//	headers[api_request.ReqHeaderAuthorization] = "MOCK_CALL"
	//}

	for k, v := range z.data.Headers() {
		headers[k] = v
	}
	return
}

func (z builderImpl) Build() (*http.Request, error) {
	l := z.Log()
	qv, err := query.Values(z.data.Query())
	if err != nil {
		l.Debug("Unable to create query", esl.Error(err))
		return nil, err
	}
	if z.method == http.MethodGet && !z.entity.IsNoAuth() {
		qv.Add("token", z.entity.Token.AccessToken)
	}
	encoded := qv.Encode()
	url := z.url
	if encoded != "" {
		url = url + "?" + encoded
	}
	rc := z.reqContent()
	req, err := nw_client.NewHttpRequest(z.method, url, rc)
	if err != nil {
		l.Debug("Unable create request", esl.Error(err))
		return nil, err
	}
	headers := z.reqHeaders()
	for k, v := range headers {
		req.Header.Add(k, v)
	}
	req.ContentLength = rc.Length()

	return req, nil
}

func (z builderImpl) Endpoint() string {
	return z.url
}

func (z builderImpl) Param() string {
	if z.data.Param() == nil || z.method == http.MethodGet {
		return ""
	}

	var reqData map[string]interface{}
	pj := z.data.ParamJson()
	if err := json.Unmarshal(pj, &reqData); err != nil {
		return string(z.data.ParamJson())
	}

	reqData["token"] = z.entity.Token.AccessToken

	pjd, err := json.Marshal(reqData)
	if err != nil {
		return string(z.data.ParamJson())
	}

	return string(pjd)
}

func (z builderImpl) With(method, url string, data api_request2.RequestData) Builder {
	z.method = method
	z.url = url
	z.data = data
	return z
}
