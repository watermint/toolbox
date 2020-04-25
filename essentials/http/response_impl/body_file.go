package response_impl

import (
	"github.com/tidwall/gjson"
	"github.com/watermint/toolbox/essentials/format/tjson"
	"github.com/watermint/toolbox/essentials/http/context"
	"github.com/watermint/toolbox/essentials/http/response"
	"github.com/watermint/toolbox/essentials/rec"
	"go.uber.org/zap"
	"io/ioutil"
)

func newFileBody(ctx context.Context, path string, contentLength int64) response.Body {
	return &bodyFileImpl{
		ctx:           ctx,
		path:          path,
		contentLength: contentLength,
	}
}

type bodyFileImpl struct {
	ctx           context.Context
	path          string
	contentLength int64
}

func (z bodyFileImpl) Json() tjson.Json {
	if j, err := z.AsJson(); err != nil {
		return tjson.Null()
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

func (z bodyFileImpl) AsJson() (tjson.Json, error) {
	l := z.ctx.Log().With(zap.String("path", z.path))
	if z.contentLength > response.MaximumJsonSize {
		l.Debug("content is too large for parse", zap.Int64("size", z.contentLength))
		return nil, response.ErrorContentIsTooLarge
	}
	content, err := ioutil.ReadFile(z.path)
	if err != nil {
		l.Debug("unable to read file", zap.Error(err))
		return nil, err
	}
	if !gjson.ValidBytes(content) {
		l.Debug("invalid bytes", zap.Any("content", rec.ByteDigest(content)))
		return nil, response.ErrorContentIsNotAJSON
	}
	return tjson.Parse(content)
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
