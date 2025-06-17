package api_auth

import (
	"encoding/base64"
	"testing"
)

func TestBasicCredential_Serialize(t *testing.T) {
	cred := BasicCredential{
		Username: "testuser",
		Password: "testpass",
	}
	
	expected := "testuser:testpass"
	if cred.Serialize() != expected {
		t.Errorf("Expected %s, got %s", expected, cred.Serialize())
	}
}

func TestBasicCredential_HeaderValue(t *testing.T) {
	cred := BasicCredential{
		Username: "user",
		Password: "pass",
	}
	
	serialized := cred.Serialize()
	encoded := base64.StdEncoding.EncodeToString([]byte(serialized))
	expected := "Basic " + encoded
	
	if cred.HeaderValue() != expected {
		t.Errorf("Expected %s, got %s", expected, cred.HeaderValue())
	}
	
	// Verify the header value is correctly formatted
	if cred.HeaderValue()[:6] != "Basic " {
		t.Error("Header value should start with 'Basic '")
	}
}

func TestNewNoAuthBasicEntity(t *testing.T) {
	entity := NewNoAuthBasicEntity()
	
	if entity.KeyName != "" {
		t.Error("KeyName should be empty")
	}
	if entity.PeerName != "" {
		t.Error("PeerName should be empty")
	}
	if entity.Credential.Username != "" || entity.Credential.Password != "" {
		t.Error("Credential should be empty")
	}
}

func TestBasicEntity_Entity(t *testing.T) {
	basicEntity := BasicEntity{
		KeyName:  "test-key",
		PeerName: "test-peer",
		Credential: BasicCredential{
			Username: "user",
			Password: "pass",
		},
		Description: "test description",
		Timestamp:   "2024-01-01T00:00:00Z",
	}
	
	entity := basicEntity.Entity()
	
	if entity.KeyName != basicEntity.KeyName {
		t.Errorf("KeyName mismatch: expected %s, got %s", basicEntity.KeyName, entity.KeyName)
	}
	if entity.PeerName != basicEntity.PeerName {
		t.Errorf("PeerName mismatch: expected %s, got %s", basicEntity.PeerName, entity.PeerName)
	}
	if entity.Scope != "" {
		t.Error("Scope should be empty for basic auth")
	}
	if entity.Credential != "user:pass" {
		t.Errorf("Credential mismatch: expected user:pass, got %s", entity.Credential)
	}
	if entity.Description != basicEntity.Description {
		t.Errorf("Description mismatch: expected %s, got %s", basicEntity.Description, entity.Description)
	}
	if entity.Timestamp != basicEntity.Timestamp {
		t.Errorf("Timestamp mismatch: expected %s, got %s", basicEntity.Timestamp, entity.Timestamp)
	}
}

func TestBasicEntity_HashSeed(t *testing.T) {
	basicEntity := BasicEntity{
		KeyName:  "key1",
		PeerName: "peer1",
		Credential: BasicCredential{
			Username: "user1",
			Password: "pass1",
		},
	}
	
	hashSeed := basicEntity.HashSeed()
	
	expected := []string{
		"a", "key1",
		"p", "peer1",
		"c", "user1:pass1",
	}
	
	if len(hashSeed) != len(expected) {
		t.Fatalf("HashSeed length mismatch: expected %d, got %d", len(expected), len(hashSeed))
	}
	
	for i, v := range expected {
		if hashSeed[i] != v {
			t.Errorf("HashSeed[%d] mismatch: expected %s, got %s", i, v, hashSeed[i])
		}
	}
}

