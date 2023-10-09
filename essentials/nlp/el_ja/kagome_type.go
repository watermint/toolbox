package el_ja

import "github.com/ikawaha/kagome/v2/tokenizer"

type Token struct {
	Index   int    `json:"index,omitempty"`
	ID      int    `json:"id,omitempty"`
	Class   string `json:"class,omitempty"`
	Start   int    `json:"start,omitempty"`
	End     int    `json:"end,omitempty"`
	Surface string `json:"surface,omitempty"`
}

func NewToken(token tokenizer.Token) Token {
	var klass string = "unknown"
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
	return Token{
		Index:   token.Index,
		ID:      token.ID,
		Class:   klass,
		Start:   token.Start,
		End:     token.End,
		Surface: token.Surface,
	}
}

func NewTokenArray(tokens []tokenizer.Token) (converted []Token) {
	converted = make([]Token, len(tokens))
	for i, token := range tokens {
		converted[i] = NewToken(token)
	}
	return
}
