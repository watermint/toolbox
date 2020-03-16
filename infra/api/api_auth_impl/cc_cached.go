package api_auth_impl

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/watermint/toolbox/infra/api/api_auth"
	"github.com/watermint/toolbox/infra/api/api_context"
	"github.com/watermint/toolbox/infra/api/api_context_impl"
	"github.com/watermint/toolbox/infra/app"
	"github.com/watermint/toolbox/infra/control/app_control"
	"go.uber.org/zap"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

type CcCachedAuth struct {
	peerName string
	tokens   map[string]string
	control  app_control.Control
	auth     api_auth.Console
}

func (z *CcCachedAuth) Auth(tokenType string) (ctx api_context.Context, err error) {
	l := z.control.Log().With(zap.String("tokenType", tokenType))
	if tok, e := z.tokens[tokenType]; e {
		l.Debug("Token found")
		tc := api_auth.TokenContainer{
			Token:     tok,
			TokenType: tokenType,
		}
		return api_context_impl.New(z.control, tc), nil
	}
	if z.auth == nil {
		l.Debug("No auth method found")
		return nil, errors.New("no authentication method")
	}
	if ctx, err = z.auth.Auth(tokenType); err != nil {
		l.Debug("Auth failure", zap.Error(err))
		return nil, err
	} else {
		l.Debug("Auth success, updating cache")
		z.updateCache(tokenType, ctx)
		return ctx, nil
	}
}

func (z *CcCachedAuth) init() {
	l := z.control.Log()
	z.tokens = make(map[string]string)

	if err := z.loadFile(); err != nil {
		l.Debug("Unable to load file", zap.Error(err))
	}
}

func (z *CcCachedAuth) cacheFile(kind string, compatible bool) string {
	if compatible {
		return filepath.Join(z.control.Workspace().Secrets(), z.peerName+"."+kind)
	} else {
		px := sha256.Sum224([]byte(z.peerName + app.BuilderKey + app.Name))
		pn := fmt.Sprintf("%x.%s", px, kind)
		return filepath.Join(z.control.Workspace().Secrets(), pn)
	}
}

func (z *CcCachedAuth) compatibleCachedFile() string {
	return z.cacheFile("tokens", true)
}
func (z *CcCachedAuth) secureCachedFile() string {
	return z.cacheFile("t", false)
}
func (z *CcCachedAuth) loadBytes(tb []byte) error {
	err := json.Unmarshal(tb, &z.tokens)
	if err != nil {
		z.control.Log().Debug("unable to unmarshal tokens file", zap.Error(err))
		return err
	}
	return nil
}

func (z *CcCachedAuth) loadFile() error {
	if app.BuilderKey != "" {
		_, err := z.loadSecureFile()
		return err
	}
	if _, err := z.loadCompatibleFile(); err == nil {
		return nil
	}
	return nil
}

func (z *CcCachedAuth) loadCompatibleFile() (exists bool, err error) {
	tf := z.compatibleCachedFile()
	_, err = os.Stat(tf)
	if os.IsNotExist(err) {
		//z.ec.Log().Debug("token file not found", zap.String("path", tf))
		return false, err
	}
	z.control.Log().Debug("Loading token file", zap.String("file", tf))
	tb, err := ioutil.ReadFile(tf)
	if err != nil {
		z.control.Log().Debug("unable to read tokens file", zap.String("path", tf), zap.Error(err))
		return false, err
	}
	return true, z.loadBytes(tb)
}

