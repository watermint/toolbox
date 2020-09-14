package eq_pipe_preserve

import "github.com/watermint/toolbox/essentials/log/esl"

func NewFactory(l esl.Logger, basePath string) Factory {
	return &simpleFactory{
		l:        l,
		basePath: basePath,
	}
}

type simpleFactory struct {
	l        esl.Logger
	basePath string
}

func (z simpleFactory) NewPreserver() Preserver {
	return NewPreserver(z.l, z.basePath)
}

func (z simpleFactory) NewRestorer(sessionId string) Restorer {
	return NewRestorer(z.l, z.basePath, sessionId)
}
