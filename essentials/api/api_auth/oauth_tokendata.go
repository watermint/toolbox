package api_auth

import (
	"encoding/json"
	"github.com/watermint/toolbox/essentials/log/esl"
	"golang.org/x/oauth2"
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
