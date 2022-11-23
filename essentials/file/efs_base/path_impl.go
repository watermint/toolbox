package efs_base

type pathImpl struct {
	namespace Namespace
	elements  []string
	spec      PathSpec
}

func (z pathImpl) String() string {
	//TODO implement me
	panic("implement me")
}

func (z pathImpl) Namespace() Namespace {
	return z.namespace
}

func (z pathImpl) Names() []string {
	return z.elements
}

func (z pathImpl) Basename() string {
	if n := len(z.elements); n < 1 {
		return ""
	} else {
		return z.elements[n-1]
	}
}

func (z pathImpl) Extname() string {
	//TODO implement me
	panic("implement me")
}

func (z pathImpl) Root() Path {
	z.elements = []string{}
	return z
}

func (z pathImpl) IsRoot() bool {
	return len(z.elements) == 0
}

func (z pathImpl) HasPrefix(prefix string) bool {
	//TODO implement me
	panic("implement me")
}

func (z pathImpl) HasSuffix(suffix string) bool {
	//TODO implement me
	panic("implement me")
}

func (z pathImpl) Depth() int {
	return len(z.elements)
}

func (z pathImpl) Relative(other Path) RelativePath {
	//TODO implement me
	panic("implement me")
}

func (z pathImpl) Resolve(relative RelativePath) (Path, PathError) {
	//TODO implement me
	panic("implement me")
}

func (z pathImpl) Sibling(other string) (Path, PathError) {
	if pe := z.spec.AssertName(other); pe != nil {
		return nil, pe
	}
	parent := z.Parent()
	sp, pe := parent.Child(other)
	if pe != nil {
		return nil, pe
	}
	if ppe := z.spec.AssertPath(z.String()); ppe != nil {
		return nil, ppe
	}
	return sp, nil
}

func (z pathImpl) Parent() Path {
	if n := len(z.elements); n < 1 {
		return z
	} else {
		z.elements = z.elements[:n-1]
		return z
	}
}

func (z pathImpl) Child(name ...string) (Path, PathError) {
	//TODO implement me
	panic("implement me")
}
