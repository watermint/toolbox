package qt_secure

import (
	"os"
	"strconv"
)

const EndToEndEnvSecure = "TOOLBOX_SECURE"

// Do not output detail logs when this flag marked as true
func IsSecureEndToEndTest() bool {
	if p, found := os.LookupEnv(EndToEndEnvSecure); found {
		if b, _ := strconv.ParseBool(p); b {
			return b
		}
	}
	return false
}
