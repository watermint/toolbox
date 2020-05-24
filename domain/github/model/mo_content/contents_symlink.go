package mo_content

type ctsSymlink struct {
	c Content
}

func (z ctsSymlink) File() (c Content, found bool) {
	return
}

func (z ctsSymlink) Dir() (c []Content, found bool) {
	return
}

func (z ctsSymlink) Symlink() (c Content, found bool) {
	return z.c, true
}

func (z ctsSymlink) Submodule() (c Content, found bool) {
	return
}
