package sc_obfuscate

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"github.com/watermint/toolbox/essentials/islet/eformat/ehex"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/app"
	"io"
)

func XapKey() []byte {
	return []byte(app.BuildInfo.Xap + app.Name)
}

func ZapKey() []byte {
	return []byte(app.BuildInfo.Zap)
}

func BuildStream() string {
	key0 := sha256.Sum256(XapKey())
	key1 := make([]byte, 32)
	copy(key1[:], key0[:])
	return ehex.ToHexString(key1)
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
