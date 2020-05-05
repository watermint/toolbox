package api_auth_impl

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/essentials/log/es_log"
	"github.com/watermint/toolbox/infra/api/api_auth"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/security/sc_token"
)

func NewConsoleCacheOnly(c app_control.Control, peerName string) api_auth.Console {
	return NewConsoleCache(c, dbx_auth.NewConsoleNoAuth(peerName))
}

func NewConsoleCache(c app_control.Control, auth api_auth.Console) api_auth.Console {
	return &Cached{
		ctl:  c,
		auth: auth,
		s:    sc_token.NewObfuscated(c, auth.PeerName()),
	}
}

func IsCacheAvailable(c app_control.Control, peerName string, scopes []string) bool {
	for _, s := range scopes {
		co := NewConsoleCacheOnly(c, peerName)
		_, err := co.Auth(s)
		if err != nil {
			return false
		}
	}
	return true
}

type Cached struct {
	ctl  app_control.Control
	auth api_auth.Console
	s    sc_token.Storage
}

func (z *Cached) PeerName() string {
	return z.auth.PeerName()
}

func (z *Cached) Purge(scope string) {
	z.s.Purge(scope)
}

func (z *Cached) Auth(scope string) (tc api_auth.Context, err error) {
	l := z.ctl.Log().With(es_log.String("peerName", z.auth.PeerName()), es_log.String("scope", scope))
	t, err := z.s.Get(scope)
	if err != nil {
		l.Debug("Unable to load from the cache", es_log.Error(err))
	} else {
		return api_auth.NewContext(t, z.auth.PeerName(), scope), nil
	}
	tc, err = z.auth.Auth(scope)
	if err != nil {
		return nil, err
	}

	l.Debug("Update cache")
	if err := z.s.Put(scope, tc.Token()); err != nil {
		l.Debug("Unable to update cache", es_log.Error(err))
		// fall thru
	}
	return tc, nil
}