func (z *CcCachedAuth) loadSecureFile() (exists bool, err error) {
	if app.BuilderKey == "" {
		z.control.Log().Debug("Use compatible token file in dev mode")
		return false, errors.New("dev mode")
	}
	tf := z.secureCachedFile()
	z.control.Log().Debug("Loading token file", zap.String("file", tf))
	_, err = os.Stat(tf)
	if os.IsNotExist(err) {
		//z.ec.Log().Debug("token file not found", zap.String("path", tf))
		return false, err
	}
	tb, err := ioutil.ReadFile(tf)
	if err != nil {
		z.control.Log().Debug("unable to read tokens file", zap.String("path", tf), zap.Error(err))
		return false, err
	}

	key := []byte(app.BuilderKey + app.Name)
	key32 := sha256.Sum224([]byte(key))
	kb := make([]byte, 32)
	copy(kb[:], key32[:])

	bk, err := aes.NewCipher([]byte(kb))
	if err != nil {
		return false, err
	}
	gcm, err := cipher.NewGCM(bk)
	if err != nil {
		return false, err
	}
	ns := gcm.NonceSize()
	nonce, ct := tb[:ns], tb[ns:]
	v, err := gcm.Open(nil, nonce, ct, nil)
	if err != nil {
		return false, err
	}
	return true, z.loadBytes(v)
}

func (z *CcCachedAuth) updateCompatible(tb []byte) error {
	tf := z.compatibleCachedFile()
	err := ioutil.WriteFile(tf, tb, 0600)
	if err != nil {
		z.control.Log().Debug("unable to write tokens into file", zap.Error(err))
		return err
	}
	return nil
}

func (z *CcCachedAuth) updateSecure(tb []byte) error {
	key := []byte(app.BuilderKey + app.Name)
	key32 := sha256.Sum224(key)
	kb := make([]byte, 32)
	copy(kb[:], key32[:])
	bk, err := aes.NewCipher(kb)
	if err != nil {
		return err
	}
	gcm, err := cipher.NewGCM(bk)
	if err != nil {
		return err
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return err
	}
	sealed := gcm.Seal(nonce, nonce, tb, nil)

	tf := z.secureCachedFile()
	err = ioutil.WriteFile(tf, sealed, 0600)
	if err != nil {
		z.control.Log().Debug("unable to write tokens into file", zap.Error(err))
		return err
	}
	return nil
}

func (z *CcCachedAuth) updateCache(tokenType string, ctx api_context.Context) {
	l := z.control.Log().With(zap.String("tokenType", tokenType))
	// Do not store tokens into file
	if z.control.IsSecure() {
		l.Debug("Skip updating cache")
		return
	}

	switch tc := ctx.(type) {
	case api_auth.TokenContext:
		z.tokens[tokenType] = tc.Token().Token
		tb, err := json.Marshal(z.tokens)
		if err != nil {
			l.Debug("unable to marshal tokens", zap.Error(err))
			return
		}
		if app.BuilderKey == "" {
			l.Debug("Updating cache with compatible method")
			z.updateCompatible(tb)
		} else {
			l.Debug("Updating cache with secure method")
			z.updateSecure(tb)
		}

	default:
		l.Debug("Token context type not supported")
	}
}

func (z *CcCachedAuth) convertToCompatible() error {
	l := z.control.Log()
	l.Debug("Loading existing token file")
	if err := z.loadFile(); err != nil {
		return err
	}
	tb, err := json.Marshal(z.tokens)
	if err != nil {
		l.Debug("unable to marshal tokens", zap.Error(err))
		return err
	}
	l.Debug("Store as compatible file")
	return z.updateCompatible(tb)
}

func CreateCompatible(c app_control.Control, peerName string) error {
	a := CcCachedAuth{
		peerName: peerName,
		control:  c,
		auth:     nil,
	}
	a.init()

	return a.convertToCompatible()
}

func CreateSecret(c app_control.Control, peerName string) error {
	l := c.Log()
	l.Debug("Converting existing compatible token file into secure version", zap.String("BUILDER_KEY", app.BuilderKey))

	a := CcCachedAuth{
		peerName: peerName,
		control:  c,
		auth:     nil,
	}
	a.init()
	if e, _ := a.loadCompatibleFile(); !e {
		l.Debug("Unable to load compatible file")
		return errors.New("compatible file not found")
	}
	t, err := json.Marshal(a.tokens)
	if err != nil {
		l.Debug("Unable to marshal token file", zap.Error(err))
		return err
	}
	if err := a.updateSecure(t); err != nil {
		l.Debug("Unable to update token file thru secure method", zap.Error(err))
		return err
	}
	l.Debug("Secure token file created")
	return nil
}
