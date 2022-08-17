package api_auth_impl

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/api/api_auth"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/security/sc_token"
	"sort"
	"strings"
)

func NewConsoleCacheOnly(c app_control.Control, peerName string, app api_auth.OAuthApp) api_auth.OAuthConsole {
	return NewConsoleCache(c, dbx_auth.NewConsoleNoAuth(peerName), app)
}

func NewConsoleCache(c app_control.Control, auth api_auth.OAuthConsole, app api_auth.OAuthApp) api_auth.OAuthConsole {
	return &Cached{
		app:  app,
		ctl:  c,
		auth: auth,
		s:    sc_token.NewObfuscated(c, auth.PeerName()),
	}
}

type Cached struct {
	app  api_auth.OAuthApp
	ctl  app_control.Control
	auth api_auth.OAuthConsole
	s    sc_token.Storage
}

func (z *Cached) PeerName() string {
	return z.auth.PeerName()
}

func (z *Cached) Purge(scope string) {
	z.s.Purge(scope)
}

func (z *Cached) Start(scopes []string) (tc api_auth.Context, err error) {
	sort.Strings(scopes)
	cacheKey := strings.Join(scopes, ",")
	l := z.ctl.Log().With(esl.String("peerName", z.auth.PeerName()), esl.Strings("scopes", scopes))
	t, err := z.s.Get(cacheKey)
	if err != nil {
		l.Debug("Unable to load from the cache", esl.Error(err))
	} else {
		return api_auth.NewContext(t, z.app.Config(scopes), z.auth.PeerName(), scopes), nil
	}
	tc, err = z.auth.Start(scopes)
	if err != nil {
		return nil, err
	}

	l.Debug("Update cache")
	if err := z.s.Put(cacheKey, tc.Token()); err != nil {
		l.Debug("Unable to update cache", esl.Error(err))
		// fall thru
	}
	return tc, nil
}
