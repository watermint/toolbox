package mo_string

type OptionalString interface {
	String
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

func (z *optString) Value() string {
	return z.str
}

func (z *optString) IsExists() bool {
	return z.str != ""
}
