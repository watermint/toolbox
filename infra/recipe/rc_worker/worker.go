package rc_worker

type Worker interface {
	Exec() error
}

type Queue interface {
	Enqueue(w Worker)
	Wait() error
}
