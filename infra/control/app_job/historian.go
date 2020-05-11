package app_job

type Historian interface {
	Histories() (histories []History, err error)
}
