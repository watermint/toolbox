package dev

import (
	"fmt"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_log"
	"github.com/watermint/toolbox/infra/recpie/app_kitchen"
	"github.com/watermint/toolbox/infra/recpie/app_vo"
	"go.uber.org/zap"
	"sync"
)

type Async struct {
}

func (z *Async) Hidden() {
}

func (z *Async) Requirement() app_vo.ValueObject {
	return &app_vo.EmptyValueObject{}
}

func (z *Async) Exec(k app_kitchen.Kitchen) error {
	lc := make(chan app_log.Log)
	zl := k.Log().WithOptions(zap.AddCallerSkip(1))

	// Logger
	go func() {
		for l := range lc {
			r := l()
			r.Level().Out(zl, r)
		}
	}()

	wg := sync.WaitGroup{}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			for j := 0; j < 1000; j++ {
				lc <- app_log.Debug(fmt.Sprintf("[Worker %d] Job %d", id, j))
			}
			wg.Done()
		}(i)
	}

	wg.Wait()

	return nil
}

func (z *Async) Test(c app_control.Control) error {
	return z.Exec(app_kitchen.NewKitchen(c, &app_vo.EmptyValueObject{}))
}
