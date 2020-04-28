package es_filehash

import (
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/quality/infra/qt_file"
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"testing"
)

func TestHashImpl_MD5(t *testing.T) {
	qt_recipe.TestWithControl(t, func(ctl app_control.Control) {
		f, err := qt_file.MakeTestFile("md5", "hello")
		if err != nil {
			t.Error(err)
			return
		}

		h := NewHash(ctl.Log())
		d, err := h.MD5(f)
		if err != nil {
			t.Error(err)
			return
		}
		if d != "5d41402abc4b2a76b9719d911017c592" {
			t.Error(d)
		}
	})
}

func TestHashImpl_SHA256(t *testing.T) {
	qt_recipe.TestWithControl(t, func(ctl app_control.Control) {
		f, err := qt_file.MakeTestFile("sha256", "hello")
		if err != nil {
			t.Error(err)
			return
		}

		h := NewHash(ctl.Log())
		d, err := h.SHA256(f)
		if err != nil {
			t.Error(err)
			return
		}
		if d != "2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824" {
			t.Error(d)
		}
	})
}
