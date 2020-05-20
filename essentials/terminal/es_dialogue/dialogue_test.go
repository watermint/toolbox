package es_dialogue

import (
	"github.com/watermint/toolbox/essentials/io/es_stdout"
	"github.com/watermint/toolbox/essentials/log/esl"
	"io"
	"io/ioutil"
	"os"
	"testing"
)

func TestNew(t *testing.T) {
	_ = New(os.Stdout)
}

func TestYesNoCont(t *testing.T) {
	if x, y := YesNoCont("yes"); !x || !y {
		t.Error(x, y)
	}
	if x, y := YesNoCont("  YES  "); !x || !y {
		t.Error(x, y)
	}
	if x, y := YesNoCont(" no  "); x || !y {
		t.Error(x, y)
	}
	if x, y := YesNoCont("  NO  "); x || !y {
		t.Error(x, y)
	}
	if x, y := YesNoCont("  CANCEL  "); y {
		t.Error(x, y)
	}
	if x, y := YesNoCont(""); y {
		t.Error(x, y)
	}
	if x, y := YesNoCont(" "); y {
		t.Error(x, y)
	}
}

func TestNonEmptyText(t *testing.T) {
	if x, y := NonEmptyText(" "); x != "" || y {
		t.Error(x, y)
	}
	if x, y := NonEmptyText(""); x != "" || y {
		t.Error(x, y)
	}
	if x, y := NonEmptyText("test"); x != "test" || !y {
		t.Error(x, y)
	}
}

func TestDenyAll(t *testing.T) {
	d := DenyAll()
	d.AskProceed(func() {})
	p := func() {}
	vc := func(t string) (cont, valid bool) {
		return false, false
	}
	vt := func(t string) (s string, valid bool) {
		return t, true
	}
	if x := d.AskCont(p, vc); x {
		t.Error(x)
	}
	if x, y := d.AskText(p, vt); x != "" || !y {
		t.Error(x, y)
	}
	if x, y := d.AskSecure(p); x != "" || !y {
		t.Error(x, y)
	}
}

func testWithFile(t *testing.T, content string, f func(in *os.File, out io.Writer)) {
	g, err := ioutil.TempFile("", "dialogue")
	if err != nil {
		t.Error(err)
		return
	}
	p := g.Name()
	defer func() {
		_ = g.Close()
		_ = os.Remove(p)
	}()

	if _, err := g.WriteString(content); err != nil {
		t.Error(err)
		return
	}
	if err := g.Sync(); err != nil {
		t.Error(err)
	}
	if _, err := g.Seek(0, io.SeekStart); err != nil {
		t.Error(err)
	}

	f(g, es_stdout.NewTestOut())
}

func TestDlgImpl_AskProceed(t *testing.T) {
	l := esl.Default()
	testWithFile(t, "yes", func(in *os.File, out io.Writer) {
		d := &dlgImpl{in: in, wr: out}
		d.AskProceed(func() {
			l.Debug("ask proceed")
		})
	})
}

func TestDlgImpl_AskCont(t *testing.T) {
	l := esl.Default()
	p := func() {
		l.Debug("ask cont")
	}
	v := func(t string) (bool, bool) {
		return t == "yes", true
	}
	testWithFile(t, "yes", func(in *os.File, out io.Writer) {
		d := &dlgImpl{in: in, wr: out}
		if x := d.AskCont(p, v); !x {
			t.Error(x)
		}
	})
	testWithFile(t, "no", func(in *os.File, out io.Writer) {
		d := &dlgImpl{in: in, wr: out}
		if x := d.AskCont(p, v); x {
			t.Error(x)
		}
	})
}

func TestDlgImpl_AskText(t *testing.T) {
	l := esl.Default()
	p := func() {
		l.Debug("ask text")
	}
	v := func(t string) (string, bool) {
		return t, t != ""
	}
	testWithFile(t, "hey", func(in *os.File, out io.Writer) {
		d := &dlgImpl{in: in, wr: out}
		if x, c := d.AskText(p, v); x != "hey" || c {
			t.Error(x, c)
		}
	})
}

func TestDlgImpl_AskSecure(t *testing.T) {
	l := esl.Default()
	testWithFile(t, "yes", func(in *os.File, out io.Writer) {
		d := &dlgImpl{in: in, wr: out}
		_, _ = d.AskSecure(func() {
			l.Debug("ask secure")
		})
	})
}
