package mo_string

type OptionalString interface {
	String() string
	IsExists() bool
}

func NewOptional(str string) OptionalString {
	return &optString{
		str: str,
	}
}

type optString struct {
	str string
}

func (z *optString) String() string {
	return z.str
}

func (z *optString) IsExists() bool {
	return z.str != ""
}
