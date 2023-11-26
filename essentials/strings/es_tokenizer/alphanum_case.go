package es_tokenizer

import "unicode"

// AlphaNumCaseTokenizer returns alpha-numeric tokens. This tokenizer ignores characters except alpha-numeric.
// This tokenizer splits token on case change. For example,
// "Powered by GoLang version1.17." is tokenized to "Powered", "by", "Go", "Lang", "version1", and "17".
func AlphaNumCaseTokenizer() Tokenizer {
	return alNumCaseTokenizer{}
}

type alNumCaseTokenizer struct {
}

func (z alNumCaseTokenizer) Tokens(s string) []string {
	runes := []rune(s)
	tokens := make([]string, 0)

	token := make([]rune, 0)
	lastLower := false

	flush := func() {
		if len(token) > 0 {
			tokens = append(tokens, string(token))
			token = make([]rune, 0)
			lastLower = false
		}
	}
	for _, r := range runes {
		switch {
		case unicode.IsLower(r):
			token = append(token, r)
			lastLower = true
		case unicode.IsUpper(r):
			if lastLower {
				flush()
			}
			token = append(token, r)
		case unicode.IsNumber(r):
			token = append(token, r)
		default:
			flush()
		}
	}
	flush()

	return tokens
}
