package oper_member

import (
	"github.com/watermint/toolbox/poc/oper"
)

type Member struct {
}

func (Member) Operations() []oper.Operation {
	return []oper.Operation{
		&Invite{},
		&List{},
	}
}
