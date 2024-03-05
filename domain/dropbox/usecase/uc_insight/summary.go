package uc_insight

type Summarizer interface {
	// Summarize summarize the scan result
	Summarize() (err error)
}
