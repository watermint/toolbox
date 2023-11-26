package eregexp

type Regexp interface {
	MatchSubExp(s string) (matches map[string]string, match bool)
}
