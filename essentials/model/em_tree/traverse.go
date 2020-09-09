package em_tree

import (
	"path/filepath"
	"strings"
)

func SumNumNode(node Node) int {
	switch n := node.(type) {
	case File:
		return 1
	case Folder:
		descendants := 0
		for _, d := range n.Descendants() {
			descendants += SumNumNode(d)
		}
		return descendants + 1
	default:
		return 0
	}
}

func SumFileSize(node Node) int64 {
	switch n := node.(type) {
	case File:
		return n.Size()
	case Folder:
		var descendants int64
		for _, d := range n.Descendants() {
			descendants += SumFileSize(d)
		}
		return descendants
	default:
		return 0
	}
}

func MaxDepth(node Node) int {
	var depth func(node Node, current int) int
	depth = func(node Node, current int) int {
		switch n := node.(type) {
		case Folder:
			max := current
			for _, d := range n.Descendants() {
				dd := depth(d, current+1)
				if max < dd {
					max = dd
				}
			}
			return max
		default:
			return current
		}
	}
	return depth(node, 0)
}

// Returns a Node of the path. Returns nil if the node not found for the path.
func ResolvePath(node Node, path string) Node {
	cleanedPath := filepath.ToSlash(filepath.Clean(path))

	switch {
	case cleanedPath == ".", cleanedPath == "/":
		return node
	case strings.HasPrefix(cleanedPath, ".."):
		return nil
	}

	current := node
	fragments := strings.Split(cleanedPath, "/")
	numFragments := len(fragments)
	for i, f := range fragments {
		if f == "" {
			continue
		}

		switch n := current.(type) {
		case File:
			if current.Name() != f {
				return nil
			}
			if i != numFragments-1 {
				return nil
			}
			return n

		case Folder:
			found := false
			for _, d := range n.Descendants() {
				if d.Name() == f {
					current = d
					found = true
					break
				}
			}
			if !found {
				return nil
			}
			if i == numFragments-1 {
				return current
			}
		}
	}
	return nil
}

// Create folders. returns false if failed to create.
func CreateFolder(root Node, path string) bool {
	pathParts := make([]string, 0)
	pathParts = append(pathParts, "/")
	pathParts = append(pathParts, strings.Split(path, "/")...)
	parent := root.(Folder)

	for i := 1; i < len(pathParts); i++ {
		pp := pathParts[:i+1]
		p := filepath.Join(pp...)
		e := ResolvePath(root, p)
		if e == nil {
			current := NewFolder(pathParts[i], []Node{})
			parent.Add(current)
			parent = current
			continue
		}
		switch n := e.(type) {
		case Folder:
			parent = n
			continue
		case File:
			return false
		}
	}
	return true
}
