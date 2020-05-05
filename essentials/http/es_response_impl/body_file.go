package es_response_impl

import (
	"github.com/tidwall/gjson"
	"github.com/watermint/toolbox/essentials/encoding/es_json"
	"github.com/watermint/toolbox/essentials/http/es_context"
	"github.com/watermint/toolbox/essentials/http/es_response"
	"github.com/watermint/toolbox/essentials/log/es_encode"
	"github.com/watermint/toolbox/essentials/log/es_log"
	"io/ioutil"
)

func newFileBody(ctx es_context.Context, path string, contentLength int64) es_response.Body {
	return &bodyFileImpl{
		ctx:           ctx,
		path:          path,
		contentLength: contentLength,
	}
}

type bodyFileImpl struct {
	ctx           es_context.Context
	path          string
	contentLength int64
}

func (z bodyFileImpl) Json() es_json.Json {
	if j, err := z.AsJson(); err != nil {
		return es_json.Null()
	} else {
		return j
	}
}

func (z bodyFileImpl) Error() error {
	return nil
}

func (z bodyFileImpl) BodyString() string {
	return ""
}

func (z bodyFileImpl) AsJson() (es_json.Json, error) {
	l := z.ctx.Log().With(es_log.String("path", z.path))
	if z.contentLength > es_response.MaximumJsonSize {
		l.Debug("content is too large for parse", es_log.Int64("size", z.contentLength))
		return nil, es_response.ErrorContentIsTooLarge
	}
	content, err := ioutil.ReadFile(z.path)
	if err != nil {
		l.Debug("unable to read file", es_log.Error(err))
		return nil, err
	}
	if !gjson.ValidBytes(content) {
		l.Debug("invalid bytes", es_log.Any("content", es_encode.ByteDigest(content)))
		return nil, es_response.ErrorContentIsNotAJSON
	}
	return es_json.Parse(content)
}

func (z bodyFileImpl) AsFile() (string, error) {
	return z.path, nil
}

func (z bodyFileImpl) ContentLength() int64 {
	return z.contentLength
}

func (z bodyFileImpl) Body() []byte {
	return []byte{}
}

func (z bodyFileImpl) File() string {
	return z.path
}

func (z bodyFileImpl) IsFile() bool {
	return true
}
