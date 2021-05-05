package sc_obfuscate

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/json"
	"errors"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/app"
	"github.com/watermint/toolbox/infra/control/app_control"
	"io"
	"io/ioutil"
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

func (z *storageImpl) key() []byte {
	return []byte(app.BuildInfo.Xap + app.Name)
}

func (z *storageImpl) Put(path string, v interface{}) error {
	l := z.c.Log().With(esl.String("path", path))
	l.Debug("Put obfuscated storage")

	d, err := json.Marshal(v)
	if err != nil {
		l.Debug("Unable to marshal", esl.Error(err))
		return err
	}

	b, err := Obfuscate(l, z.key(), d)
	if err != nil {
		l.Debug("Unable to obfuscate", esl.Error(err))
		return err
	}

	if err := ioutil.WriteFile(path, b, 0600); err != nil {
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

	b, err := ioutil.ReadFile(path)
	if err != nil {
		l.Debug("Unable to load file", esl.Error(err))
		return err
	}

	d, err := Deobfuscate(l, z.key(), b)
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

func Deobfuscate(l esl.Logger, key, b []byte) (d []byte, err error) {
	l.Debug("Decrypting")
	key32 := sha256.Sum224(key)
	kb := make([]byte, 32)
	copy(kb[:], key32[:])

	bk, err := aes.NewCipher(kb)
	if err != nil {
		l.Debug("Unable to create cipher", esl.Error(err))
		return nil, err
	}
	gcm, err := cipher.NewGCM(bk)
	if err != nil {
		l.Debug("Unable to create GCM", esl.Error(err))
		return nil, err
	}
	ns := gcm.NonceSize()
	nonce, ct := b[:ns], b[ns:]
	v, err := gcm.Open(nil, nonce, ct, nil)
	if err != nil {
		l.Debug("Unable to decrypt", esl.Error(err))
		return nil, err
	}
	return v, nil
}

func Obfuscate(l esl.Logger, key, d []byte) (b []byte, err error) {
	key32 := sha256.Sum224(key)
	kb := make([]byte, 32)
	copy(kb[:], key32[:])
	bk, err := aes.NewCipher(kb)
	if err != nil {
		l.Debug("Unable to create cipher", esl.Error(err))
		return nil, err
	}
	gcm, err := cipher.NewGCM(bk)
	if err != nil {
		l.Debug("Unable to create GCM", esl.Error(err))
		return nil, err
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		l.Debug("Unable to read", esl.Error(err))
		return nil, err
	}
	return gcm.Seal(nonce, nonce, d, nil), nil
}
