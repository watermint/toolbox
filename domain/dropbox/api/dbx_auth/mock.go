package dbx_auth

import (
	"github.com/watermint/toolbox/infra/api/api_auth"
	"golang.org/x/oauth2"
	"strings"
)

func NewMock(peerName string) api_auth.OAuthConsole {
	return &MockConsoleAuth{peerName: peerName}
}

// Deprecated: NewMockWithPreset
func NewMockWithPreset(peerName string, preset map[string]*oauth2.Token) api_auth.OAuthConsole {
	return &MockConsoleAuth{peerName: peerName, preset: preset}
}

type MockContext struct {
	peerName string
	scopes   []string
	preset   *oauth2.Token
}

func (z *MockContext) Config() *oauth2.Config {
	return &oauth2.Config{}
}

func (z *MockContext) Token() *oauth2.Token {
	return z.preset
}

func (z *MockContext) Scopes() []string {
	return z.scopes
}

func (z *MockContext) PeerName() string {
	return z.peerName
}

func (z *MockContext) Description() string {
	return ""
}

func (z *MockContext) Supplemental() string {
	return ""
}

func (z *MockContext) IsNoAuth() bool {
	return false
}

type MockConsoleAuth struct {
	peerName string
	preset   map[string]*oauth2.Token
}

func (z *MockConsoleAuth) PeerName() string {
	return z.peerName
}

func (z *MockConsoleAuth) Start(scopes []string) (token api_auth.OAuthContext, err error) {
	emptyMock := &MockContext{
		peerName: z.peerName,
		scopes:   scopes,
		preset:   &oauth2.Token{},
	}
	if z.preset == nil {
		return emptyMock, nil
	}
	presetKey := strings.Join(scopes, ",")
	if t, ok := z.preset[presetKey]; ok {
		return &MockContext{
			peerName: z.peerName,
			scopes:   scopes,
			preset:   t,
		}, nil
	}
	return emptyMock, nil
}
