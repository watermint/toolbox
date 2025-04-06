package efscommon

import (
	"strings"

	"github.com/watermint/toolbox/essentials/file/efs_alpha"
)

type NameOpt func(opts nameOpts) nameOpts

type nameOpts struct {
	invalidChars           []rune
	reservedNames          []string
	reservedNameIgnoreCase []string
	maxLength              int
}

func (z nameOpts) Apply(opts []NameOpt) nameOpts {
	switch len(opts) {
	case 0:
		return z
	case 1:
		return opts[0](z)
	default:
		return opts[0](z).Apply(opts[1:])
	}
}

func (z nameOpts) Accept(name string) error {
	if l := len(name); 0 < z.maxLength && z.maxLength < l {
		return NewNameOutcomeNameTooLong(l, z.maxLength)
	}
	for _, r := range name {
		for _, ic := range z.invalidChars {
			if r == ic {
				return NewNameOutcomeInvalidChar(string(r))
			}
		}
	}
	for _, r := range z.reservedNames {
		if r == name {
			return NewNameOutcomeNameReserved(r)
		}
	}
	nameLower := strings.ToLower(name)
	for _, r := range z.reservedNameIgnoreCase {
		if r == nameLower {
			return NewNameOutcomeNameReserved(r)
		}
	}
	return nil
}

func DefineNameInvalidChars(chars ...rune) NameOpt {
	return func(opts nameOpts) nameOpts {
		opts.invalidChars = chars
		return opts
	}
}

func DefineNameReservedNames(names ...string) NameOpt {
	return func(opts nameOpts) nameOpts {
		opts.reservedNames = names
		return opts
	}
}

func DefineNameReservedNameIgnoreCase(names ...string) NameOpt {
	return func(opts nameOpts) nameOpts {
		opts.reservedNameIgnoreCase = names
		return opts
	}
}

// DefineMaxNameLength define maximum name length (inclusive)
func DefineMaxNameLength(max int) NameOpt {
	return func(opts nameOpts) nameOpts {
		opts.maxLength = max
		return opts
	}
}

func NewName(opts ...NameOpt) efs_alpha.Name {
	return &nameImpl{
		opts: nameOpts{}.Apply(opts),
	}
}

type nameImpl struct {
	opts nameOpts
}

func (z nameImpl) Accept(name string) error {
	return z.opts.Accept(name)
}
