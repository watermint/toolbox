package sc_storage

import (
	"encoding/json"
	"errors"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/security/sc_obfuscate"
	"os"
)

var (
	ErrorStorageNotFound = errors.New("storage not found")
)

type Storage interface {
	// Put the value into the path. The `v` should be serializable to JSON format.
	Put(path string, v interface{}) error

	// Retrieve the value from the path.
	Get(path string, v interface{}) error
}

func NewStorage(c app_control.Control) Storage {
	return &storageImpl{c: c}
}

type storageImpl struct {
	c app_control.Control
}

func (z *storageImpl) Put(path string, v interface{}) error {
	l := z.c.Log().With(esl.String("path", path))
	l.Debug("Put obfuscated storage")

	d, err := json.Marshal(v)
	if err != nil {
		l.Debug("Unable to marshal", esl.Error(err))
		return err
	}

	b, err := sc_obfuscate.Obfuscate(l, sc_obfuscate.XapKey(), d)
	if err != nil {
		l.Debug("Unable to obfuscate", esl.Error(err))
		return err
	}

	if err := os.WriteFile(path, b, 0600); err != nil {
		l.Debug("Unable to write", esl.Error(err))
		return err
	}

	return nil
}

func (z *storageImpl) Get(path string, v interface{}) error {
	l := z.c.Log().With(esl.String("path", path))
	l.Debug("Get obfuscated storage")

	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		l.Debug("Load skipped (the file not found)")
		return ErrorStorageNotFound
	}

	b, err := os.ReadFile(path)
	if err != nil {
		l.Debug("Unable to load file", esl.Error(err))
		return err
	}

	d, err := sc_obfuscate.Deobfuscate(l, sc_obfuscate.XapKey(), b)
	if err != nil {
		l.Debug("Unable to deobfuscate sequence", esl.Error(err))
		return err
	}

	if err := json.Unmarshal(d, v); err != nil {
		l.Debug("Unable to unmarshal", esl.Error(err))
		return err
	}
	return nil
}
