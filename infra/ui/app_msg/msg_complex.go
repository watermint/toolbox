package app_msg

type MessageComplex interface {
	Message
	Messages() []Message
}

func Join(messages ...Message) MessageComplex {
	return &msgComplex{
		messages: messages,
		optional: false,
	}
}

type msgComplex struct {
	messages []Message
	optional bool
}

func (z msgComplex) Optional() bool {
	return z.optional
}

func (z msgComplex) Key() string {
	return KeyComplex
}

func (z msgComplex) Params() []P {
	return []P{
		{KeyComplexMessages: z.messages},
	}
}

func (z msgComplex) With(key string, value interface{}) Message {
	return z
}

func (z msgComplex) AsOptional() MessageOptional {
	z.optional = true
	return z
}

func (z msgComplex) Messages() []Message {
	return z.messages
}
