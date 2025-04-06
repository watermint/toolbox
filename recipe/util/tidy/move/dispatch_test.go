package move

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_definitions"
	"github.com/watermint/toolbox/quality/infra/qt_file"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
)

func TestLocal_Exec(t *testing.T) {
	qtr_endtoend.TestRecipe(t, &Dispatch{})
}

func TestNamePattern_Match(t *testing.T) {
	np, err := newNamePattern(`TBX_(\d{4})-(\d{2})-(\d{2})`)
	if err != nil {
		t.Error(err)
		return
	}
	if !np.Match("TBX_2020-04-01") {
		t.Error("does not match")
	}
	if np.Match("TBX_2020-040-01") {
		t.Error("should not match")
	}
}

func TestNamePattern_MatchValues(t *testing.T) {
	np, err := newNamePattern(`TBX_(\d{4})-(\d{2})-(\d{2})`)
	if err != nil {
		t.Error(err)
		return
	}
	mv := np.MatchValues("TBX_2020-04-01.pdf")
	if mv["M0"] != "TBX_2020-04-01" {
		t.Error("invalid")
	}
	if mv["M1"] != "2020" {
		t.Error("invalid")
	}
	if mv["M2"] != "04" {
		t.Error("invalid")
	}
	if mv["M3"] != "01" {
		t.Error("invalid")
	}
}

func TestNamePattern_Compile(t *testing.T) {
	np, err := newNamePattern(`TBX_(\d{4})-(\d{2})-(\d{2})`)
	if err != nil {
		t.Error(err)
		return
	}
	c, err := np.Compile("TBX_2020-04-01.pdf", "toolbox-{{.M1}}-{{.M2}}-{{.M3}}")
	if err != nil {
		t.Error(err)
		return
	}
	if c != "toolbox-2020-04-01" {
		t.Error("invalid")
	}
}

func TestLocalPattern_Preview(t *testing.T) {
	src, err := qt_file.MakeTestFolder("src", true)
	if err != nil {
		t.Error(err)
		return
	}
	defer func() {
		os.RemoveAll(src)
	}()
	dst, err := qt_file.MakeTestFolder("dst", false)
	if err != nil {
		t.Error(err)
		return
	}
	defer func() {
		os.RemoveAll(dst)
	}()

	qtr_endtoend.TestWithControl(t, func(ctl app_control.Control) {
		lp := &DispatchPattern{}
		lp.preview(src, dst, ctl)
	})
}

func TestLocalPattern_Move(t *testing.T) {
	src, err := qt_file.MakeTestFolder("src", false)
	if err != nil {
		t.Error(err)
		return
	}
	defer func() {
		os.RemoveAll(src)
	}()
	dst, err := qt_file.MakeTestFolder("dst", false)
	if err != nil {
		t.Error(err)
		return
	}
	defer func() {
		os.RemoveAll(dst)
	}()

	name := "TBX-2020-04-01.txt"
	srcPath := filepath.Join(src, name)
	dstPath := filepath.Join(dst, name)

	err = os.WriteFile(srcPath, []byte(app_definitions.BuildId), 0644)
	if err != nil {
		return
	}

	qtr_endtoend.TestWithControl(t, func(ctl app_control.Control) {
		lp := &DispatchPattern{}
		if err := lp.move(srcPath, dstPath, ctl); err != nil {
			t.Error(err)
			return
		}
		_, err := os.Lstat(dstPath)
		if err != nil {
			t.Error(err)
		}
	})
}

func TestLocalPattern_Exec(t *testing.T) {
	src, err := qt_file.MakeTestFolder("src", false)
	if err != nil {
		t.Error(err)
		return
	}
	defer func() {
		os.RemoveAll(src)
	}()
	dst, err := qt_file.MakeTestFolder("dst", false)
	if err != nil {
		t.Error(err)
		return
	}
	defer func() {
		os.RemoveAll(dst)
	}()

	name := "TBX-2020-04-01.txt"
	srcFile := filepath.Join(src, name)

	err = os.WriteFile(srcFile, []byte(app_definitions.BuildId), 0644)
	if err != nil {
		return
	}

	qtr_endtoend.TestWithControl(t, func(ctl app_control.Control) {
		lp := &DispatchPattern{
			Suffix:            "txt",
			SourcePath:        src,
			SourceFilePattern: `TBX-(\d{4})-(\d{2})-(\d{2})`,
			DestPathPattern:   dst + "/{{.M1}}",
			DestFilePattern:   "{{.M1}}-{{.M2}}-{{.M3}}",
		}
		if err := lp.Exec(ctl, lp.move); err != nil {
			t.Error(err)
			return
		}
		_, err := os.Lstat(filepath.Join(dst, "2020", "2020-04-01.txt"))
		if err != nil {
			t.Error(err)
		}
	})
}
