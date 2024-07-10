package es_uuid

import (
	"github.com/watermint/toolbox/essentials/go/es_idiom_deprecated/eoutcome"
	"github.com/watermint/toolbox/essentials/strings/es_hex"
)

func New(data []byte) (UUID, eoutcome.ParseOutcome) {
	if len(data) != 16 {
		return nil, eoutcome.NewParseInvalidFormat("UUID data length must be 16 bytes")
	}
	d := make([]byte, 16)
	copy(d[:], data[:])
	return uuidData{d}, eoutcome.NewParseSuccess()
}

type uuidData struct {
	u []byte
}

func (z uuidData) Metadata() *UUIDMetadata {
	return &UUIDMetadata{
		UUID:    z.String(),
		Version: int(z.Version()),
		Variant: int(z.Variant()),
	}
}

func (z uuidData) Bytes() []byte {
	return z.u
}

func (z uuidData) Equals(x UUID) bool {
	if x == nil {
		return false
	}
	return z.String() == x.String()
}

func (z uuidData) IsNil() bool {
	for _, x := range z.u {
		if x != 0x00 {
			return false
		}
	}
	return true
}

func (z uuidData) Version() Version {
	x := z.u[6] >> 4
	switch x {
	case 1:
		return Version1
	case 2:
		return Version2
	case 3:
		return Version3
	case 4:
		return Version4
	case 5:
		return Version5
	case 6:
		return Version6
	case 7:
		return Version7
	case 8:
		return Version8
	default:
		return Version(x)
	}
}

func (z uuidData) Variant() Variant {
	x := z.u[8] >> 4
	switch {
	case x&0b1000 == 0:
		return Variant0
	case x&0b1100 == 0b1000:
		return Variant1
	case x&0b1110 == 0b1100:
		return Variant2
	default:
		return VariantReserved
	}
}

func (z uuidData) String() string {
	return es_hex.ToHexString(z.u[0:4]) + "-" +
		es_hex.ToHexString(z.u[4:6]) + "-" +
		es_hex.ToHexString(z.u[6:8]) + "-" +
		es_hex.ToHexString(z.u[8:10]) + "-" +
		es_hex.ToHexString(z.u[10:16])
}

func (z uuidData) Urn() string {
	return "urn:uuid:" + z.String()
}
