package eq_pipe_preserve

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/quality/infra/qt_file"
	"strings"
	"testing"
)

func TestSimplePreserverRestorer_Single(t *testing.T) {
	type Meta struct {
		Version int    `json:"version"`
		Comment string `json:"comment"`
	}

	qt_file.TestWithTestFolder(t, "simple", false, func(path string) {
		factory := NewFactory(esl.Default(), path)
		preserver := factory.NewPreserver()
		err := preserver.Start()
		if err != nil {
			t.Error(err)
			return
		}

		if err := preserver.Add([]byte("Hello")); err != nil {
			t.Error(err)
			return
		}

		meta := &Meta{
			Version: 123,
			Comment: "Hey",
		}
		metaBin, err := json.Marshal(meta)
		if err != nil {
			t.Error(err)
			return
		}

		sessionId, err := preserver.Commit(metaBin)
		if err != nil {
			t.Error(err)
			return
		}

		restorer := factory.NewRestorer(sessionId)

		infoLoader := func(d []byte) error {
			m := &Meta{}
			if err = json.Unmarshal(d, m); err != nil {
				t.Error(err)
				return err
			}
			if m.Version != 123 || m.Comment != "Hey" {
				t.Error(m)
				return errors.New("invalid")
			}
			return nil
		}

		loader := func(d []byte) error {
			if "Hello" != string(d) {
				t.Error(d)
			}
			return nil
		}

		err = restorer.Restore(infoLoader, loader)

		if err != nil {
			t.Error(err)
		}
	})
}

func TestSimplePreserverRestorer_Bulk(t *testing.T) {
	qt_file.TestWithTestFolder(t, "simple", false, func(path string) {
		l := esl.Default()
		preserver := NewPreserver(l, path)
		err := preserver.Start()
		if err != nil {
			t.Error(err)
			return
		}

		dataCount := 100
		for i := 0; i < dataCount; i++ {
			d := fmt.Sprintf("Hello%03d", i)
			if err := preserver.Add([]byte(d)); err != nil {
				t.Error(err)
				return
			}
		}

		sessionId, err := preserver.Commit([]byte{})
		if err != nil {
			t.Error(err)
			return
		}

		restorer := NewRestorer(l, path, sessionId)

		restoreCount := 0

		infoLoader := func(d []byte) error { return nil }
		loader := func(d []byte) error {
			l.Info("Restored", esl.ByteString("Data", d))

			if !strings.HasPrefix(string(d), "Hello") {
				t.Error(d)
			}
			restoreCount++
			return nil
		}

		err = restorer.Restore(infoLoader, loader)
		if restoreCount != dataCount {
			t.Error(err)
		}

		if err != nil {
			t.Error(err)
		}
	})
}
