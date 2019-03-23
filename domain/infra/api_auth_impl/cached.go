package api_auth_impl

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"github.com/watermint/toolbox/app"
	"github.com/watermint/toolbox/domain/infra/api_auth"
	"github.com/watermint/toolbox/domain/infra/api_context"
	"github.com/watermint/toolbox/domain/infra/api_context_impl"
	"go.uber.org/zap"
	"io/ioutil"
	"os"
)

type CachedAuth struct {
	peerName string
	tokens   map[string]string
	ec       *app.ExecContext
	auth     api_auth.Auth
}

func (z *CachedAuth) Auth(tokenType string) (ctx api_context.Context, err error) {
	if tok, e := z.tokens[tokenType]; e {
		tc := api_auth.TokenContainer{
			Token:     tok,
			TokenType: tokenType,
		}
		return api_context_impl.New(z.ec, tc), nil
	}
	if ctx, err = z.auth.Auth(tokenType); err != nil {
		return nil, err
	} else {
		z.updateCache(tokenType, ctx)
		return ctx, nil
	}
}

func (z *CachedAuth) init() {
	z.tokens = make(map[string]string)

	if z.loadFile() == nil {
		return // return on success
	}
}

func (z *CachedAuth) cacheFile() string {
	px := sha256.Sum224([]byte(z.peerName))
	pn := fmt.Sprintf("%x.tokens", px)
	return z.ec.FileOnSecretsPath(pn)
}

func (z *CachedAuth) loadFile() error {
	tf := z.cacheFile()
	_, err := os.Stat(tf)
	if os.IsNotExist(err) {
		z.ec.Log().Debug("token file not found", zap.String("path", tf))
		return err
	}
	tb, err := ioutil.ReadFile(tf)
	if err != nil {
		z.ec.Log().Debug("unable to read tokens file", zap.String("path", tf), zap.Error(err))
		return err
	}
	err = json.Unmarshal(tb, &z.tokens)
	if err != nil {
		z.ec.Log().Debug("unable to unmarshal tokens file", zap.Error(err))
		return err
	}
	return nil
}

func (z *CachedAuth) updateCache(tokenType string, ctx api_context.Context) {
	// Do not store tokens into file
	if z.ec.NoCacheToken() {
		return
	}

	switch tc := ctx.(type) {
	case api_auth.TokenContext:
		z.tokens[tokenType] = tc.Token().Token
		tb, err := json.Marshal(z.tokens)
		if err != nil {
			z.ec.Log().Debug("unable to marshal tokens", zap.Error(err))
			return
		}
		tf := z.cacheFile()
		err = ioutil.WriteFile(tf, tb, 0600)
		if err != nil {
			z.ec.Log().Debug("unable to write tokens into file", zap.Error(err))
			return
		}
	}
}
