package es_open

import (
	"errors"
	"github.com/watermint/toolbox/essentials/islet/edesktop"
	"github.com/watermint/toolbox/essentials/log/esl"
)

var (
	ErrorAlwaysError = errors.New("always fail")
)

type Open interface {
	// Open a file, folder, or URI using the OS's default app.
	Open(input string) error
}

func New() Open {
	return &openWrapper{}
}

func NewTestDummy() Open {
	return &dummyOpener{}
}
func NewTestError() Open {
	return &errorOpener{}
}

type openWrapper struct {
}

func (z *openWrapper) Open(input string) error {
	return edesktop.CurrentDesktop().Open(input).Cause()
}

type dummyOpener struct {
}

func (z *dummyOpener) Open(input string) error {
	l := esl.Default()
	l.Debug("Open", esl.String("input", input))
	return nil
}

type errorOpener struct {
}

func (z *errorOpener) Open(input string) error {
	l := esl.Default()
	l.Debug("Open", esl.String("input", input))
	return ErrorAlwaysError
}
