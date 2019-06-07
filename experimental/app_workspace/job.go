package app_workspace

type Job interface {
	// Path for job
	Job() string

	// Job ID
	JobId() string

	// Log path for job
	Log() string

	// Create or get child folder of job folder
	Descendant(name string) (path string, err error)
}
