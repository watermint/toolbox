package es_dialogue

func DenyAll() Dialogue {
	return &denyImpl{}
}

type denyImpl struct {
}

func (z denyImpl) AskProceed(p Prompt) {
}

func (z denyImpl) AskCont(p Prompt, v VerifyCont) (c bool) {
	return false
}

func (z denyImpl) AskText(p Prompt, v VerifyText) (t string, cancel bool) {
	return "", true
}

func (z denyImpl) AskSecure(p Prompt) (t string, cancel bool) {
	return "", true
}
