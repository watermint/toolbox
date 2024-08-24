package es_uuid

import "crypto/rand"

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
