package api_auth_impl

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/json"
	"errors"
	"github.com/watermint/toolbox/infra/api/api_auth"
	"github.com/watermint/toolbox/infra/api/api_context"
	"github.com/watermint/toolbox/infra/api/api_context_impl"
	"github.com/watermint/toolbox/infra/app"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/security/sc_random"
	"go.uber.org/zap"
	"golang.org/x/oauth2"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
)

func NewWeb(control app_control.Control) api_auth.Web {
	w := &Web{
		control:       control,
		app:           NewApp(control),
		sessions:      map[string]WebAuthSession{},
		sessionTokens: map[string]string{},
	}
	return w
}

type WebAuthSession struct {
	TokenType   string
	RedirectUrl string
	PeerName    string
}

type Web struct {
	control       app_control.Control
	app           api_auth.App
	sessions      map[string]WebAuthSession
	sessionTokens map[string]string
	sessionLock   sync.Mutex
	databaseLock  sync.Mutex // database write lock
}

func (z *Web) generatePeerName() string {
	return sc_random.MustGenerateRandomString(8)
}

func (z *Web) New(tokenType, redirectUrl string) (state, url string, err error) {
	z.sessionLock.Lock()
	defer z.sessionLock.Unlock()

	peerName := z.generatePeerName()

	l := z.control.Log().With(
		zap.String("peerName", peerName),
		zap.String("tokenType", tokenType),
	)
	l.Debug("Start Web OAuth sequence", zap.String("redirectUrl", redirectUrl))

	for {
		state, err = sc_random.GenerateRandomString(12)
		if err != nil {
			l.Error("Unable to generate `state`", zap.Error(err))
			return "", "", err
		}
		if _, ok := z.sessions[state]; !ok {
			z.sessions[state] = WebAuthSession{
				TokenType:   tokenType,
				RedirectUrl: redirectUrl,
				PeerName:    peerName,
			}
			break
		}
		// recreate state when conflicted state found in existing sessions
	}

	cfg := z.app.Config(tokenType)
	url = cfg.AuthCodeURL(
		state,
		oauth2.SetAuthURLParam("client_id", cfg.ClientID),
		oauth2.SetAuthURLParam("response_type", "code"),
		oauth2.SetAuthURLParam("redirect_uri", redirectUrl),
	)
	l.Debug("Auth url generated", zap.String("url", url))
	return
}

func (z *Web) Auth(state, code string) (peerName string, ctx api_context.Context, err error) {
	z.sessionLock.Lock()
	defer z.sessionLock.Unlock()

	l := z.control.Log().With(
		zap.String("peerName", peerName),
		zap.String("state", state),
	)
	l.Debug("Start auth sequence")

	session, ok := z.sessions[state]
	if !ok {
		l.Debug("State not found")
		return "", nil, errors.New("state not found")
	}

	cfg := z.app.Config(session.TokenType)
	cfg.RedirectURL = session.RedirectUrl
	token, err := cfg.Exchange(context.Background(), code)
	if err != nil {
		l.Debug("Auth failed", zap.Error(err))
		return "", nil, err
	}

	tc := api_auth.TokenContainer{
		Token:     token.AccessToken,
		TokenType: session.TokenType,
		PeerName:  session.PeerName,
	}
	ctx = api_context_impl.New(z.control, tc)

	desc, suppl, err := VerifyToken(session.TokenType, ctx)
	if err != nil {
		l.Debug("Verification failed", zap.Error(err))
		return "", nil, err
	}
	tc.Description = desc
	tc.Supplemental = suppl

	z.sessionTokens[state] = token.AccessToken
	z.updateDatabase(tc)

	l.Debug("Successfully finished auth sequence")
	return tc.PeerName, ctx, nil
}

func (z *Web) Get(state string) (peerName string, ctx api_context.Context, err error) {
	if t, ok := z.sessionTokens[state]; ok {
		if c, ok := z.sessions[state]; ok {
			tc := api_auth.TokenContainer{
				Token:     t,
				TokenType: c.TokenType,
				PeerName:  c.PeerName,
			}
			ctx = api_context_impl.New(z.control, tc)

			return c.PeerName, ctx, nil
		}
	}
	z.control.Log().Debug("State not found", zap.String("state", state))
	return "", nil, errors.New("state not found")
}

func (z *Web) List(tokenType string) (token []api_auth.TokenContainer, err error) {
	token = make([]api_auth.TokenContainer, 0)
	tf := z.databaseFile(tokenType)
	l := z.control.Log().With(zap.String("tokenType", tokenType))

	_, err = os.Stat(tf)
	if os.IsNotExist(err) {
		return token, nil
	}
	tb, err := ioutil.ReadFile(tf)
	if err != nil {
		l.Debug("unable to read tokens file", zap.String("path", tf), zap.Error(err))
		return
	}

	bk, err := aes.NewCipher(z.databaseKey())
	if err != nil {
		l.Debug("unable to create new cipher", zap.Error(err))
		return
	}
	gcm, err := cipher.NewGCM(bk)
	if err != nil {
		l.Debug("unable to create new GCM", zap.Error(err))
		return
	}
	ns := gcm.NonceSize()
	nonce, ct := tb[:ns], tb[ns:]
	v, err := gcm.Open(nil, nonce, ct, nil)
	if err != nil {
		l.Debug("unable to open gcm", zap.Error(err))
		return
	}
	err = json.Unmarshal(v, &token)
	if err != nil {
		l.Debug("unable to unmarshal tokens file", zap.Error(err))
		return
	}
	return
}

func (z *Web) databaseKey() []byte {
	key := []byte(app.BuilderKey + app.Name)
	key32 := sha256.Sum224([]byte(key))
	kb := make([]byte, 32)
	copy(kb[:], key32[:])

	return kb
}

func (z *Web) databaseFile(tokenType string) string {
	p := z.control.Workspace().Secrets()
	return filepath.Join(p, tokenType+".t")
}

func (z *Web) updateDatabase(tc api_auth.TokenContainer) {
	z.databaseLock.Lock()
	defer z.databaseLock.Unlock()

	l := z.control.Log()

	tokens, err := z.List(tc.TokenType)
	if err != nil {
		l.Debug("Unable to retrieve existing tokens", zap.Error(err))
		return
	}

	tokens = append(tokens, tc)

	tb, err := json.Marshal(tokens)
	if err != nil {
		l.Debug("Unable to marshal tokens into JSON", zap.Error(err))
		return
	}
	bk, err := aes.NewCipher(z.databaseKey())
	if err != nil {
		l.Debug("Unable to create new cipher", zap.Error(err))
		return
	}
	gcm, err := cipher.NewGCM(bk)
	if err != nil {
		l.Debug("Unable to create new GCM", zap.Error(err))
		return
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		l.Debug("Unable to prepare nonce", zap.Error(err))
		return
	}
	sealed := gcm.Seal(nonce, nonce, tb, nil)

	tf := z.databaseFile(tc.TokenType)
	err = ioutil.WriteFile(tf, sealed, 0600)
	if err != nil {
		l.Debug("unable to write tokens into file", zap.Error(err))
		return
	}
}
