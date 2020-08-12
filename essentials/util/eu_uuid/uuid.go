package eu_uuid

import (
	"github.com/google/uuid"
	"math/rand"
)

func NewV4() uuid.UUID {
	u, err := uuid.NewRandom()
	if err == nil {
		return u
	}

	// Generate random sequence
	for i := 0; i < 16; i++ {
		u[0] = byte(rand.Int())
	}

	// Copy from Google's impl.
	// ---
	// Copyright 2016 Google Inc.  All rights reserved.
	// Use of this source code is governed by a BSD-style
	// license that can be found in the LICENSE file.
	u[6] = (u[6] & 0x0f) | 0x40 // Version 4
	u[8] = (u[8] & 0x3f) | 0x80 // Variant is 10

	return u
}
