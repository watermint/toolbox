package app_worker

type Worker func() error

type Queue interface {
	Enqueue(w Worker)
	Wait()
}
