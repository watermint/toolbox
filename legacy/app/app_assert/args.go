package app_assert

import (
	"errors"
	app2 "github.com/watermint/toolbox/legacy/app"
	"github.com/watermint/toolbox/legacy/app/app_ui"
)

type Args interface {
	Validate(args []string, opts ...ArgAssertOption) (msg app_ui.UIMessage, err error)
}

type ArgAssertOption func(o *argOpts) *argOpts

func NoArgs() ArgAssertOption {
	return func(o *argOpts) *argOpts {
		b := true
		o.noArgs = &b
		return o
	}
}

func MinNArgs(n int) ArgAssertOption {
	return func(o *argOpts) *argOpts {
		o.minArgs = &n
		return o
	}
}

func MaxNArgs(n int) ArgAssertOption {
	return func(o *argOpts) *argOpts {
		o.maxArgs = &n
		return o
	}
}

func AssertArgs(ec *app2.ExecContext, args []string, opts ...ArgAssertOption) (msg app_ui.UIMessage, err error) {
	v := &argsImpl{
		ec: ec,
	}
	return v.Validate(args, opts...)
}

type argOpts struct {
	noArgs  *bool
	minArgs *int
	maxArgs *int
}

type argsImpl struct {
	ec *app2.ExecContext
}

func (z *argsImpl) validateNoArgs(args []string, noArgs bool) (msg app_ui.UIMessage, err error) {
	if noArgs && len(args) > 0 {
		return z.ec.Msg("app.common.assert.err.no_args"), errors.New("violation: no argument")
	}
	return z.ec.Msg(app2.MsgNoError), nil
}

func (z *argsImpl) validateMinNArgs(args []string, minArgs int) (msg app_ui.UIMessage, err error) {
	if minArgs > len(args) {
		msg = z.ec.Msg("app.common.assert.err.min_args").WithData(struct {
			N int
		}{
			N: minArgs,
		})
		return msg, errors.New("violation: min n arguments")
	}
	return z.ec.Msg(app2.MsgNoError), nil
}

func (z *argsImpl) validateMaxNArgs(args []string, maxArgs int) (msg app_ui.UIMessage, err error) {
	if maxArgs < len(args) {
		msg = z.ec.Msg("app.common.assert.err.max_args").WithData(struct {
			N int
		}{
			N: maxArgs,
		})
		return msg, errors.New("violation: max n arguments")
	}
	return z.ec.Msg(app2.MsgNoError), nil
}

func (z *argsImpl) Validate(args []string, opts ...ArgAssertOption) (msg app_ui.UIMessage, err error) {
	ao := &argOpts{}
	for _, o := range opts {
		o(ao)
	}

	if ao.noArgs != nil {
		return z.validateNoArgs(args, *ao.noArgs)
	}
	if ao.minArgs != nil {
		return z.validateMinNArgs(args, *ao.minArgs)
	}
	if ao.maxArgs != nil {
		return z.validateMaxNArgs(args, *ao.maxArgs)
	}

	return z.ec.Msg(app2.MsgNoError), nil
}
