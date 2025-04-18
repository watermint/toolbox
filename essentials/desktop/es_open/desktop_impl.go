package es_open

func CurrentDesktop() Desktop {
	return &desktopImpl{}
}

type desktopImpl struct {
}

func (z desktopImpl) Open(p string) error {
	return desktopOpen(p)
}
