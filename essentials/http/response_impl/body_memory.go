package response_impl

import (
	"github.com/tidwall/gjson"
	"github.com/watermint/toolbox/essentials/format/tjson"
	"github.com/watermint/toolbox/essentials/http/context"
	"github.com/watermint/toolbox/essentials/http/response"
	"github.com/watermint/toolbox/essentials/rec"
	"go.uber.org/zap"
	"io/ioutil"
	"os"
)

func newMemoryBody(ctx context.Context, content []byte) response.Body {
	return &bodyMemoryImpl{
		ctx:     ctx,
		content: content,
	}
}

type bodyMemoryImpl struct {
	ctx     context.Context
	content []byte
}

func (z bodyMemoryImpl) Json() tjson.Json {
	if j, err := z.AsJson(); err != nil {
		return tjson.Null()
	} else {
		return j
	}
}

func (z bodyMemoryImpl) Error() error {
	return nil
}

func (z bodyMemoryImpl) BodyString() string {
	return string(z.content)
}

func (z bodyMemoryImpl) AsJson() (tjson.Json, error) {
	l := z.ctx.Log()
	if !gjson.ValidBytes(z.content) {
		l.Debug("Invalid bytes", zap.Any("bytes", rec.ByteDigest(z.content)))
		return nil, response.ErrorContentIsNotAJSON
	}
	return tjson.Parse(z.content)
}

func toFile(ctx context.Context, content []byte) (string, error) {
	l := ctx.Log()
	p, err := ioutil.TempFile("", ctx.ClientHash())
	if err != nil {
		l.Debug("Unable to create temp file", zap.Error(err))
		return "", err
	}
	cleanupOnError := func() {
		if err := p.Close(); err != nil {
			l.Debug("unable to close", zap.Error(err))
		}
		if err := os.Remove(p.Name()); err != nil {
			l.Debug("unable to remove", zap.Error(err))
		}
	}
	if err := ioutil.WriteFile(p.Name(), content, 0600); err != nil {
		l.Debug("Unable to write", zap.Error(err))
		cleanupOnError()
		return "", err
	}
	if err := p.Close(); err != nil {
		l.Debug("unable to close", zap.Error(err))
		cleanupOnError()
		return "", err
	}
	return p.Name(), nil
}

func (z bodyMemoryImpl) AsFile() (string, error) {
	return toFile(z.ctx, z.content)
}

func (z bodyMemoryImpl) ContentLength() int64 {
	return int64(len(z.content))
}

func (z bodyMemoryImpl) Body() []byte {
	return z.content
}

func (z bodyMemoryImpl) File() string {
	return ""
}

func (z bodyMemoryImpl) IsFile() bool {
	return false
}
