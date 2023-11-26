package efscommon

import (
	"github.com/watermint/toolbox/essentials/islet/efs"
	"strings"
)

type Path interface {
	efs.Path
}

const (
	pathSeparator = "/"
)

var (
	rootPathElements = make([]string, 0)
)

func joinPathElements(elements []string) string {
	if len(elements) < 1 {
		return pathSeparator
	}
	return pathSeparator + strings.Join(elements, pathSeparator)
}

func NewPath(fs efs.FileSystem, ns efs.Namespace, pathElements []string) Path {
	return &pathImpl{
		pathElements: pathElements,
		namespace:    ns,
		fs:           fs,
	}
}

type pathImpl struct {
	pathElements []string
	namespace    efs.Namespace
	fs           efs.FileSystem
}

func (z pathImpl) String() string {
	return joinPathElements(z.pathElements)
}

func (z pathImpl) Identity() string {
	nsId := z.namespace.Identity()
	if nsId == "" {
		return joinPathElements(z.pathElements)
	}
	return nsId + nsIdentitySeparator + joinPathElements(z.pathElements)
}

func (z pathImpl) FileSystem() efs.FileSystem {
	return z.fs
}

func (z pathImpl) Parent() efs.Path {
	if len(z.pathElements) < 1 {
		return z
	}
	z.pathElements = z.pathElements[:len(z.pathElements)-2]
	return z
}

func (z pathImpl) Basename() string {
	if len(z.pathElements) < 1 {
		return ""
	} else {
		return z.pathElements[len(z.pathElements)-1]
	}
}

func (z pathImpl) Extname() string {
	base := z.Basename()
	pos := strings.LastIndex(base, ".")
	if pos < 0 {
		return ""
	}
	return base[pos:]
}

func (z pathImpl) Root() efs.Path {
	z.pathElements = rootPathElements
	return z
}

func (z pathImpl) IsRoot() bool {
	return len(z.pathElements) < 1
}

func (z pathImpl) Namespace() efs.Namespace {
	return z.namespace
}

func (z pathImpl) Child(name ...string) (efs.Path, efs.ChildOutcome) {
	if len(name) < 1 {
		return z, NewChildOutcomeSuccess()
	}
	elements := make([]string, len(z.pathElements)+len(name))
	existingOffset := len(z.pathElements)
	for i, n := range name {
		noc := z.fs.NameRule().Accept(n)
		switch {
		case noc.IsOk():
			elements[existingOffset+i] = n
		default:
			return nil, NewChildOutcomeByNameOutcome(noc)
		}
	}
	copy(elements[0:existingOffset-1], z.pathElements)

	z.pathElements = elements
	return z, NewChildOutcomeSuccess()
}
