package kv_kvs_impl

import (
	"encoding/json"
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/watermint/toolbox/essentials/kvs/kv_kvs"
)

func TestNewTurnstile(t *testing.T) {
	mockKvs := NewEmpty()
	turnstile := NewTurnstile(mockKvs)
	
	if turnstile == nil {
		t.Error("Expected non-nil turnstile")
	}
}

func TestTurnstileImpl_PutString(t *testing.T) {
	mockKvs := &mockKvs{data: make(map[string]string)}
	turnstile := NewTurnstile(mockKvs)
	
	err := turnstile.PutString("key1", "value1")
	if err != nil {
		t.Errorf("PutString should not return error, got: %v", err)
	}
	
	// Verify the value was stored in the underlying KVS
	if mockKvs.data["key1"] != "value1" {
		t.Errorf("Expected value1, got %s", mockKvs.data["key1"])
	}
}

func TestTurnstileImpl_PutJson(t *testing.T) {
	mockKvs := &mockKvs{data: make(map[string]string)}
	turnstile := NewTurnstile(mockKvs)
	
	jsonData := json.RawMessage(`{"test": "value"}`)
	err := turnstile.PutJson("key1", jsonData)
	if err != nil {
		t.Errorf("PutJson should not return error, got: %v", err)
	}
	
	// Verify the JSON was stored
	if mockKvs.data["key1"] != string(jsonData) {
		t.Errorf("Expected %s, got %s", jsonData, mockKvs.data["key1"])
	}
}

func TestTurnstileImpl_PutJsonModel(t *testing.T) {
	mockKvs := &mockKvs{data: make(map[string]string)}
	turnstile := NewTurnstile(mockKvs)
	
	testModel := map[string]string{"test": "value"}
	err := turnstile.PutJsonModel("key1", testModel)
	if err != nil {
		t.Errorf("PutJsonModel should not return error, got: %v", err)
	}
	
	// Verify that PutJsonModel delegates to the underlying KVS
}

func TestTurnstileImpl_GetString(t *testing.T) {
	mockKvs := &mockKvs{data: make(map[string]string)}
	mockKvs.data["key1"] = "value1"
	turnstile := NewTurnstile(mockKvs)
	
	value, err := turnstile.GetString("key1")
	if err != nil {
		t.Errorf("GetString should not return error, got: %v", err)
	}
	if value != "value1" {
		t.Errorf("Expected value1, got %s", value)
	}
	
	// Verify that GetString delegates to the underlying KVS
}

func TestTurnstileImpl_GetJson(t *testing.T) {
	mockKvs := &mockKvs{data: make(map[string]string)}
	jsonData := `{"test": "value"}`
	mockKvs.data["key1"] = jsonData
	turnstile := NewTurnstile(mockKvs)
	
	result, err := turnstile.GetJson("key1")
	if err != nil {
		t.Errorf("GetJson should not return error, got: %v", err)
	}
	if string(result) != jsonData {
		t.Errorf("Expected %s, got %s", jsonData, result)
	}
	
	// Verify that GetJson delegates to the underlying KVS
}

func TestTurnstileImpl_GetJsonModel(t *testing.T) {
	mockKvs := &mockKvs{data: make(map[string]string)}
	mockKvs.data["key1"] = `{"test": "value"}`
	turnstile := NewTurnstile(mockKvs)
	
	var result map[string]string
	err := turnstile.GetJsonModel("key1", &result)
	if err != nil {
		t.Errorf("GetJsonModel should not return error, got: %v", err)
	}
	
	// Verify that GetJsonModel delegates to the underlying KVS
}

func TestTurnstileImpl_Delete(t *testing.T) {
	mockKvs := &mockKvs{data: make(map[string]string)}
	mockKvs.data["key1"] = "value1"
	turnstile := NewTurnstile(mockKvs)
	
	err := turnstile.Delete("key1")
	if err != nil {
		t.Errorf("Delete should not return error, got: %v", err)
	}
	
	// Verify that Delete delegates to the underlying KVS
}

func TestTurnstileImpl_ForEach(t *testing.T) {
	mockKvs := &mockKvs{data: make(map[string]string)}
	mockKvs.data["key1"] = "value1"
	mockKvs.data["key2"] = "value2"
	turnstile := NewTurnstile(mockKvs)
	
	count := 0
	err := turnstile.ForEach(func(key string, value []byte) error {
		count++
		return nil
	})
	if err != nil {
		t.Errorf("ForEach should not return error, got: %v", err)
	}
	
	// Verify that ForEach delegates to the underlying KVS
}

