package es_regexp2

type Regexp interface {
	MatchSubExp(s string) (matches map[string]string, match bool)
}
