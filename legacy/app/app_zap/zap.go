package app_zap

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"github.com/watermint/toolbox/infra/app"
	app2 "github.com/watermint/toolbox/legacy/app"
)

func Unzap(ec *app2.ExecContext) (b []byte, err error) {
	tas, err := ec.ResourceBytes("toolbox.appkeys.secret")
	if err != nil {
		return nil, err
	}
	key := []byte(app.Zap)
	zap32 := sha256.Sum256([]byte(key))
	zap := make([]byte, 32)
	copy(zap[:], zap32[:])
	block, err := aes.NewCipher(zap)
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	ns := gcm.NonceSize()
	nonce, ct := tas[:ns], tas[ns:]
	v, err := gcm.Open(nil, nonce, ct, nil)
	if err != nil {
		return nil, err
	}
	return v, nil
}
