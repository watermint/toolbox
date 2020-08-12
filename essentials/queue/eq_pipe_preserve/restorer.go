package eq_pipe_preserve

type Restorer interface {
	// Restore
	Restore(infoLoader func(info []byte) error, loader func(d []byte) error) error
}
