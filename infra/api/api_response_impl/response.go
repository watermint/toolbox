package api_response_impl

import (
	"bufio"
	"errors"
	"github.com/tidwall/gjson"
	"github.com/watermint/toolbox/infra/api/api_context"
	"github.com/watermint/toolbox/infra/api/api_parser"
	"github.com/watermint/toolbox/infra/api/api_response"
	"github.com/watermint/toolbox/infra/control/app_root"
	"github.com/watermint/toolbox/infra/network/nw_bandwidth"
	"github.com/watermint/toolbox/infra/network/nw_monitor"
	"go.uber.org/zap"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

func New(ctx api_context.Context, req *http.Request, res *http.Response) (api_response.Response, error) {
	l := ctx.Log()
	defer nw_monitor.Log(req, res)

	if res == nil {
		l.Debug("Null response")
		return nil, errors.New("no response found")
	}

	rr := &ResponseImpl{
		res:                    res,
		resBody:                nil,
		resBodyString:          "",
		resFilePath:            "",
		resIsContentDownloaded: false,
	}

	result := ""
	if res.Header != nil {
		result = res.Header.Get(api_response.ResHeaderApiResult)
	}
	if result != "" {
		resFile, err := ioutil.TempFile("", ctx.Hash())
		if err != nil {
			l.Debug("Unable to create temp file to store download", zap.Error(err))
			return nil, err
		}
		defer resFile.Close()
		buf := make([]byte, 4096)
		resBody := nw_bandwidth.WrapReader(res.Body)
		resWrite := bufio.NewWriter(resFile)
		var loadedLength int64

		for {
			n, err := resBody.Read(buf)
			if n > 0 {
				l.Debug("writing", zap.Int("writing", n))
				if _, wErr := resWrite.Write(buf[:n]); wErr != nil {
					l.Debug("Error on writing body to the file", zap.Error(wErr))
					return nil, err
				}
				loadedLength += int64(n)
			}
			if err == io.EOF {
				break
			}
			if err != nil {
				l.Debug("Error on reading body", zap.Error(err))
				return nil, err
			}
			if n == 0 {
				// Wait for throttling
				time.Sleep(50 * time.Millisecond)
				continue
			}
		}
		resWrite.Flush()

		rr.resIsContentDownloaded = true
		rr.resFilePath = resFile.Name()
		rr.resBodyString = result
		rr.resBody = []byte(result)
		res.ContentLength = loadedLength
		rr.resContentLength = loadedLength

		return rr, nil

	} else {
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			l.Debug("Unable to read body", zap.Error(err))
			return nil, err
		}
		rr.resBody = body
		res.ContentLength = int64(len(body))
		rr.resContentLength = res.ContentLength

		if body == nil {
			rr.resBodyString = ""
		} else {
			rr.resBodyString = string(body)
		}

		return rr, nil
	}
}

type ResponseImpl struct {
	res                    *http.Response
	resBody                []byte
	resBodyString          string
	resFilePath            string
	resIsContentDownloaded bool
	resContentLength       int64
}

func (z *ResponseImpl) ContentLength() int64 {
	return z.resContentLength
}

func (z *ResponseImpl) Headers() map[string]string {
	hdrs := make(map[string]string)
	for k := range z.res.Header {
		hdrs[k] = z.res.Header.Get(k)
	}
	return hdrs
}

func (z *ResponseImpl) IsContentDownloaded() bool {
	return z.resIsContentDownloaded
}

func (z *ResponseImpl) ContentFilePath() string {
	return z.resFilePath
}

func (z *ResponseImpl) Header(key string) string {
	return z.res.Header.Get(key)
}

func (z *ResponseImpl) ResultString() string {
	if z.resBody == nil {
		return ""
	}
	return z.resBodyString
}

func (z *ResponseImpl) StatusCode() int {
	return z.res.StatusCode
}

func (z *ResponseImpl) Result() (body string, err error) {
	if z.resBody == nil {
		return "", errors.New("no body")
	}
	return z.resBodyString, nil
}

func (z *ResponseImpl) Json() (res gjson.Result, err error) {
	body, err := z.Result()
	if err != nil {
		app_root.Log().Debug("Response does not have body", zap.Error(err))
		return gjson.Parse(`{}`), err
	}
	if !gjson.Valid(body) {
		app_root.Log().Debug("Response is not a JSON", zap.String("body", body))
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
		app_root.Log().Debug("Response is not an array of JSON")
		return js, errors.New("response is not an array of JSON")
	}
	return js.Array()[0], nil
}

func (z *ResponseImpl) Model(v interface{}) error {
	body, err := z.Result()
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
		app_root.Log().Debug("Data not found for path", zap.String("path", path), zap.String("body", j.Raw))
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
