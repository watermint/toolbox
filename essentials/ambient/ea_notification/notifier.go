package ea_notification

type Notifier interface {
	// Force suppress for test
	SuppressForce()

	// Suppress progress notification
	Suppress()

	// Resume progress notification
	Resume()

	// Notify report on progress
	OnProgress(f func())
}

var (
	global = &repoImpl{
		suppress: false,
	}
)

func Global() Notifier {
	return global
}

type repoImpl struct {
	suppress      bool
	suppressForce bool
}

func (z *repoImpl) SuppressForce() {
	z.suppressForce = true
}

func (z *repoImpl) Suppress() {
	if !z.suppressForce {
		z.suppress = true
	}
}

func (z *repoImpl) Resume() {
	z.suppress = false
}

func (z *repoImpl) OnProgress(f func()) {
	if !z.suppress && !z.suppressForce {
		f()
	}
}
