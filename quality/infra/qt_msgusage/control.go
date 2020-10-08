package qt_msgusage

import "sort"

type Usage interface {
	NotFound(key string)
	Touch(key string)
	Missing() []string
	Used() []string
}

var (
	record = &misImpl{
		missing: make(map[string]bool),
		touch:   make(map[string]bool),
	}
)

func Record() Usage {
	return record
}

type misImpl struct {
	missing map[string]bool
	touch   map[string]bool
}

func (z *misImpl) Touch(key string) {
	z.touch[key] = true
}

func (z *misImpl) Used() []string {
	touches := make([]string, 0)
	for t := range z.touch {
		touches = append(touches, t)
	}
	return touches
}

func (z *misImpl) Missing() []string {
	m := make([]string, 0)
	for k := range z.missing {
		m = append(m, k)
	}
	sort.Strings(m)
	return m
}

func (z *misImpl) NotFound(key string) {
	z.missing[key] = true
}
