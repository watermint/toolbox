package rc_worker

type Worker interface {
	Exec() error
}

type LegacyQueue interface {
	Enqueue(w Worker)
	Wait() error
}
