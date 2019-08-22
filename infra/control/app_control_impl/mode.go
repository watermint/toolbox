package app_control_impl

import (
	"github.com/watermint/toolbox/infra/app"
)

func isProduction() bool {
	return app.Hash != ""
}
