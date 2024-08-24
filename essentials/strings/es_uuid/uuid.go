package es_uuid

import (
	"fmt"
	"github.com/watermint/toolbox/essentials/go/es_idiom_deprecated/eoutcome"
	"github.com/watermint/toolbox/essentials/strings/es_hex"
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

	// Version6 UUIDv6 is a field-compatible version of UUIDv1
	Version6

	// Version7 UUIDv7 features a time-ordered value field derived from the widely implemented and well-known
	// Unix Epoch timestamp source, the number of milliseconds since midnight 1 Jan 1970 UTC, leap seconds excluded.
	Version7

	// Version8 UUIDv8 provides a format for experimental or vendor-specific use cases.
	Version8
)

const (
	Variant0 Variant = iota
	Variant1
	Variant2
	VariantReserved
)

type UUIDMetadata struct {
	UUID    string `json:"uuid"`
	Version int    `json:"version"`
	Variant int    `json:"variant"`
}

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

	// Bytes returns the UUID as byte array
	Bytes() []byte

	// Metadata returns the metadata of UUID
	Metadata() *UUIDMetadata
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
	ud, outcome := es_hex.Parse(woh)
	if outcome.IsError() {
		return nil, outcome
	}
	return New(ud)
}
