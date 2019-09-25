package app_worker

import "github.com/watermint/toolbox/infra/control/app_control"

type Worker func(ctl app_control.Control) error

type Queue interface {
	Enqueue(w Worker)
	Wait()
	Launch(concurrency int)
}
