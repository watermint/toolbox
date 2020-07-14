package dbx_auth

import (
	"github.com/watermint/toolbox/infra/api/api_auth"
	"golang.org/x/oauth2"
	"strings"
)

func NewMock(peerName string) api_auth.Console {
	return &MockConsoleAuth{peerName: peerName}
}

func NewMockWithPreset(peerName string, preset map[string]*oauth2.Token) api_auth.Console {
	return &MockConsoleAuth{peerName: peerName, preset: preset}
}

type MockContext struct {
	peerName string
	scopes   []string
	preset   *oauth2.Token
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

func (z *MockConsoleAuth) Auth(scopes []string) (token api_auth.Context, err error) {
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
