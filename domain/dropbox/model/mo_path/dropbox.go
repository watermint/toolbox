package mo_path

import (
	"path/filepath"
	"strings"
)

type DropboxPath interface {
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
	ChildPath(elem ...string) DropboxPath

	IsValid() bool
}

type dropboxPathImpl struct {
	ns        string
	id        string
	path      string
	pathEmpty bool
}

func (z *dropboxPathImpl) IsValid() bool {
	return !z.pathEmpty
}

func (z *dropboxPathImpl) Value() string {
	switch {
	case z.ns != "":
		// root of the namespace
		if z.path == "" {
			return "ns:" + z.ns
		}
		// z.path always starts with '/' if it's not empty
		return "ns:" + z.ns + z.path

	case z.id != "":
		// root of the folder id
		if z.path == "" {
			return "id:" + z.id
		}
		// z.path always starts with '/' if it's not empty
		return "id:" + z.id + z.path

	default:
		if z.path == "/" {
			return ""
		} else {
			return z.path
		}
	}
}

func (z *dropboxPathImpl) ChildPath(elem ...string) DropboxPath {
	a := make([]string, 0)
	a = append(a, z.path)
	a = append(a, elem...)
	path := filepath.ToSlash(filepath.Join(a...))
	if path != "" && !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	return &dropboxPathImpl{
		ns:   z.ns,
		id:   z.id,
		path: path,
	}
}

func (z *dropboxPathImpl) Namespace() (namespace string, exist bool) {
	return z.ns, z.ns != ""
}

func (z *dropboxPathImpl) Id() (id string, exist bool) {
	return z.id, z.id != ""
}

func (z *dropboxPathImpl) LogicalPath() string {
	if z.path == "" {
		return "/"
	}
	return z.path
}

func (z *dropboxPathImpl) Path() string {
	return z.Value()
}

// Create new `Path` instance.
// Windows style paths are automatically replaced for API.
func NewDropboxPath(path string) DropboxPath {
	ps1 := strings.Split(path, "\\")
	ps2 := strings.Join(ps1, "/")
	ps3 := strings.ReplaceAll(ps2, "//", "/")
	emp := false
	if path == "" {
		emp = true
	}
	if ps3 == "/" {
		ps3 = ""
	}

	ns := ""
	id := ""
	pe := ps3

	if strings.HasPrefix(ps3, "ns:") {
		p := strings.Index(ps3, "/")
		if p < 0 {
			ns = ps3[3:]
			pe = ""
		} else {
			ns = ps3[3:p]
			pe = ps3[p:]
		}
	}
	if strings.HasPrefix(ps3, "id:") {
		p := strings.Index(ps3, "/")
		if p < 0 {
			id = ps3[3:]
			pe = ""
		} else {
			id = ps3[3:p]
			pe = ps3[p:]
		}
	}

	return &dropboxPathImpl{
		ns:        ns,
		id:        id,
		path:      pe,
		pathEmpty: emp,
	}
}

// Create new `Path` instance. No validation & modification
func NewPathDisplay(path string) DropboxPath {
	return &dropboxPathImpl{path: path}
}
