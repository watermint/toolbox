package api_auth

import (
	"encoding/json"
	"github.com/watermint/toolbox/essentials/log/esl"
	"golang.org/x/oauth2"
	"sort"
	"strings"
	"time"
)

type OAuthTokenData struct {
	// AccessToken is the token that authorizes and authenticates the requests.
	AccessToken string `json:"access_token"`

	// RefreshToken is a token that's used by the application (as opposed to the user) to refresh the access token if it expires.
	RefreshToken string `json:"refresh_token,omitempty"`

	// Expiry is the optional expiration time of the access token.
	Expiry time.Time `json:"expiry,omitempty"`
}

func (z OAuthTokenData) OAuthToken() *oauth2.Token {
	return &oauth2.Token{
		AccessToken:  z.AccessToken,
		RefreshToken: z.RefreshToken,
		Expiry:       z.Expiry,
	}
}

func (z OAuthTokenData) Serialize() string {
	serialized, err := json.Marshal(&z)
	if err != nil {
		l := esl.Default()
		l.Debug("Unable to marshal", esl.Error(err))
		return ""
	}
	return string(serialized)
}

func DeserializeOAuthTokenData(d string) (OAuthTokenData, error) {
	otd := OAuthTokenData{}
	if err := json.Unmarshal([]byte(d), &otd); err != nil {
		return otd, err
	}
	return otd, nil
}

const (
	OAuthScopeSeparator = " "
)

type OAuthEntity struct {
	// App key name to retrieve client_id/client_secret
	KeyName string

	// Scopes
	Scopes []string

	// Peer name
	PeerName string

	// Serialized credential
	Token OAuthTokenData

	// Supplemental information (e.g. email address of the authenticated account)
	Description string
}

func (z OAuthEntity) Entity() Entity {
	return Entity{
		KeyName:     z.KeyName,
		Scope:       OAuthScopeSerialize(z.Scopes),
		PeerName:    z.PeerName,
		Credential:  z.Token.Serialize(),
		Description: z.Description,
	}
}

func OAuthScopeSerialize(scopes []string) string {
	sc := make([]string, len(scopes))
	copy(sc[:], scopes[:])
	sort.Strings(sc)
	return strings.Join(sc, OAuthScopeSeparator)
}

func DeserializeOAuthEntity(e Entity) (OAuthEntity, error) {
	scopes := strings.Split(e.Scope, OAuthScopeSeparator)
	token, err := DeserializeOAuthTokenData(e.Credential)
	if err != nil {
		return OAuthEntity{}, err
	}
	return OAuthEntity{
		KeyName:     e.KeyName,
		Scopes:      scopes,
		PeerName:    e.PeerName,
		Token:       token,
		Description: e.Description,
	}, nil
}

type OAuthRepository interface {
	// Put to store the Entity.
	Put(entity OAuthEntity)

	// Get to retrieve the Entity.
	Get(keyName string, scopes []string, peerName string) (entity OAuthEntity, found bool)

	// Delete to purge the Entity from the repository.
	Delete(keyName string, scopes []string, peerName string)

	// List to retrieve all Entities matches appKeyName/scope combination.
	List(keyName string, scopes []string) (entities []OAuthEntity)

	// Close the repository
	Close()
}
