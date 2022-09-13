package api_auth

import (
	"encoding/base64"
	"errors"
	"strings"
)

type BasicCredential struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (z BasicCredential) Serialize() string {
	// RFC 7617
	return z.Username + ":" + z.Password
}

func (z BasicCredential) HeaderValue() string {
	return "Basic " + base64.StdEncoding.EncodeToString([]byte(z.Serialize()))
}

type BasicEntity struct {
	// App key name
	KeyName string `json:"key_name"`

	// Peer name
	PeerName string `json:"peer_name"`

	// Credential
	Credential BasicCredential `json:"credential"`

	// Supplemental information (e.g. email address of the authenticated account)
	Description string `json:"description,omitempty"`

	// Timestamp of the entity created/updated (RFC3339 format)
	Timestamp string `json:"timestamp,omitempty"`
}

func (z BasicEntity) Entity() Entity {
	return Entity{
		KeyName:     z.KeyName,
		Scope:       "",
		PeerName:    z.PeerName,
		Credential:  z.Credential.Serialize(),
		Description: z.Description,
		Timestamp:   z.Timestamp,
	}
}

func DeserializeBasicCredential(credential string) (BasicCredential, error) {
	dat := strings.Split(credential, ":")
	if len(dat) != 2 {
		return BasicCredential{}, errors.New("invalid format")
	} else {
		return BasicCredential{
			Username: dat[0],
			Password: dat[1],
		}, nil
	}
}

func DeserializeBasicEntity(e Entity) (BasicEntity, error) {
	cred, err := DeserializeBasicCredential(e.Credential)
	if err != nil {
		return BasicEntity{}, err
	}
	return BasicEntity{
		KeyName:     e.KeyName,
		PeerName:    e.PeerName,
		Credential:  cred,
		Description: e.Description,
		Timestamp:   e.Timestamp,
	}, nil
}
