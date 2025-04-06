package uc_file_size

import (
	"strings"
	"sync"

	"github.com/watermint/toolbox/domain/dropbox/model/mo_file"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_file_size"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
)

// Sum up file entries. The implementation will not consider namespaces.
// The implementation is thread-safe.
type Sum interface {
	// Evaluate folder entries
	Eval(path string, entries []mo_file.Entry)

	// Retrieve results
	Each(f func(path mo_path.DropboxPath, size mo_file_size.Size))
}

func NewSum(depth int) Sum {
	return &sumImpl{
		depth: depth,
		sizes: make(map[string]mo_file_size.Size),
	}
}

type sumImpl struct {
	depth      int
	sizes      map[string]mo_file_size.Size
	sizesMutex sync.Mutex
}

// paths of entry that limited by depth
func (z *sumImpl) pathsOfEntry(path string) (paths []string) {
	components := strings.Split(path, "/")
	switch len(components) {
	case 0, 1:
		return []string{"/"}
	default:
		if components[0] == "" {
			components = components[1:]
		}
	}
	paths = make([]string, 0)
	x := min(z.depth, len(components)+1)
	for i := 0; i < x; i++ {
		paths = append(paths, "/"+strings.Join(components[:i], "/"))
	}

	// Remove duplicates using a map
	uniqueMap := make(map[string]bool)
	uniquePaths := make([]string, 0, len(paths))
	for _, p := range paths {
		if !uniqueMap[p] {
			uniqueMap[p] = true
			uniquePaths = append(uniquePaths, p)
		}
	}
	return uniquePaths
}

func (z *sumImpl) Eval(path string, entries []mo_file.Entry) {
	z.sizesMutex.Lock()
	defer z.sizesMutex.Unlock()

	s := mo_file_size.Size{Path: path}.Eval(entries)
	for _, path := range z.pathsOfEntry(path) {
		if size, ok := z.sizes[path]; ok {
			z.sizes[path] = size.Plus(path, s)
		} else {
			z.sizes[path] = s
		}
	}
}

func (z *sumImpl) Each(f func(path mo_path.DropboxPath, size mo_file_size.Size)) {
	for path, size := range z.sizes {
		f(mo_path.NewDropboxPath(path), size)
	}
}
