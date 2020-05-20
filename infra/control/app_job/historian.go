package app_job

type Historian interface {
	// Histories, that guarantee sorted by job id.
	Histories() (histories []History, err error)
}
