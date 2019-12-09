package qt_missingmsg

type Message interface {
	NotFound(key string)
	Missing() []string
}
