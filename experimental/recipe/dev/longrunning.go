package dev

import (
	"github.com/watermint/toolbox/experimental/app_kitchen"
	"github.com/watermint/toolbox/experimental/app_vo"
	"go.uber.org/zap"
)

type LongRunning struct {
}

func (*LongRunning) Hidden() {
}

func (*LongRunning) Requirement() app_vo.ValueObject {
	return &app_vo.EmptyValueObject{}
}

func (*LongRunning) Exec(k app_kitchen.Kitchen) error {
	for i := 0; i < 10000; i++ {
		for j := 0; j < 10000; j++ {
			k.Log().Debug("LongRunner", zap.Int("i", i), zap.Int("j", j))
		}
	}
	return nil
}
