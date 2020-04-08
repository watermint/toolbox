package app_config

import (
	"io/ioutil"
	"testing"
)

func TestConfigImpl(t *testing.T) {
	p, err := ioutil.TempDir("", "config")
	if err != nil {
		t.Error(err)
		return
	}
	cf := NewConfig(p)
	entries, err := cf.List()
	if err != nil {
		t.Error(err)
	}
	if len(entries) > 0 {
		t.Error("invalid")
	}

	if err := cf.Put("experimental", true); err != nil {
		t.Error(err)
	}
	if v, err := cf.Get("experimental"); err != nil {
		t.Error(err)
	} else {
		b := v.(bool)
		if !b {
			t.Error("invalid")
		}
	}
}
