package app_msg

type Param func() *ParamContainer
type ParamContainer struct {
	Key   string
	Value interface{}
}

func P(key string, value interface{}) Param {
	return func() *ParamContainer {
		ph := &ParamContainer{}
		ph.Key = key
		ph.Value = value
		return ph
	}
}

type Message interface {
	Key() string
	Params() []Param
}

type messageImpl struct {
	K string
	P []Param
}

func (z *messageImpl) Key() string {
	return z.K
}

func (z *messageImpl) Params() []Param {
	return z.P
}

func M(key string, p ...Param) Message {
	return &messageImpl{
		K: key,
		P: p,
	}
}

func Raw(text string) Message {
	return &messageImpl{
		K: "raw",
		P: []Param{P("Raw", text)},
	}
}
