package efscommon

import (
	efs_deprecated2 "github.com/watermint/toolbox/essentials/file/efs_deprecated"
	"strings"
)

type Path interface {
	efs_deprecated2.Path
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

func NewPath(fs efs_deprecated2.FileSystem, ns efs_deprecated2.Namespace, pathElements []string) Path {
	return &pathImpl{
		pathElements: pathElements,
		namespace:    ns,
		fs:           fs,
	}
}

type pathImpl struct {
	pathElements []string
	namespace    efs_deprecated2.Namespace
	fs           efs_deprecated2.FileSystem
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

func (z pathImpl) FileSystem() efs_deprecated2.FileSystem {
	return z.fs
}

func (z pathImpl) Parent() efs_deprecated2.Path {
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

func (z pathImpl) Root() efs_deprecated2.Path {
	z.pathElements = rootPathElements
	return z
}

func (z pathImpl) IsRoot() bool {
	return len(z.pathElements) < 1
}

func (z pathImpl) Namespace() efs_deprecated2.Namespace {
	return z.namespace
}

func (z pathImpl) Child(name ...string) (efs_deprecated2.Path, efs_deprecated2.ChildOutcome) {
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
