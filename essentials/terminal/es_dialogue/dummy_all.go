package es_dialogue

import "strings"

func DummyAll() Dialogue {
	return &dummyAll{}
}

type dummyAll struct {
}

func (z dummyAll) AskProceed(p Prompt) {
}

func (z dummyAll) AskCont(p Prompt, v VerifyCont) (c bool) {
	p()
	return true
}

func (z dummyAll) AskText(p Prompt, v VerifyText) (t string, cancel bool) {
	p()
	return strings.Repeat("x", 8), false
}

func (z dummyAll) AskSecure(p Prompt) (t string, cancel bool) {
	p()
	return strings.Repeat("x", 8), false
}
