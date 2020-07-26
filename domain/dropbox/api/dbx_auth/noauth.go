package dbx_auth

import (
	"errors"
	"github.com/watermint/toolbox/infra/api/api_auth"
)

var (
	ErrorNoAuthDefined = errors.New("no auth defined")
)

func NewConsoleNoAuth(peerName string) api_auth.Console {
	return &NoAuth{peerName: peerName}
}

type NoAuth struct {
	peerName string
}

func (z *NoAuth) PeerName() string {
	return z.peerName
}

func (z *NoAuth) Auth(scopes []string) (tc api_auth.Context, err error) {
	return nil, ErrorNoAuthDefined
}
