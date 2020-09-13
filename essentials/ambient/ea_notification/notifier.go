package ea_notification

type Notifier interface {
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
	suppress bool
}

func (z *repoImpl) Suppress() {
	z.suppress = true
}

func (z *repoImpl) Resume() {
	z.suppress = false
}

func (z *repoImpl) OnProgress(f func()) {
	if !z.suppress {
		f()
	}
}
