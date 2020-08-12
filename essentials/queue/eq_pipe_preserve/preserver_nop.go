package eq_pipe_preserve

import "errors"

var (
	ErrorPreserveIsNotSupported = errors.New("this pipe does not support preserve")
)

func NopPreserver() Preserver {
	return &nopPreserver{}
}

type nopPreserver struct {
}

func (z nopPreserver) Start() error {
	return ErrorSessionIsNotAvailable
}

func (z nopPreserver) Add(d []byte) error {
	return ErrorSessionIsNotAvailable
}

func (z nopPreserver) Commit(info []byte) (sessionId string, err error) {
	return "", ErrorSessionIsNotAvailable
}
