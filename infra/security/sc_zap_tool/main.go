package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base32"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

var (
	keyEnvNames = []string{
		"HOME",
		"HOSTNAME",
		"CI_BUILD_REF",
		"CI_JOB_ID",
	}
)

func getKey() string {
	seed := time.Now().String()
	envs := os.Environ()
	env := make(map[string]string)
	for _, e := range envs {
		i := strings.Index(e, "=")
		if i < 0 {
			// skip
			continue
		}
		k := e[:i]
		v := e[i+1:]
		env[k] = v
	}
	for _, k := range keyEnvNames {
		if v, e := env[k]; e {
			seed += v
		}
	}
	hash := make([]byte, 32)
	sha2 := sha256.Sum256([]byte(seed))
	copy(hash[:], sha2[:])

	b32 := base32.StdEncoding.WithPadding('_').EncodeToString(hash)
	return strings.ReplaceAll(b32, "_", "")
}

const (
	exitCantOpenKey = iota + 1
	exitCantReadKey
	exitCantWriteZap
	exitCantCreateBlock
	exitCantCreateGCM
	exitCantPrepareNonce
	exitCantWriteSecret
)

func main() {
	keyPath := "resources/toolbox.appkeys"
	secretPath := keyPath + ".secret"

	keyFile, err := os.Open(keyPath)
	if err != nil {
		os.Exit(exitCantOpenKey)
	}

	defer keyFile.Close()
	keyContent, err := ioutil.ReadAll(keyFile)
	if err != nil {
		os.Exit(exitCantReadKey)
	}

	key := getKey()
	if err := ioutil.WriteFile("/tmp/toolbox.zap", []byte(key), 0600); err != nil {
		os.Exit(exitCantWriteZap)
	}

	zap32 := sha256.Sum256([]byte(key))
	zap := make([]byte, 32)
	copy(zap[:], zap32[:])

	block, err := aes.NewCipher([]byte(zap))
	if err != nil {
		os.Exit(exitCantCreateBlock)
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		os.Exit(exitCantCreateGCM)
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		os.Exit(exitCantPrepareNonce)
	}
	sealed := gcm.Seal(nonce, nonce, keyContent, nil)
	if err := ioutil.WriteFile(secretPath, sealed, 0600); err != nil {
		os.Exit(exitCantWriteSecret)
	}
}
