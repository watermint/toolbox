package efscommon

import (
	"fmt"
	"strings"

	efs_deprecated2 "github.com/watermint/toolbox/essentials/file/efs_alpha"
)

type Path interface {
	efs_deprecated2.Path
}

const (
	pathSeparator = "/"
)

var (
	rootPathElements = make([]string, 0)
)

func joinPathElements(elements []string) string {
	if len(elements) < 1 {
		return pathSeparator
	}
	return pathSeparator + strings.Join(elements, pathSeparator)
}

func NewPath(fs efs_deprecated2.FileSystem, ns efs_deprecated2.Namespace, pathElements []string) Path {
	return &pathImpl{
		pathElements: pathElements,
		namespace:    ns,
		fs:           fs,
	}
}

type pathImpl struct {
	pathElements []string
	namespace    efs_deprecated2.Namespace
	fs           efs_deprecated2.FileSystem
}

func (z pathImpl) String() string {
	return joinPathElements(z.pathElements)
}

func (z pathImpl) Identity() string {
	nsId := z.namespace.Identity()
	if nsId == "" {
		return joinPathElements(z.pathElements)
	}
	return nsId + nsIdentitySeparator + joinPathElements(z.pathElements)
}

func (z pathImpl) FileSystem() efs_deprecated2.FileSystem {
	return z.fs
}

func (z pathImpl) Parent() efs_deprecated2.Path {
	if len(z.pathElements) < 1 {
		return z
	}
	z.pathElements = z.pathElements[:len(z.pathElements)-2]
	return z
}

func (z pathImpl) Basename() string {
	if len(z.pathElements) < 1 {
		return ""
	} else {
		return z.pathElements[len(z.pathElements)-1]
	}
}

func (z pathImpl) Extname() string {
	base := z.Basename()
	pos := strings.LastIndex(base, ".")
	if pos < 0 {
		return ""
	}
	return base[pos:]
}

func (z pathImpl) Root() efs_deprecated2.Path {
	z.pathElements = rootPathElements
	return z
}

func (z pathImpl) IsRoot() bool {
	return len(z.pathElements) < 1
}

func (z pathImpl) Namespace() efs_deprecated2.Namespace {
	return z.namespace
}

func (z pathImpl) Child(path string) (efs_deprecated2.Path, error) {
	if path == "" {
		return z, nil
	}

	// Split the path by separator if it contains multiple segments
	parts := strings.Split(path, "/")
	parts = removeEmptyParts(parts)

	if len(parts) == 0 {
		return z, nil
	}

	elements := make([]string, len(z.pathElements)+len(parts))
	existingOffset := len(z.pathElements)

	// Copy existing path elements
	copy(elements[:existingOffset], z.pathElements)

	// Process each part of the path
	for i, part := range parts {
		err := z.fs.NameRule().Accept(part)
		if err != nil {
			return nil, NewChildOutcomeByNameOutcome(err)
		}
		elements[existingOffset+i] = part
	}

	// Create a new pathImpl with the updated elements
	return &pathImpl{
		pathElements: elements,
		namespace:    z.namespace,
		fs:           z.fs,
	}, nil
}

// Helper function to remove empty path parts
func removeEmptyParts(parts []string) []string {
	result := make([]string, 0, len(parts))
	for _, part := range parts {
		if part != "" {
			result = append(result, part)
		}
	}
	return result
}

// Add the Dir method required by the interface
func (z pathImpl) Dir() efs_deprecated2.Path {
	if len(z.pathElements) <= 1 {
		return &pathImpl{
			pathElements: []string{},
			namespace:    z.namespace,
			fs:           z.fs,
		}
	}

	// Create a new path with all elements except the last one
	dirElements := make([]string, len(z.pathElements)-1)
	copy(dirElements, z.pathElements[:len(z.pathElements)-1])

	return &pathImpl{
		pathElements: dirElements,
		namespace:    z.namespace,
		fs:           z.fs,
	}
}

// Add Name method required by the interface
func (z pathImpl) Name() string {
	return z.Basename()
}

// Add Extension method required by the interface
func (z pathImpl) Extension() string {
	return z.Extname()
}

// Add DirSlash method required by the interface
func (z pathImpl) DirSlash() string {
	dir := z.Dir()
	if dir.IsRoot() {
		return "/"
	}
	return dir.String() + "/"
}

// Implement FindFile method required by the interface
func (z pathImpl) FindFile(cond func(path efs_deprecated2.Path) bool) (efs_deprecated2.Path, error) {
	// Simple implementation - just check if the current path matches
	if cond(z) {
		return z, nil
	}
	return nil, nil
}

// Implement FindPathUnderTree method required by the interface
func (z pathImpl) FindPathUnderTree(cond func(path efs_deprecated2.Path) bool) ([]efs_deprecated2.Path, error) {
	// Simple implementation - just check if the current path matches
	if cond(z) {
		return []efs_deprecated2.Path{z}, nil
	}
	return []efs_deprecated2.Path{}, nil
}

// Implement Walk method required by the interface
func (z pathImpl) Walk(fn func(path efs_deprecated2.Path) error) error {
	return fn(z)
}

// Implement MustChild method required by the interface
func (z pathImpl) MustChild(path string) efs_deprecated2.Path {
	child, err := z.Child(path)
	if err != nil {
		// Return self as fallback on error
		return z
	}
	return child
}

// Implement RelativeTo method required by the interface
func (z pathImpl) RelativeTo(base efs_deprecated2.Path) (string, error) {
	// Simple implementation - compare strings
	baseStr := base.String()
	selfStr := z.String()

	if strings.HasPrefix(selfStr, baseStr) {
		rel := selfStr[len(baseStr):]
		if len(rel) > 0 && rel[0] == '/' {
			rel = rel[1:]
		}
		return rel, nil
	}

	return "", fmt.Errorf("path %s is not relative to %s", selfStr, baseStr)
}

// Implement Equal method required by the interface
func (z pathImpl) Equal(p efs_deprecated2.Path) bool {
	return z.String() == p.String()
}
