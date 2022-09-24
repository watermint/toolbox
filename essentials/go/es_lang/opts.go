package es_lang

func ApplyOpts[T any](v T, opts []func(v T) T) T {
	switch len(opts) {
	case 0:
		return v
	case 1:
		return opts[0](v)
	default:
		return ApplyOpts(opts[0](v), opts[1:])
	}
}
