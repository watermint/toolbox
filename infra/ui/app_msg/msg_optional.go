package app_msg

type MessageOptional interface {
	Message
	Optional() bool
}

type messageOptionalImpl struct {
	K string
	P []P
}

func (z *messageOptionalImpl) Optional() bool {
	return true
}

func (z *messageOptionalImpl) AsOptional() MessageOptional {
	return &messageOptionalImpl{
		K: z.K,
		P: z.P,
	}
}

func (z *messageOptionalImpl) With(key string, value interface{}) Message {
	np := make([]P, 0)
	np = append(np, P{key: value})
	np = append(np, z.P...)
	return &messageOptionalImpl{
		K: z.K,
		P: np,
	}
}

func (z *messageOptionalImpl) Key() string {
	return z.K
}

func (z *messageOptionalImpl) Params() []P {
	return z.P
}
