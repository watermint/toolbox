package em_tree

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
