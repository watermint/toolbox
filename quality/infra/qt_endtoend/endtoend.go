package qt_endtoend

import (
	"github.com/watermint/toolbox/infra/control/app_root"
	"go.uber.org/zap"
	"os"
	"strconv"
	"testing"
)

const (
	DeployPeer          = "deploy"
	DeployEnvToken      = "TOOLBOX_DEPLOY_TOKEN"
	EndToEndPeer        = "end_to_end_test"
	EndToEndEnvTestSkip = "TOOLBOX_SKIPENDTOENDTEST"
	EndToEndEnvToken    = "TOOLBOX_ENDTOEND_TOKEN"
	TestResourceEnv     = "TOOLBOX_TEST_RESOURCE"

	// Keys for ControlTestExtension keys
	CtlTestExtUseMock = "use_mock"
)

func IsSkipEndToEndTest() bool {
	if p, found := os.LookupEnv(EndToEndEnvTestSkip); found {
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
