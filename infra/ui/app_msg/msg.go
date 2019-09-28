package app_msg

type P map[string]interface{}

type Message interface {
	Key() string
	Params() []P
}

type messageImpl struct {
	K string
	P []P
}

func (z *messageImpl) Key() string {
	return z.K
}

func (z *messageImpl) Params() []P {
	return z.P
}

func M(key string, p ...P) Message {
	return &messageImpl{
		K: key,
		P: p,
	}
}

func Raw(text string) Message {
	return &messageImpl{
		K: "raw",
		P: []P{
			{
				"Raw": text,
			},
		},
	}
}
