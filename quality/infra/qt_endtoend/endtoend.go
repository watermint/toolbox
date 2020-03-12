package qt_endtoend

import (
	"github.com/watermint/toolbox/infra/control/app_root"
	"go.uber.org/zap"
	"os"
	"strconv"
	"testing"
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

	// Recover: `testing: Short called before Init`
	defer func() {
		if r := recover(); r != nil {
			l := app_root.Log()
			l.Debug("Recover", zap.Any("recover", r))
		}
	}()
	if testing.Short() {
		return true
	}
	return false
}
