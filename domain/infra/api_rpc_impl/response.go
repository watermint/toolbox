package api_rpc_impl

import (
	"errors"
	"github.com/tidwall/gjson"
	"github.com/watermint/toolbox/app"
	"github.com/watermint/toolbox/domain/infra/api_parser"
	"go.uber.org/zap"
	"net/http"
)

type ResponseImpl struct {
	resHeader     http.Header
	resBody       []byte
	resBodyString string
	resStatusCode int
}

func (z *ResponseImpl) StatusCode() int {
	return z.resStatusCode
}

func (z *ResponseImpl) Body() (body string, err error) {
	if z.resBody == nil {
		return "", errors.New("no body")
	}
	return z.resBodyString, nil
}

func (z *ResponseImpl) Json() (res gjson.Result, err error) {
	body, err := z.Body()
	if err != nil {
		app.Root().Log().Debug("Response does not have body", zap.Error(err))
		return gjson.Parse(`{}`), err
	}
	if !gjson.Valid(body) {
		app.Root().Log().Debug("Response is not a JSON", zap.String("body", body))
		return gjson.Parse(`{}`), errors.New("not a json data")
	}
	return gjson.Parse(body), nil
}

func (z *ResponseImpl) JsonArrayFirst() (res gjson.Result, err error) {
	js, err := z.Json()
	if err != nil {
		return js, err
	}
	if !js.IsArray() {
		app.Root().Log().Debug("Response is not an array of JSON")
		return js, errors.New("response is not an array of JSON")
	}
	return js.Array()[0], nil
}

func (z *ResponseImpl) Model(v interface{}) error {
	body, err := z.Body()
	if err != nil {
		return err
	}
	return api_parser.ParseModelString(v, body)
}

func (z *ResponseImpl) ModelWithPath(v interface{}, path string) error {
	j, err := z.Json()
	if err != nil {
		return err
	}
	p := j.Get(path)
	if !p.Exists() {
		app.Root().Log().Debug("Data not found for path", zap.String("path", path), zap.String("body", j.Raw))
		return errors.New("data not found for path")
	}
	return api_parser.ParseModel(v, p)
}

func (z *ResponseImpl) ModelArrayFirst(v interface{}) error {
	j, err := z.JsonArrayFirst()
	if err != nil {
		return err
	}
	return api_parser.ParseModel(v, j)
}