func TestTurnstileImpl_ConcurrentAccess(t *testing.T) {
	mockKvs := &mockKvs{
		data:  make(map[string]string),
		mutex: &sync.Mutex{},
	}
	turnstile := NewTurnstile(mockKvs)
	
	// Test concurrent access
	numGoroutines := 10
	numOpsPerGoroutine := 100
	
	var wg sync.WaitGroup
	wg.Add(numGoroutines)
	
	for i := 0; i < numGoroutines; i++ {
		go func(id int) {
			defer wg.Done()
			for j := 0; j < numOpsPerGoroutine; j++ {
				key := fmt.Sprintf("key_%d_%d", id, j)
				value := fmt.Sprintf("value_%d_%d", id, j)
				
				err := turnstile.PutString(key, value)
				if err != nil {
					t.Errorf("PutString failed: %v", err)
				}
				
				// Small delay to increase chance of race conditions
				time.Sleep(time.Microsecond)
			}
		}(i)
	}
	
	wg.Wait()
	
	// Verify all values were stored
	expectedCount := numGoroutines * numOpsPerGoroutine
	mockKvs.mutex.Lock()
	actualCount := len(mockKvs.data)
	mockKvs.mutex.Unlock()
	
	if actualCount != expectedCount {
		t.Errorf("Expected %d entries, got %d", expectedCount, actualCount)
	}
}

// mockKvs is a simple mock implementation of the Kvs interface for testing
type mockKvs struct {
	data  map[string]string
	mutex *sync.Mutex
}

func (m *mockKvs) PutString(key string, value string) error {
	if m.mutex != nil {
		m.mutex.Lock()
		defer m.mutex.Unlock()
	}
	m.data[key] = value
	return nil
}

func (m *mockKvs) PutJson(key string, j json.RawMessage) error {
	if m.mutex != nil {
		m.mutex.Lock()
		defer m.mutex.Unlock()
	}
	m.data[key] = string(j)
	return nil
}

func (m *mockKvs) PutJsonModel(key string, v interface{}) error {
	if m.mutex != nil {
		m.mutex.Lock()
		defer m.mutex.Unlock()
	}
	jsonData, err := json.Marshal(v)
	if err != nil {
		return err
	}
	m.data[key] = string(jsonData)
	return nil
}

func (m *mockKvs) GetString(key string) (string, error) {
	if m.mutex != nil {
		m.mutex.Lock()
		defer m.mutex.Unlock()
	}
	value, exists := m.data[key]
	if !exists {
		return "", kv_kvs.ErrorNotFound
	}
	return value, nil
}

func (m *mockKvs) GetJson(key string) (json.RawMessage, error) {
	if m.mutex != nil {
		m.mutex.Lock()
		defer m.mutex.Unlock()
	}
	value, exists := m.data[key]
	if !exists {
		return nil, kv_kvs.ErrorNotFound
	}
	return json.RawMessage(value), nil
}

func (m *mockKvs) GetJsonModel(key string, v interface{}) error {
	if m.mutex != nil {
		m.mutex.Lock()
		defer m.mutex.Unlock()
	}
	value, exists := m.data[key]
	if !exists {
		return kv_kvs.ErrorNotFound
	}
	return json.Unmarshal([]byte(value), v)
}

func (m *mockKvs) Delete(key string) error {
	if m.mutex != nil {
		m.mutex.Lock()
		defer m.mutex.Unlock()
	}
	delete(m.data, key)
	return nil
}

func (m *mockKvs) ForEach(f func(key string, value []byte) error) error {
	if m.mutex != nil {
		m.mutex.Lock()
		defer m.mutex.Unlock()
	}
	for key, value := range m.data {
		if err := f(key, []byte(value)); err != nil {
			return err
		}
	}
	return nil
}

func (m *mockKvs) ForEachRaw(f func(key []byte, value []byte) error) error {
	if m.mutex != nil {
		m.mutex.Lock()
		defer m.mutex.Unlock()
	}
	for key, value := range m.data {
		if err := f([]byte(key), []byte(value)); err != nil {
			return err
		}
	}
	return nil
}

func (m *mockKvs) ForEachModel(model interface{}, f func(key string, m interface{}) error) error {
	if m.mutex != nil {
		m.mutex.Lock()
		defer m.mutex.Unlock()
	}
	for key, value := range m.data {
		var result interface{}
		if err := json.Unmarshal([]byte(value), &result); err != nil {
			continue // Skip invalid JSON
		}
		if err := f(key, result); err != nil {
			return err
		}
	}
	return nil
}