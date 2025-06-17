package api_auth

import (
	"sync"
	"testing"
)

// mockRepository implements Repository interface for testing
type mockRepository struct {
	mu       sync.Mutex
	entities map[string]Entity
	closed   bool
}

func newMockRepository() *mockRepository {
	return &mockRepository{
		entities: make(map[string]Entity),
	}
}

func (m *mockRepository) makeKey(keyName, scope, peerName string) string {
	return keyName + "|" + scope + "|" + peerName
}

func (m *mockRepository) Put(entity Entity) {
	m.mu.Lock()
	defer m.mu.Unlock()
	
	if m.closed {
		panic("repository is closed")
	}
	
	key := m.makeKey(entity.KeyName, entity.Scope, entity.PeerName)
	m.entities[key] = entity
}

func (m *mockRepository) Get(keyName, scope, peerName string) (entity Entity, found bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	
	if m.closed {
		panic("repository is closed")
	}
	
	key := m.makeKey(keyName, scope, peerName)
	entity, found = m.entities[key]
	return
}

func (m *mockRepository) Delete(keyName, scope, peerName string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	
	if m.closed {
		panic("repository is closed")
	}
	
	key := m.makeKey(keyName, scope, peerName)
	delete(m.entities, key)
}

func (m *mockRepository) List(keyName, scope string) []Entity {
	m.mu.Lock()
	defer m.mu.Unlock()
	
	if m.closed {
		panic("repository is closed")
	}
	
	var result []Entity
	prefix := keyName + "|" + scope + "|"
	
	for key, entity := range m.entities {
		if len(key) >= len(prefix) && key[:len(prefix)] == prefix {
			result = append(result, entity)
		}
	}
	
	return result
}

func (m *mockRepository) Close() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.closed = true
}

func (m *mockRepository) All() []Entity {
	m.mu.Lock()
	defer m.mu.Unlock()
	
	var result []Entity
	for _, entity := range m.entities {
		result = append(result, entity)
	}
	return result
}

func TestRepository_PutAndGet(t *testing.T) {
	repo := newMockRepository()
	defer repo.Close()
	
	entity := Entity{
		KeyName:     "test-key",
		Scope:       "test-scope",
		PeerName:    "test-peer",
		Credential:  "test-credential",
		Description: "test description",
		Timestamp:   "2024-01-01T00:00:00Z",
	}
	
	// Put entity
	repo.Put(entity)
	
	// Get entity
	retrieved, found := repo.Get("test-key", "test-scope", "test-peer")
	if !found {
		t.Error("Entity not found after Put")
	}
	
	if retrieved.KeyName != entity.KeyName {
		t.Errorf("KeyName mismatch: expected %s, got %s", entity.KeyName, retrieved.KeyName)
	}
	if retrieved.Credential != entity.Credential {
		t.Errorf("Credential mismatch: expected %s, got %s", entity.Credential, retrieved.Credential)
	}
	
	// Get non-existent entity
	_, found = repo.Get("non-existent", "test-scope", "test-peer")
	if found {
		t.Error("Non-existent entity should not be found")
	}
}

func TestRepository_Delete(t *testing.T) {
	repo := newMockRepository()
	defer repo.Close()
	
	entity := Entity{
		KeyName:  "delete-key",
		Scope:    "delete-scope",
		PeerName: "delete-peer",
	}
	
	// Put entity
	repo.Put(entity)
	
	// Verify it exists
	_, found := repo.Get("delete-key", "delete-scope", "delete-peer")
	if !found {
		t.Error("Entity should exist before delete")
	}
	
	// Delete entity
	repo.Delete("delete-key", "delete-scope", "delete-peer")
	
	// Verify it's gone
	_, found = repo.Get("delete-key", "delete-scope", "delete-peer")
	if found {
		t.Error("Entity should not exist after delete")
	}
}

func TestRepository_List(t *testing.T) {
	repo := newMockRepository()
	defer repo.Close()
	
	// Put multiple entities with same key and scope
	entities := []Entity{
		{KeyName: "list-key", Scope: "list-scope", PeerName: "peer1"},
		{KeyName: "list-key", Scope: "list-scope", PeerName: "peer2"},
		{KeyName: "list-key", Scope: "list-scope", PeerName: "peer3"},
		{KeyName: "other-key", Scope: "list-scope", PeerName: "peer4"},
		{KeyName: "list-key", Scope: "other-scope", PeerName: "peer5"},
	}
	
	for _, e := range entities {
		repo.Put(e)
	}
	
	// List entities for list-key/list-scope
	results := repo.List("list-key", "list-scope")
	if len(results) != 3 {
		t.Errorf("Expected 3 entities, got %d", len(results))
	}
	
	// Verify all results match the criteria
	for _, r := range results {
		if r.KeyName != "list-key" || r.Scope != "list-scope" {
			t.Errorf("Unexpected entity in results: %+v", r)
		}
	}
	
	// List entities for other combinations
	results = repo.List("other-key", "list-scope")
	if len(results) != 1 {
		t.Errorf("Expected 1 entity for other-key, got %d", len(results))
	}
	
	results = repo.List("list-key", "other-scope")
	if len(results) != 1 {
		t.Errorf("Expected 1 entity for other-scope, got %d", len(results))
	}
	
	// List non-existent combination
	results = repo.List("non-existent", "non-existent")
	if len(results) != 0 {
		t.Errorf("Expected 0 entities for non-existent, got %d", len(results))
	}
}

func TestRepositoryTraversable_All(t *testing.T) {
	repo := newMockRepository()
	defer repo.Close()
	
	// Put some entities
	entities := []Entity{
		{KeyName: "key1", Scope: "scope1", PeerName: "peer1"},
		{KeyName: "key2", Scope: "scope2", PeerName: "peer2"},
		{KeyName: "key3", Scope: "scope3", PeerName: "peer3"},
	}
	
	for _, e := range entities {
		repo.Put(e)
	}
	
	// Get all entities
	all := repo.All()
	if len(all) != 3 {
		t.Errorf("Expected 3 entities, got %d", len(all))
	}
}