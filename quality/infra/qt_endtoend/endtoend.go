package qt_endtoend

import (
	"os"
	"strconv"
)

const (
	EndToEndPeer        = "end_to_end_test"
	EndToEndTestSkipEnv = "TOOLBOX_SKIPENDTOENDTEST"

	// Keys for ControlTestExtension keys
	CtlTestExtUseMock = "use_mock"
)

func IsSkipEndToEndTest() bool {
	if p, found := os.LookupEnv(EndToEndTestSkipEnv); found {
		if b, _ := strconv.ParseBool(p); b {
			return true
		}
	}
	return false
}
