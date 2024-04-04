package app_license_key

import (
	"encoding/base32"
	"encoding/json"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/control/app_definitions"
	"github.com/watermint/toolbox/infra/control/app_workspace"
	"github.com/watermint/toolbox/infra/security/sc_obfuscate"
	"github.com/watermint/toolbox/resources"
	"golang.org/x/crypto/sha3"
	"os"
	"path/filepath"
	"strings"
)

type LicenseKey struct {
	Version string   `json:"version"`
	Keys    []string `json:"keys"`
}

func AvailableKeys() (licenses []string) {
	licenses = make([]string, 0)
	if relLic := resources.ReleaseLicense(app_definitions.Version.Major, app_definitions.BuildInfo.HashedSalt()); relLic != "" {
		licenses = append(licenses, relLic)
	}

	if envLic, found := os.LookupEnv(app_definitions.EnvNameToolboxLicenseKey); found {
		envLicSplit := strings.Split(envLic, ",")
		for _, el := range envLicSplit {
			licenses = append(licenses, strings.TrimSpace(el))
		}
	}

	return licenses
}

func AddKey(ws app_workspace.Workspace, key string) {
	keys := loadKeys(ws)
	keys = append(keys, key)
	saveKeys(keys, ws)
}

func keyFileKey() []byte {
	hash := sha3.Sum512([]byte("LOCAL_LICENSE:" + app_definitions.BuildInfo.LicenseSalt))
	return hash[:]
}

func licenseKeyFile() string {
	hash := sha3.Sum512([]byte("LOCAL_LICENSE:" + app_definitions.BuildInfo.LicenseSalt))
	return "license_key_" + strings.ToLower(base32.StdEncoding.EncodeToString(hash[:]))[:16] + ".key"
}

func loadKeys(ws app_workspace.Workspace) (keys []string) {
	l := esl.Default()
	keys = make([]string, 0)
	path := filepath.Join(ws.Secrets(), licenseKeyFile())
	if _, err := os.Stat(path); os.IsNotExist(err) {
		l.Debug("No license key file found", esl.String("path", path))
		return
	}
	data, err := os.ReadFile(path)
	if err != nil {
		l.Debug("Unable to read license key file", esl.Error(err))
		return
	}

	d, err := sc_obfuscate.Deobfuscate(l, keyFileKey(), data)
	if err != nil {
		l.Debug("Unable to deobfuscate license key file", esl.Error(err))
		return
	}

	lk := &LicenseKey{}
	if err := json.Unmarshal(d, lk); err != nil {
		l.Debug("Unable to unmarshal license key file", esl.Error(err))
		return
	}
	l.Debug("Loaded license keys", esl.Int("numKeys", len(lk.Keys)))
	keys = append(keys, lk.Keys...)

	return
}

func saveKeys(keys []string, ws app_workspace.Workspace) {
	l := esl.Default()
	path := filepath.Join(ws.Secrets(), licenseKeyFile())
	l.Debug("Save license keys", esl.String("path", path))
	lk := &LicenseKey{
		Version: "1",
		Keys:    keys,
	}
	d, err := json.Marshal(lk)
	if err != nil {
		l.Debug("Unable to marshal license key", esl.Error(err))
		return
	}
	e, err := sc_obfuscate.Obfuscate(l, keyFileKey(), d)
	if err != nil {
		l.Debug("Unable to obfuscate license key", esl.Error(err))
		return
	}
	if err := os.WriteFile(path, e, 0600); err != nil {
		l.Debug("Unable to write license key file", esl.Error(err))
		return
	}
}
