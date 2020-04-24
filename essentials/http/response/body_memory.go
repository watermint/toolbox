package response

import (
	"github.com/tidwall/gjson"
	"github.com/watermint/toolbox/essentials/format/tjson"
	"github.com/watermint/toolbox/essentials/rec"
	"github.com/watermint/toolbox/infra/api/api_context"
	"go.uber.org/zap"
	"io/ioutil"
	"os"
)

func newMemoryBody(ctx api_context.Context, content []byte) Body {
	return &bodyMemoryImpl{
		ctx:     ctx,
		content: content,
	}
}

type bodyMemoryImpl struct {
	ctx     api_context.Context
	content []byte
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
		return nil, ErrorContentIsNotAJSON
	}
	return tjson.Parse(z.content)
}

func (z bodyMemoryImpl) AsFile() (string, error) {
	l := z.ctx.Log()
	p, err := ioutil.TempFile("", z.ctx.ClientHash())
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
	if err := ioutil.WriteFile(p.Name(), z.content, 0600); err != nil {
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
