package api_response

import (
	"errors"
	"github.com/tidwall/gjson"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/infra/api/api_parser"
	"github.com/watermint/toolbox/infra/control/app_root"
	"go.uber.org/zap"
	"net/http"
)

const (
	DropboxApiResHeaderRetryAfter    = "Rewind-After"
	DropboxApiResHeaderResult        = "Dropbox-API-Result"
	DropboxApiErrorBadInputParam     = 400
	DropboxApiErrorBadOrExpiredToken = 401
	DropboxApiErrorAccessError       = 403
	DropboxApiErrorEndpointSpecific  = 409
	DropboxApiErrorNoPermission      = 422
	DropboxApiErrorRateLimit         = 429
)

var (
	ErrorNoResponseBody  = errors.New("no response body")
	ErrorInvalidFormat   = errors.New("unexpected response body format")
	ErrorNotFoundForPath = errors.New("data not found for the path")
	ErrorNoResponse      = errors.New("no response")
)

type Response interface {
	// Response code. Returns -1 if a response does not contain status code.
	StatusCode() int

	// Returns body string. Returns empty & error if a response does not contain body.
	Result() (body string, err error)

	// Returns body string.
	ResultString() string

	// Content length.
	ContentLength() int64

	// True if the content stored in the file
	IsContentDownloaded() bool

	// Path to the content saved, if the content downloaded. Otherwise empty string.
	ContentFilePath() mo_path.FileSystemPath

	// Header for key. Returns empty if no header for the key.
	Header(key string) string

	// Key value pair of headers
	Headers() map[string]string

	// Returns JSON result. Returns empty & error if a response is not a JSON document.
	Json() (res gjson.Result, err error)

	// Returns first element of the array.
	// Returns empty & error if a response is not an array of JSON
	JsonArrayFirst() (res gjson.Result, err error)

	// Parse model.
	Model(v interface{}) error

	// Parse model with given JSON path.
	ModelWithPath(v interface{}, path string) error

	// Parse model for a first element of the array of JSON.
	ModelArrayFirst(v interface{}) error
}

func New(res *http.Response, body []byte) Response {
	bodyStr := ""
	if body != nil {
		bodyStr = string(body)
	}
	return &responseImpl{
		res:                    res,
		resBody:                body,
		resBodyString:          bodyStr,
		resFilePath:            nil,
		resIsContentDownloaded: false,
		resContentLength:       res.ContentLength,
	}
}

func NewDownload(res *http.Response, path mo_path.FileSystemPath, body string, contentLength int64) Response {
	return &responseImpl{
		res:                    res,
		resBody:                []byte(body),
		resBodyString:          body,
		resFilePath:            path,
		resIsContentDownloaded: true,
		resContentLength:       contentLength,
	}
}

type responseImpl struct {
	res                    *http.Response
	resBody                []byte
	resBodyString          string
	resFilePath            mo_path.FileSystemPath
	resIsContentDownloaded bool
	resContentLength       int64
}

func (z *responseImpl) ContentLength() int64 {
	return z.resContentLength
}

func (z *responseImpl) Headers() map[string]string {
	hdrs := make(map[string]string)
	for k := range z.res.Header {
		hdrs[k] = z.res.Header.Get(k)
	}
	return hdrs
}

func (z *responseImpl) IsContentDownloaded() bool {
	return z.resIsContentDownloaded
}

func (z *responseImpl) ContentFilePath() mo_path.FileSystemPath {
	return z.resFilePath
}

func (z *responseImpl) Header(key string) string {
	return z.res.Header.Get(key)
}

func (z *responseImpl) ResultString() string {
	if z.resBody == nil {
		return ""
	}
	return z.resBodyString
}

func (z *responseImpl) StatusCode() int {
	return z.res.StatusCode
}

func (z *responseImpl) Result() (body string, err error) {
	if z.resBody == nil {
		return "", ErrorNoResponseBody
	}
	return z.resBodyString, nil
}

func (z *responseImpl) Json() (res gjson.Result, err error) {
	body, err := z.Result()
	if err != nil {
		app_root.Log().Debug("Response does not have body", zap.Error(err))
		return gjson.Parse(`{}`), err
	}
	if !gjson.Valid(body) {
		app_root.Log().Debug("Response is not a JSON", zap.String("body", body))
		return gjson.Parse(`{}`), ErrorInvalidFormat
	}
	return gjson.Parse(body), nil
}

func (z *responseImpl) JsonArrayFirst() (res gjson.Result, err error) {
	js, err := z.Json()
	if err != nil {
		return js, err
	}
	if !js.IsArray() {
		app_root.Log().Debug("Response is not an array of JSON")
		return js, ErrorInvalidFormat
	}
	return js.Array()[0], nil
}

func (z *responseImpl) Model(v interface{}) error {
	body, err := z.Result()
	if err != nil {
		return err
	}
	return api_parser.ParseModelString(v, body)
}

func (z *responseImpl) ModelWithPath(v interface{}, path string) error {
	j, err := z.Json()
	if err != nil {
		return err
	}
	p := j.Get(path)
	if !p.Exists() {
		app_root.Log().Debug("Data not found for path", zap.String("path", path), zap.String("body", j.Raw))
		return ErrorNotFoundForPath
	}
	return api_parser.ParseModel(v, p)
}

func (z *responseImpl) ModelArrayFirst(v interface{}) error {
	j, err := z.JsonArrayFirst()
	if err != nil {
		return err
	}
	return api_parser.ParseModel(v, j)
}
