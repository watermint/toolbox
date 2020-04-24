package response

import (
	"github.com/tidwall/gjson"
	"github.com/watermint/toolbox/essentials/format/tjson"
	"github.com/watermint/toolbox/essentials/rec"
	"github.com/watermint/toolbox/infra/api/api_context"
	"go.uber.org/zap"
	"io/ioutil"
)

func newFileBody(ctx api_context.Context, path string, contentLength int64) Body {
	return &bodyFileImpl{
		ctx:           ctx,
		path:          path,
		contentLength: contentLength,
	}
}

type bodyFileImpl struct {
	ctx           api_context.Context
	path          string
	contentLength int64
}

func (z bodyFileImpl) BodyString() string {
	return ""
}

func (z bodyFileImpl) AsJson() (tjson.Json, error) {
	l := z.ctx.Log().With(zap.String("path", z.path))
	if z.contentLength > MaximumJsonSize {
		l.Debug("content is too large for parse", zap.Int64("size", z.contentLength))
		return nil, ErrorContentIsTooLarge
	}
	content, err := ioutil.ReadFile(z.path)
	if err != nil {
		l.Debug("unable to read file", zap.Error(err))
		return nil, err
	}
	if !gjson.ValidBytes(content) {
		l.Debug("invalid bytes", zap.Any("content", rec.ByteDigest(content)))
		return nil, ErrorContentIsNotAJSON
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
