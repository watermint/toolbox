package qt_endtoend

import (
	"github.com/watermint/toolbox/essentials/runtime/es_env"
	"github.com/watermint/toolbox/infra/app"
	"github.com/watermint/toolbox/infra/control/app_root"
	"go.uber.org/zap"
	"testing"
)

func IsSkipEndToEndTest() bool {
	if es_env.IsEnabled(app.EnvNameEndToEndSkipTest) {
		return true
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
