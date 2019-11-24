package mo_path

import (
	"strings"
)

type Path interface {
	// Path format for Dropbox API
	Path() string

	// Namespace ID & true, if exists. Otherwise "" and false.
	Namespace() (namespace string, exist bool)

	// File/Folder ID & true, if exists. Otherwise "" and false.
	Id() (id string, exist bool)

	// Logical part of the path. That doesn't contain namespace or file/folder id.
	// Returns `/` if the path point to root.
	LogicalPath() string

	// Child path
	ChildPath(name string) Path
}

type pathImpl struct {
	path string
}

func (z *pathImpl) String() string {
	return z.path
}

func (z *pathImpl) ChildPath(name string) Path {
	return NewPathDisplay(z.path + "/" + name)
}

func (z *pathImpl) Namespace() (namespace string, exist bool) {
	if strings.HasPrefix(z.path, "ns:") {
		p := strings.Index(z.path, "/")
		if p < 0 {
			return z.path[3:], true
		}
		return z.path[3:p], true
	}
	return "", false
}

func (z *pathImpl) Id() (id string, exist bool) {
	if strings.HasPrefix(z.path, "id:") {
		p := strings.Index(z.path, "/")
		if p < 0 {
			return z.path[3:], true
		}
		return z.path[3:p], true
	}
	return "", false
}

func (z *pathImpl) LogicalPath() string {
	if z.path == "" {
		return "/"
	}
	p := strings.Index(z.path, "/")
	if strings.HasPrefix(z.path, "ns:") {
		if p < 0 {
			return "/"
		}
		return z.path[p:]
	}
	if strings.HasPrefix(z.path, "id:") {
		if p < 0 {
			return "/"
		}
		return z.path[p:]
	}
	return z.path
}

func (z *pathImpl) Path() string {
	return z.path
}

// Create new `Path` instance.
// Windows style paths are automatically replaced for API.
func NewPath(path string) Path {
	ps1 := strings.Split(path, "\\")
	ps2 := strings.Join(ps1, "/")
	ps3 := strings.ReplaceAll(ps2, "//", "/")
	if ps3 == "/" {
		ps3 = ""
	}

	return &pathImpl{path: ps3}
}

// Create new `Path` instance. No validation & modification
func NewPathDisplay(path string) Path {
	return &pathImpl{path: path}
}
