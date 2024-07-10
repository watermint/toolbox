package es_uuid

import (
	"crypto/rand"
	"encoding/binary"
	"errors"
	"time"
)

var (
	ErrInvalidVersion = errors.New("invalid version")
)

func NewV7() UUID {
	return NewV7WithTimestamp(time.Now())
}

func NewV7WithTimestamp(ts time.Time) UUID {
	u := make([]byte, 16)

	epochMillis := ts.UnixNano() / 1_000_000 // nano -> milli
	epochMillisBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(epochMillisBytes, uint64(epochMillis))

	// Copy the 48-bit timestamp
	copy(u[0:6], epochMillisBytes[2:8])

	// Generate random sequence
	_, err := rand.Read(u[6:])
	if err != nil {
		panic(err)
	}

	// Set the version
	u[6] = (u[6] & 0x0f) | 0x70 // Version 7

	// Set the variant
	u[8] = (u[8] & 0x3f) | 0x80 // Variant is 10

	return &uuidData{
		u: u,
	}
}

func TimestampFromUUIDV7(u UUID) (time.Time, error) {
	if u.Version() != Version7 {
		return time.Time{}, ErrInvalidVersion
	}

	// Extract the timestamp from the UUID
	epochMillisBytes := make([]byte, 8)
	copy(epochMillisBytes[2:8], u.Bytes()[0:6])
	epochMillis := int64(binary.BigEndian.Uint64(epochMillisBytes))

	// Convert the timestamp to time.Time
	return time.Unix(0, epochMillis*1_000_000), nil
}
