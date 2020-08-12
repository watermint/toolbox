package eq_pipe_preserve

func NopFactory() Factory {
	return &nopFactory{}
}

type nopFactory struct {
}

func (z nopFactory) NewPreserver() Preserver {
	return NopPreserver()
}

func (z nopFactory) NewRestorer(sessionId string) Restorer {
	return NopRestorer()
}
