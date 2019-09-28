package app_worker

type Worker interface {
	Exec() error
}

type Queue interface {
	Enqueue(w Worker)
	Wait()
}
