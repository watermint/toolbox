package mo_content

type ctsFile struct {
	c Content
}

func (z ctsFile) File() (c Content, found bool) {
	return z.c, true
}

func (z ctsFile) Dir() (c []Content, found bool) {
	return
}

func (z ctsFile) Symlink() (c Content, found bool) {
	return
}

func (z ctsFile) Submodule() (c Content, found bool) {
	return
}
