package api_auth

import (
	"testing"
	"time"
)

func TestEntity_NoCredential(t *testing.T) {
	timestamp := time.Now().Format(time.RFC3339)
	
	entity := Entity{
		KeyName:     "test-key",
		Scope:       "read write",
		PeerName:    "test-peer",
		Credential:  "secret-credential",
		Description: "test description",
		Timestamp:   timestamp,
	}

	noCred := entity.NoCredential()

	// Verify that credential is not included
	if noCred.KeyName != entity.KeyName {
		t.Errorf("KeyName mismatch: expected %s, got %s", entity.KeyName, noCred.KeyName)
	}
	if noCred.Scope != entity.Scope {
		t.Errorf("Scope mismatch: expected %s, got %s", entity.Scope, noCred.Scope)
	}
	if noCred.PeerName != entity.PeerName {
		t.Errorf("PeerName mismatch: expected %s, got %s", entity.PeerName, noCred.PeerName)
	}
	if noCred.Description != entity.Description {
		t.Errorf("Description mismatch: expected %s, got %s", entity.Description, noCred.Description)
	}
	if noCred.Timestamp != entity.Timestamp {
		t.Errorf("Timestamp mismatch: expected %s, got %s", entity.Timestamp, noCred.Timestamp)
	}
}

func TestEntity_Fields(t *testing.T) {
	// Test that all fields can be set and retrieved
	entity := Entity{
		KeyName:     "app-key",
		Scope:       "full-access",
		PeerName:    "peer-1",
		Credential:  "encrypted-token",
		Description: "Test account",
		Timestamp:   "2024-01-01T00:00:00Z",
	}

	if entity.KeyName != "app-key" {
		t.Errorf("KeyName not set correctly")
	}
	if entity.Scope != "full-access" {
		t.Errorf("Scope not set correctly")
	}
	if entity.PeerName != "peer-1" {
		t.Errorf("PeerName not set correctly")
	}
	if entity.Credential != "encrypted-token" {
		t.Errorf("Credential not set correctly")
	}
	if entity.Description != "Test account" {
		t.Errorf("Description not set correctly")
	}
	if entity.Timestamp != "2024-01-01T00:00:00Z" {
		t.Errorf("Timestamp not set correctly")
	}
}

func TestEntityNoCredential_Fields(t *testing.T) {
	// Test that EntityNoCredential doesn't have credential field
	entity := EntityNoCredential{
		KeyName:     "app-key",
		Scope:       "read-only",
		PeerName:    "peer-2",
		Description: "Read-only account",
		Timestamp:   "2024-01-02T00:00:00Z",
	}

	if entity.KeyName != "app-key" {
		t.Errorf("KeyName not set correctly")
	}
	if entity.Scope != "read-only" {
		t.Errorf("Scope not set correctly")
	}
	if entity.PeerName != "peer-2" {
		t.Errorf("PeerName not set correctly")
	}
	if entity.Description != "Read-only account" {
		t.Errorf("Description not set correctly")
	}
	if entity.Timestamp != "2024-01-02T00:00:00Z" {
		t.Errorf("Timestamp not set correctly")
	}
}