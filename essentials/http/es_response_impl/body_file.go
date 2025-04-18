package es_response_impl

import (
	"os"

	"github.com/tidwall/gjson"
	"github.com/watermint/toolbox/essentials/encoding/es_json"
	"github.com/watermint/toolbox/essentials/http/es_client"
	"github.com/watermint/toolbox/essentials/http/es_response"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/log/esl_encode"
)

func newFileBody(ctx es_client.Client, path string, contentLength int64) es_response.Body {
	return &bodyFileImpl{
		ctx:           ctx,
		path:          path,
		contentLength: contentLength,
	}
}

type bodyFileImpl struct {
	ctx           es_client.Client
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
	l := z.ctx.Log().With(esl.String("path", z.path))
	if z.contentLength > es_response.MaximumJsonSize {
		l.Debug("content is too large for parse", esl.Int64("size", z.contentLength))
		return nil, es_response.ErrorContentIsTooLarge
	}
	content, err := os.ReadFile(z.path)
	if err != nil {
		l.Debug("unable to read file", esl.Error(err))
		return nil, err
	}
	if !gjson.ValidBytes(content) {
		l.Debug("invalid bytes", esl.Any("content", esl_encode.ByteDigest(content)))
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
