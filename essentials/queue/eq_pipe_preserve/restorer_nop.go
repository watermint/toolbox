package eq_pipe_preserve

func NopRestorer() Restorer {
	return &nopRestorer{}
}

type nopRestorer struct {
}

func (z nopRestorer) Restore(infoLoader func(info []byte) error, loader func(d []byte) error) error {
	return nil
}
