package dbx_util

import "regexp"

var (
	RegexEmail = regexp.MustCompile(`^['&A-Za-z0-9._%+-]+@[A-Za-z0-9-][A-Za-z0-9.-]*\.[A-Za-z]{2,15}$`)
)
