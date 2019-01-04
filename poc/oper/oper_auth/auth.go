package oper_auth

import (
	"github.com/watermint/toolbox/poc/oper/oper_api"
)

type Authenticator interface {
	Auth(t oper_api.DropboxApiToken) (oper_api.DropboxApiToken, error)
}
