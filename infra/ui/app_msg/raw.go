package app_msg

func Raw(text interface{}) Message {
	return &messageImpl{
		K: KeyRaw,
		P: []P{
			{KeyParamRaw: text},
		},
	}
}
