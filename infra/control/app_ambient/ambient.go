package app_ambient

// TODO: Ad-hoc impl. should be implemented under essentials
type Ambient interface {
	// Suppress progress UI messages
	SuppressProgress()

	// Resume progress UI messages
	ResumeProgress()

	// Check progress suppress mode
	OnProgress(f func())
}

var (
	Current = newAmbient()
)

func newAmbient() Ambient {
	return &ambImpl{}
}

type ambImpl struct {
	suppressProgress bool
}

func (z *ambImpl) OnProgress(f func()) {
	if !z.suppressProgress {
		f()
	}
}

func (z *ambImpl) SuppressProgress() {
	z.suppressProgress = true
}

func (z *ambImpl) ResumeProgress() {
	z.suppressProgress = false
}
