package es_mailaddr

import "regexp"

func EscapeSpecial(email string, alt string) string {
	// RFC 5322, atext (excl ALPHA/DIGIT)
	specials := []rune("!#$%&'*+-/=?^_`{|}~")
	emailRune := []rune(email)
	e := make([]rune, 0)

	for _, c := range emailRune {
		isSpecial := false
		for _, s := range specials {
			if c == s {
				isSpecial = true
				break
			}
		}
		if isSpecial {
			e = append(e, []rune(alt)...)
		} else {
			e = append(e, c)
		}
	}
	return string(e)
}

var (
	emailPattern = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
)

func IsEmailAddr(email string) bool {
	return emailPattern.MatchString(email)
}
