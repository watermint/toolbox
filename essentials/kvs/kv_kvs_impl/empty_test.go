package kv_kvs_impl

import (
	"encoding/json"
	"testing"
	"github.com/watermint/toolbox/essentials/kvs/kv_kvs"
)

func TestNewEmpty(t *testing.T) {
	kvs := NewEmpty()
	if kvs == nil {
		t.Error("Expected non-nil KVS")
	}
}

func TestEmptyImpl_PutOperations(t *testing.T) {
	kvs := NewEmpty()
	
	// All put operations should succeed without error
	if err := kvs.PutString("key", "value"); err != nil {
		t.Errorf("PutString should not return error, got: %v", err)
	}
	
	if err := kvs.PutJson("key", json.RawMessage(`{"test": "value"}`)); err != nil {
		t.Errorf("PutJson should not return error, got: %v", err)
	}
	
	testModel := map[string]string{"test": "value"}
	if err := kvs.PutJsonModel("key", testModel); err != nil {
		t.Errorf("PutJsonModel should not return error, got: %v", err)
	}
}

func TestEmptyImpl_GetOperations(t *testing.T) {
	kvs := NewEmpty()
	
	// All get operations should return not found error
	value, err := kvs.GetString("key")
	if err != kv_kvs.ErrorNotFound {
		t.Errorf("GetString should return ErrorNotFound, got: %v", err)
	}
	if value != "" {
		t.Errorf("GetString should return empty string, got: %s", value)
	}
	
	jsonMsg, err := kvs.GetJson("key")
	if err != kv_kvs.ErrorNotFound {
		t.Errorf("GetJson should return ErrorNotFound, got: %v", err)
	}
	if jsonMsg != nil {
		t.Errorf("GetJson should return nil, got: %v", jsonMsg)
	}
	
	var testModel map[string]string
	err = kvs.GetJsonModel("key", &testModel)
	if err != kv_kvs.ErrorNotFound {
		t.Errorf("GetJsonModel should return ErrorNotFound, got: %v", err)
	}
}

func TestEmptyImpl_Delete(t *testing.T) {
	kvs := NewEmpty()
	
	// Delete should succeed without error
	if err := kvs.Delete("key"); err != nil {
		t.Errorf("Delete should not return error, got: %v", err)
	}
}

func TestEmptyImpl_ForEachOperations(t *testing.T) {
	kvs := NewEmpty()
	
	// ForEach should not call the function (no entries)
	called := false
	err := kvs.ForEach(func(key string, value []byte) error {
		called = true
		return nil
	})
	if err != nil {
		t.Errorf("ForEach should not return error, got: %v", err)
	}
	if called {
		t.Error("ForEach should not call function on empty KVS")
	}
	
	// ForEachRaw should not call the function (no entries)
	called = false
	err = kvs.ForEachRaw(func(key []byte, value []byte) error {
		called = true
		return nil
	})
	if err != nil {
		t.Errorf("ForEachRaw should not return error, got: %v", err)
	}
	if called {
		t.Error("ForEachRaw should not call function on empty KVS")
	}
	
	// ForEachModel should not call the function (no entries)
	called = false
	var testModel map[string]string
	err = kvs.ForEachModel(testModel, func(key string, m interface{}) error {
		called = true
		return nil
	})
	if err != nil {
		t.Errorf("ForEachModel should not return error, got: %v", err)
	}
	if called {
		t.Error("ForEachModel should not call function on empty KVS")
	}
}

