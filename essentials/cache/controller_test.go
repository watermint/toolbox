package cache

import (
	"github.com/watermint/toolbox/essentials/kvs/kv_kvs"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/quality/infra/qt_file"
	"testing"
)

func TestCtlImpl_Single(t *testing.T) {
	qt_file.TestWithTestFolder(t, "cache", false, func(path string) {
		cc := New(path, esl.Default())
		if err := cc.Startup(); err != nil {
			t.Error(err)
		}

		c := cc.Cache("abc", "123")
		if err := c.Kvs().PutString("mango", "愛文芒果"); err != nil {
			t.Error(err)
		}
		if v, err := c.Kvs().GetString("mango"); err != nil || v != "愛文芒果" {
			t.Error(err, v)
		}

		cc.Shutdown()
	})
}

func TestCtlImpl_MultiProcess(t *testing.T) {
	qt_file.TestWithTestFolder(t, "cache", false, func(path string) {
		cc1 := New(path, esl.Default())
		if err := cc1.Startup(); err != nil {
			t.Error(err)
		}
		c1 := cc1.Cache("abc", "123")
		if err := c1.Kvs().PutString("mango", "愛文芒果"); err != nil {
			t.Error(err)
		}
		if v, err := c1.Kvs().GetString("mango"); err != nil || v != "愛文芒果" {
			t.Error(err, v)
		}

		// 2nd process should not acquire cache
		cc2 := New(path, esl.Default())
		if err := cc2.Startup(); err != nil {
			t.Error(err)
		}
		c2 := cc1.Cache("abc", "123")
		if err := c2.Kvs().PutString("mango", "玉文芒果"); err != nil {
			t.Error(err)
		}
		if v, err := c2.Kvs().GetString("mango"); err != kv_kvs.ErrorNotFound {
			t.Error(v, err)
		}

		// 1st shutdown, 3rd open. 3rd should acquire cache
		cc1.Shutdown()
		cc3 := New(path, esl.Default())

		if err := cc3.Startup(); err != nil {
			t.Error(err)
		}
		c3 := cc3.Cache("abc", "123")
		if v, err := c3.Kvs().GetString("mango"); err != nil || v != "愛文芒果" {
			t.Error(err, v)
		}

		cc3.Shutdown()
		cc2.Shutdown()
	})
}