func TestDeserializeBasicCredential(t *testing.T) {
	tests := []struct {
		name        string
		credential  string
		wantUser    string
		wantPass    string
		wantErr     bool
	}{
		{
			name:       "valid credential",
			credential: "user:pass",
			wantUser:   "user",
			wantPass:   "pass",
			wantErr:    false,
		},
		{
			name:       "empty password",
			credential: "user:",
			wantUser:   "user",
			wantPass:   "",
			wantErr:    false,
		},
		{
			name:       "empty username",
			credential: ":pass",
			wantUser:   "",
			wantPass:   "pass",
			wantErr:    false,
		},
		{
			name:       "no colon",
			credential: "userpass",
			wantErr:    true,
		},
		{
			name:       "multiple colons",
			credential: "user:pass:extra",
			wantErr:    true,
		},
		{
			name:       "empty string",
			credential: "",
			wantErr:    true,
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cred, err := DeserializeBasicCredential(tt.credential)
			
			if tt.wantErr {
				if err == nil {
					t.Error("Expected error but got none")
				}
				return
			}
			
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
				return
			}
			
			if cred.Username != tt.wantUser {
				t.Errorf("Username mismatch: expected %s, got %s", tt.wantUser, cred.Username)
			}
			if cred.Password != tt.wantPass {
				t.Errorf("Password mismatch: expected %s, got %s", tt.wantPass, cred.Password)
			}
		})
	}
}

func TestDeserializeBasicEntity(t *testing.T) {
	// Test valid entity
	entity := Entity{
		KeyName:     "test-key",
		PeerName:    "test-peer",
		Credential:  "user:pass",
		Description: "test",
		Timestamp:   "2024-01-01T00:00:00Z",
	}
	
	basicEntity, err := DeserializeBasicEntity(entity)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	
	if basicEntity.KeyName != entity.KeyName {
		t.Errorf("KeyName mismatch: expected %s, got %s", entity.KeyName, basicEntity.KeyName)
	}
	if basicEntity.PeerName != entity.PeerName {
		t.Errorf("PeerName mismatch: expected %s, got %s", entity.PeerName, basicEntity.PeerName)
	}
	if basicEntity.Credential.Username != "user" {
		t.Errorf("Username mismatch: expected user, got %s", basicEntity.Credential.Username)
	}
	if basicEntity.Credential.Password != "pass" {
		t.Errorf("Password mismatch: expected pass, got %s", basicEntity.Credential.Password)
	}
	if basicEntity.Description != entity.Description {
		t.Errorf("Description mismatch: expected %s, got %s", entity.Description, basicEntity.Description)
	}
	if basicEntity.Timestamp != entity.Timestamp {
		t.Errorf("Timestamp mismatch: expected %s, got %s", entity.Timestamp, basicEntity.Timestamp)
	}
	
	// Test invalid credential format
	invalidEntity := Entity{
		KeyName:    "test-key",
		PeerName:   "test-peer",
		Credential: "invalid-no-colon",
	}
	
	_, err = DeserializeBasicEntity(invalidEntity)
	if err == nil {
		t.Error("Expected error for invalid credential format")
	}
}

func TestBasicEntity_RoundTrip(t *testing.T) {
	// Test that we can convert BasicEntity -> Entity -> BasicEntity
	original := BasicEntity{
		KeyName:  "round-trip-key",
		PeerName: "round-trip-peer",
		Credential: BasicCredential{
			Username: "rtuser",
			Password: "rtpass",
		},
		Description: "round trip test",
		Timestamp:   "2024-01-01T12:00:00Z",
	}
	
	// Convert to Entity
	entity := original.Entity()
	
	// Convert back to BasicEntity
	restored, err := DeserializeBasicEntity(entity)
	if err != nil {
		t.Fatalf("Failed to deserialize: %v", err)
	}
	
	// Verify all fields match
	if restored.KeyName != original.KeyName {
		t.Errorf("KeyName mismatch after round trip")
	}
	if restored.PeerName != original.PeerName {
		t.Errorf("PeerName mismatch after round trip")
	}
	if restored.Credential.Username != original.Credential.Username {
		t.Errorf("Username mismatch after round trip")
	}
	if restored.Credential.Password != original.Credential.Password {
		t.Errorf("Password mismatch after round trip")
	}
	if restored.Description != original.Description {
		t.Errorf("Description mismatch after round trip")
	}
	if restored.Timestamp != original.Timestamp {
		t.Errorf("Timestamp mismatch after round trip")
	}
}