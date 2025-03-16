package el_ja

import "github.com/ikawaha/kagome/v2/tokenizer"

type Token struct {
	Index            int      `json:"index,omitempty"`
	ID               int      `json:"id,omitempty"`
	Class            string   `json:"class,omitempty"`
	Start            int      `json:"start,omitempty"`
	End              int      `json:"end,omitempty"`
	Surface          string   `json:"surface,omitempty"`
	Features         []string `json:"features,omitempty"`
	BaseForm         string   `json:"base_form,omitempty"`
	InflectionalForm string   `json:"inflectional_form,omitempty"`
	InflectionalType string   `json:"inflectional_type,omitempty"`
}

func NewToken(token tokenizer.Token) Token {
	var klass = "unknown"
	switch token.Class {
	case tokenizer.DUMMY:
		klass = "dummy"
	case tokenizer.KNOWN:
		klass = "known"
	case tokenizer.UNKNOWN:
		klass = "unknown"
	case tokenizer.USER:
		klass = "user"
	}

	t := Token{
		Index:    token.Index,
		ID:       token.ID,
		Class:    klass,
		Start:    token.Start,
		End:      token.End,
		Surface:  token.Surface,
		Features: token.Features(),
	}

	if v, ok := token.BaseForm(); ok {
		t.BaseForm = v
	}
	if v, ok := token.InflectionalForm(); ok {
		t.InflectionalForm = v
	}
	if v, ok := token.InflectionalType(); ok {
		t.InflectionalType = v
	}

	return t
}

func NewTokenArray(tokens []tokenizer.Token) (converted []Token) {
	converted = make([]Token, len(tokens))
	for i, token := range tokens {
		converted[i] = NewToken(token)
	}
	return
}
