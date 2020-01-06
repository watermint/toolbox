package app_workspace

type Job interface {
	// Path for job
	Job() string

	// Job ID
	JobId() string

	// Log path for job
	Log() string

	// Test
	Test() string

	// Report path for job
	Report() string

	// Path for KVS storage
	KVS() string

	// Create or get child folder of job folder
	Descendant(name string) (path string, err error)
}
