package etokenizer

// Tokenizer break a string into tokens.
type Tokenizer interface {
	Tokens(s string) []string
}
