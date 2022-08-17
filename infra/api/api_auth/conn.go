package api_auth

import (
	"crypto/sha256"
	"fmt"
	"sort"
	"strings"
)

type AuthType struct {
	Type   string   `json:"type"`
	Scopes []string `json:"scopes"`
}

func (z AuthType) Id() string {
	scopes := make([]string, len(z.Scopes))
	copy(scopes[:], z.Scopes[:])
	sort.Strings(scopes)

	v := strings.Join(scopes, ",")
	v = v + ";" + z.Type

	sum := sha256.Sum256([]byte(v))
	return fmt.Sprintf("%x", sum)
}
