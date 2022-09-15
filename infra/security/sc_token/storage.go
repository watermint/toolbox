package sc_token

import (
	"crypto/sha256"
	"fmt"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/app"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/security/sc_storage"
	"golang.org/x/oauth2"
	"os"
	"path/filepath"
)

type Storage interface {
	PeerName() string
	Get(scope string) (token *oauth2.Token, err error)
	Put(scope string, token *oauth2.Token) error
	Purge(scope string)
}

func storagePath(c app_control.Control, peerName, scope, suffix string) string {
	s := sha256.Sum224([]byte(peerName + scope + app.BuildInfo.Xap))
	return filepath.Join(c.Workspace().Secrets(), fmt.Sprintf("%x.%s", s, suffix))
}

func NewObfuscated(c app_control.Control, peerName string) Storage {
	return &ObfuscatedStorage{
		peerName: peerName,
		c:        c,
		s:        sc_storage.NewStorage(c),
	}
}

type ObfuscatedStorage struct {
	peerName string
	c        app_control.Control
	s        sc_storage.Storage
}

func (z *ObfuscatedStorage) PeerName() string {
	return z.peerName
}

func (z *ObfuscatedStorage) pathAndLog(scope string) (path string, l esl.Logger) {
	l = z.c.Log().With(esl.String("peerName", z.peerName))
	p := z.path(scope)
	l = l.With(esl.String("path", p))
	return p, l
}

func (z *ObfuscatedStorage) path(scope string) string {
	return storagePath(z.c, z.peerName, scope, "obf")
}

func (z *ObfuscatedStorage) Purge(scope string) {
	p, l := z.pathAndLog(scope)

	l.Debug("Purge obfuscate storage")
	if err := os.Remove(p); err != nil {
		l.Debug("Unable to purge", esl.Error(err))
	}
}

func (z *ObfuscatedStorage) Get(scope string) (token *oauth2.Token, err error) {
	p, l := z.pathAndLog(scope)

	l.Debug("Load obfuscated storage")
	token = &oauth2.Token{}
	if err = z.s.Get(p, token); err != nil {
		return nil, err
	}
	return token, nil
}

func (z *ObfuscatedStorage) Put(scope string, token *oauth2.Token) error {
	p, l := z.pathAndLog(scope)

	l.Debug("Load obfuscated storage")
	if err := z.s.Put(p, token); err != nil {
		l.Debug("Unable to store", esl.Error(err))
		return err
	}
	return nil
}
