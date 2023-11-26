package euuid

import (
	"fmt"
	"github.com/watermint/toolbox/essentials/islet/eformat/ehex"
	"github.com/watermint/toolbox/essentials/islet/eidiom/eoutcome"
	"regexp"
)

type Version int
type Variant int

const (
	// Version1 Date-time and MAC address
	Version1 Version = iota + 1

	// Version2 Date-time and MAC address, DCE security version
	Version2

	// Version3 Namespace name-based (MD5)
	Version3

	// Version4 Random UUID
	Version4

	// Version5 Namespace name-based (SHA1)
	Version5
)

const (
	Variant0 Variant = iota
	Variant1
	Variant2
	VariantReserved
)

type UUID interface {
	fmt.Stringer

	// Urn returns URN form like `urn:uuid:123e4567-e89b-12d3-a456-426655440000`.
	Urn() string

	// Version of UUID
	Version() Version

	// Variant of UUID
	Variant() Variant

	// IsNil true if the UUID is zero
	IsNil() bool

	// Equals true if same value
	Equals(x UUID) bool
}

const (
	uuidRePattern = `^[0-9a-fA-F]{8}\b-[0-9a-fA-F]{4}\b-[0-9a-fA-F]{4}\b-[0-9a-fA-F]{4}\b-[0-9a-fA-F]{12}$`
)

var (
	uuidRe = regexp.MustCompile(uuidRePattern)
)

func IsUUID(uuid string) bool {
	return uuidRe.MatchString(uuid)
}

func Parse(uuid string) (u UUID, outcome eoutcome.ParseOutcome) {
	if !IsUUID(uuid) {
		return nil, eoutcome.NewParseInvalidFormat("the given string does not conform UUID format")
	}

	// ----+----|----+----|----+----|----+-
	// 123e4567-e89b-12d3-a456-426655440000
	woh := uuid[0:8] + uuid[9:13] + uuid[14:18] + uuid[19:23] + uuid[24:36]
	ud, outcome := ehex.Parse(woh)
	if outcome.IsError() {
		return nil, outcome
	}
	return New(ud)
}
