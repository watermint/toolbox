package ut_mailaddr

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
