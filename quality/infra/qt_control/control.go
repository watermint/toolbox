package qt_control

type Message interface {
	NotFound(key string)
	Missing() []string
}
