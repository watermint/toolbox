package app_job

import (
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"time"
)

type History interface {
	JobId() string
	RecipeName() string
	Recipe() (r rc_recipe.Spec, found bool)
	AppName() string
	AppVersion() string
	TimeStart() (t time.Time, found bool)
	TimeFinish() (t time.Time, found bool)

	// Archive job history
	Archive() (path string, err error)
}

type Historian interface {
	Histories() (histories []History)
}
