package es_regexp

import "regexp"

type Regexp interface {
	MatchSubExp(s string) (matches map[string]string, match bool)
}

func MustNew(regex string) Regexp {
	return &reImpl{re: regexp.MustCompile(regex)}
}

func New(regex string) (r Regexp, err error) {
	re, err := regexp.Compile(regex)
	if err != nil {
		return nil, err
	}
	return &reImpl{re: re}, nil
}

type reImpl struct {
	re *regexp.Regexp
}

func (z reImpl) MatchSubExp(s string) (matches map[string]string, match bool) {
	m := z.re.FindStringSubmatch(s)
	if m == nil {
		return nil, false
	}
	matches = make(map[string]string)
	subNames := z.re.SubexpNames()

	for i, sm := range subNames {
		matches[sm] = m[i]
	}
	return matches, true
}
