package oper_member

import "go.uber.org/zap"

type List struct {
	Logger *zap.Logger
}

func (z *List) Exec() {
	z.Logger.Info("Member List")
}
