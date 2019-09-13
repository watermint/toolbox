package dev

import (
	"fmt"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recpie/app_kitchen"
	"github.com/watermint/toolbox/infra/recpie/app_vo"
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
	type Report struct {
		Worker int `json:"worker"`
		Job    int `json:"job"`
	}

	wg := sync.WaitGroup{}
	rep, err := k.Report("async", &Report{})
	if err != nil {
		return err
	}
	defer rep.Close()

	concurrency := 8

	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go func(id int) {
			for j := 0; j < 10000; j++ {
				k.Log().Debug(fmt.Sprintf("[Worker %d] Job %d", id, j))
				rep.Row(&Report{i, j})
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
