package main

import (
	"crypto/sha256"
	"encoding/base32"
	"github.com/watermint/toolbox/infra/control/app_root"
	"github.com/watermint/toolbox/infra/security/sc_obfuscate"
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
	exitCantWriteSecret
	exitCantObfuscate
)

func main() {
	keyPath := "resources/toolbox.appkeys"
	secretPath := keyPath + ".secret"
	l := app_root.Log()

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

	b, err := sc_obfuscate.Obfuscate(l, []byte(key), keyContent)
	if err != nil {
		os.Exit(exitCantObfuscate)
	}
	if err := ioutil.WriteFile(secretPath, b, 0600); err != nil {
		os.Exit(exitCantWriteSecret)
	}
}
