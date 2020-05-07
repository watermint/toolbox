package es_filesystem

import (
	"errors"
	"net/http"
)

var (
	ErrAlwaysFail = errors.New("always fail")
)

type Empty struct {
}

func (z Empty) Open(name string) (http.File, error) {
	return nil, ErrAlwaysFail
}
