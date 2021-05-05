package sc_zap

import (
	"crypto/sha256"
	"encoding/base32"
	"errors"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/app"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_resource"
	"github.com/watermint/toolbox/infra/security/sc_obfuscate"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func Unzap(ctl app_control.Control) (b []byte, err error) {
	tas, err := app_resource.Bundle().Keys().Bytes("toolbox.appkeys.secret")
	if err != nil {
		return nil, err
	}
	return sc_obfuscate.Deobfuscate(ctl.Log(), []byte(app.BuildInfo.Zap), tas)
}

var (
	keyEnvNames = []string{
		"HOME",
		"HOSTNAME",
		"CI_BUILD_REF",
		"CI_JOB_ID",
		"CIRCLE_BUILD_NUM",
		"CIRCLE_NODE_INDEX",
	}
)

var (
	ErrorObfuscateFailure = errors.New("obfuscation failure")
	ErrorCantFileWrite    = errors.New("cant write file")
)

func NewZap(extraSeed string) string {
	seeds := make([]byte, 0)
	seeds = strconv.AppendInt(seeds, time.Now().Unix(), 16)
	seeds = append(seeds, app.BuildId...)
	seeds = append(seeds, extraSeed...)

	for _, k := range keyEnvNames {
		if v, ok := os.LookupEnv(k); ok {
			seeds = append(seeds, k...)
			seeds = append(seeds, v...)
		}
	}
	hash := make([]byte, 32)
	sha2 := sha256.Sum256(seeds)
	copy(hash[:], sha2[:])

	b32 := base32.StdEncoding.WithPadding('_').EncodeToString(hash)
	return strings.ReplaceAll(b32, "_", "")
}

func Zap(zap string, prjRoot string, data []byte) error {
	secretPath := filepath.Join(prjRoot, "resources/keys/toolbox.appkeys.secret")
	l := esl.Default()

	b, err := sc_obfuscate.Obfuscate(l, []byte(zap), data)
	if err != nil {
		return ErrorObfuscateFailure
	}
	if err := ioutil.WriteFile(secretPath, b, 0600); err != nil {
		return ErrorCantFileWrite
	}
	return nil
}
