package app_job

import (
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"time"
)

type History interface {
	JobId() string
	JobPath() string
	RecipeName() string
	Recipe() (r rc_recipe.Spec, found bool)
	AppName() string
	AppVersion() string
	TimeStart() (t time.Time, found bool)
	TimeFinish() (t time.Time, found bool)

	// Archive this job data
	Archive() (path string, err error)

	// Delete this job data
	Delete() error

	// Log files. logs are guaranteed sorted by their file names.
	Logs() (logs []LogFile, err error)
}
