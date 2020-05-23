package mo_content

type ctsDir struct {
	entries []Content
}

func (z ctsDir) File() (c Content, found bool) {
	return
}

func (z ctsDir) Dir() (c []Content, found bool) {
	return z.entries, true
}

func (z ctsDir) Symlink() (c Content, found bool) {
	return
}

func (z ctsDir) Submodule() (c Content, found bool) {
	return
}
