package qt_missingmsg

type Missing interface {
	NotFound(key string)
	Missing() []string
}

var (
	record = &misImpl{
		missing: make(map[string]bool),
	}
)

func Record() Missing {
	return record
}

type misImpl struct {
	missing map[string]bool
}

func (z *misImpl) Missing() []string {
	m := make([]string, 0)
	for k := range z.missing {
		m = append(m, k)
	}
	return m
}

func (z *misImpl) NotFound(key string) {
	z.missing[key] = true
}
