package dbx_auth

import (
	"github.com/watermint/toolbox/infra/api/api_auth"
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

func (z *NoAuth) Auth(scope string) (tc api_auth.Context, err error) {
	return nil, ErrorNoAuthDefined
}
