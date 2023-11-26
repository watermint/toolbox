package euuid

import (
	"crypto/rand"
	"github.com/watermint/toolbox/essentials/islet/eformat/ehex"
	"github.com/watermint/toolbox/essentials/islet/eidiom/eoutcome"
)

func NewV4() UUID {
	u := make([]byte, 16)

	// Generate random sequence
	_, err := rand.Read(u)
	if err != nil {
		panic(err)
	}

	// Copy from Google's impl.
	// ---
	// Copyright 2016 Google Inc.  All rights reserved.
	// Use of this source code is governed by a BSD-style
	// license that can be found in the LICENSE file.
	u[6] = (u[6] & 0x0f) | 0x40 // Version 4
	u[8] = (u[8] & 0x3f) | 0x80 // Variant is 10

	return &uuidData{
		u: u,
	}
}

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
	return ehex.ToHexString(z.u[0:4]) + "-" +
		ehex.ToHexString(z.u[4:6]) + "-" +
		ehex.ToHexString(z.u[6:8]) + "-" +
		ehex.ToHexString(z.u[8:10]) + "-" +
		ehex.ToHexString(z.u[10:16])
}

func (z uuidData) Urn() string {
	return "urn:uuid:" + z.String()
}
