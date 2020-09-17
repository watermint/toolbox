package mo_filter

import (
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"net/mail"
	"strings"
)

func NewEmailFilter() FilterOpt {
	return &emailFilterOpt{}
}

type emailFilterOpt struct {
	email string
}

func (z *emailFilterOpt) Accept(v interface{}) bool {
	return ExpectString(v, func(s string) bool {
		if strings.ToLower(s) == strings.ToLower(z.email) {
			return true
		}
		addr, err := mail.ParseAddress(s)
		if err != nil {
			return false
		}
		return addr.Name == z.email || strings.ToLower(addr.Address) == strings.ToLower(z.email)
	})
}

func (z *emailFilterOpt) Bind() interface{} {
	return &z.email
}

func (z *emailFilterOpt) NameSuffix() string {
	return "Email"
}

func (z *emailFilterOpt) Desc() app_msg.Message {
	return MFilter.DescFilterEmail
}

func (z *emailFilterOpt) Enabled() bool {
	return z.email != ""
}
