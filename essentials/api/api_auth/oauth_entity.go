package api_auth

import (
	"sort"
	"strings"
)

const (
	OAuthScopeSeparator = " "
)

func NewNoAuthOAuthEntity() OAuthEntity {
	return OAuthEntity{
		KeyName:     "",
		Scopes:      []string{},
		PeerName:    "",
		Token:       OAuthTokenData{},
		Description: "",
	}
}

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

	// Timestamp of the entity created/updated (RFC3339 format)
	Timestamp string
}

func (z OAuthEntity) IsNoAuth() bool {
	return z.KeyName == ""
}

func (z OAuthEntity) Entity() Entity {
	return Entity{
		KeyName:     z.KeyName,
		Scope:       OAuthScopeSerialize(z.Scopes),
		PeerName:    z.PeerName,
		Credential:  z.Token.Serialize(),
		Description: z.Description,
		Timestamp:   z.Timestamp,
	}
}

func (z OAuthEntity) HashSeed() []string {
	seed := make([]string, 0)
	seed = append(seed, []string{
		"a", z.KeyName,
		"p", z.PeerName,
		"t", z.Token.AccessToken,
	}...)
	seed = append(seed, "s")
	seed = append(seed, z.Scopes...)
	return seed
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
		Timestamp:   e.Timestamp,
	}, nil
}
