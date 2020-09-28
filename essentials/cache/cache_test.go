package cache

import (
	"github.com/watermint/toolbox/essentials/kvs/kv_kvs"
	"github.com/watermint/toolbox/essentials/kvs/kv_storage_impl"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/quality/infra/qt_file"
	"path/filepath"
	"testing"
	"time"
)

func TestNoCache(t *testing.T) {
	nc := NoCache()
	if err := nc.Kvs().Lock(); err != nil {
		t.Error(err)
	}
	if err := nc.Kvs().PutString("mango", "apple mango"); err != nil {
		t.Error(err)
	}
	if _, err := nc.Kvs().GetString("mango"); err != kv_kvs.ErrorNotFound {
		t.Error(err)
	}
}

func TestCacheImpl_EvictIfRequired(t *testing.T) {
	qt_file.TestWithTestFolder(t, "cache", false, func(path string) {
		factory := kv_storage_impl.NewFactory(path, esl.Default())
		info := Info{
			Version:   CurrentVersion,
			Namespace: "abc",
			Name:      "123",
			Expire:    time.Now(),
			Last:      time.Now(),
		}

		cache := newCache(info, filepath.Join(path, "info.json"), path, factory, esl.Default(), 10*time.Millisecond)
		if err := cache.Open(); err != nil {
			t.Error(err)
			return
		}

		if err := cache.Kvs().PutString("mango", "apple mango"); err != nil {
			t.Error(err)
		}
		if v, err := cache.Kvs().GetString("mango"); err != nil || v != "apple mango" {
			t.Error(v, err)
		}
		cache.Close()

		time.Sleep(11 * time.Millisecond)

		if err := cache.Open(); err != nil {
			t.Error(err)
			return
		}
		if v, err := cache.Kvs().GetString("mango"); err != kv_kvs.ErrorNotFound {
			t.Error(v, err)
		}
	})
}

func TestCacheImpl_EvictEvenIfItLocked(t *testing.T) {
	qt_file.TestWithTestFolder(t, "cache", false, func(path string) {
		factory := kv_storage_impl.NewFactory(path, esl.Default())
		info := Info{
			Version:   CurrentVersion,
			Namespace: "abc",
			Name:      "123",
			Expire:    time.Now(),
			Last:      time.Now(),
		}

		cache := newCache(info, filepath.Join(path, "info.json"), path, factory, esl.Default(), 10*time.Millisecond)
		if err := cache.Open(); err != nil {
			t.Error(err)
			return
		}

		if err := cache.Kvs().PutString("mango", "apple mango"); err != nil {
			t.Error(err)
		}
		if v, err := cache.Kvs().GetString("mango"); err != nil || v != "apple mango" {
			t.Error(v, err)
		}

		time.Sleep(11 * time.Millisecond)

		cache2 := newCache(info, filepath.Join(path, "info.json"), path, factory, esl.Default(), 10*time.Millisecond)
		if err := cache2.Open(); err != nil {
			t.Error(err)
			return
		}
		if v, err := cache2.Kvs().GetString("mango"); err != kv_kvs.ErrorNotFound {
			t.Error(v, err)
		}
	})
}
