package app_msg

type P map[string]interface{}

func CreateMessage(key string, p ...P) Message {
	return &messageImpl{
		K: key,
		P: p,
	}
}

type Message interface {
	Key() string
	Params() []P
	With(key string, value interface{}) Message
	AsOptional() MessageOptional
}

type messageImpl struct {
	K string
	P []P
}

func (z *messageImpl) With(key string, value interface{}) Message {
	np := make([]P, 0)
	np = append(np, P{key: value})
	np = append(np, z.P...)
	return &messageImpl{
		K: z.K,
		P: np,
	}
}

func (z *messageImpl) AsOptional() MessageOptional {
	return &messageOptionalImpl{
		K: z.K,
		P: z.P,
	}
}

func (z *messageImpl) Key() string {
	return z.K
}

func (z *messageImpl) Params() []P {
	return z.P
}
