package resources

import (
	"encoding/json"
	"os"
)

const (
	// EnvNameToolboxLicenseSalt Env variable name for license salt
	EnvNameToolboxLicenseSalt = "TOOLBOX_LICENSE_SALT"
)

type LicenseResource struct {
	Salt string `json:"salt"`
}

func LicenseSalt() string {
	saltFromResource := func() string {
		licData, err := CurrentBundle.Keys().Bytes("license.json")
		if err != nil {
			return ""
		}
		lic := &LicenseResource{}
		if err := json.Unmarshal(licData, lic); err != nil {
			return ""
		}
		return lic.Salt
	}

	saltFromEnv := func() string {
		return os.Getenv(EnvNameToolboxLicenseSalt)
	}

	if salt := saltFromResource(); salt != "" {
		return salt
	}
	return saltFromEnv()
}
