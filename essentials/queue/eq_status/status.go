package eq_status

type Watcher interface {
	OnDequeue()
	OnEnqueue()
}
