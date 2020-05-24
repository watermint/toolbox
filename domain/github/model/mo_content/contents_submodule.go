package mo_content

type ctsSubmodule struct {
	c Content
}

func (z ctsSubmodule) File() (c Content, found bool) {
	return
}

func (z ctsSubmodule) Dir() (c []Content, found bool) {
	return
}

func (z ctsSubmodule) Symlink() (c Content, found bool) {
	return
}

func (z ctsSubmodule) Submodule() (c Content, found bool) {
	return z.c, true
}
