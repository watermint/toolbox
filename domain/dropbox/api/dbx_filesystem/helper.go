package dbx_filesystem

func NewHelper(rootInfo *RootInfo) *FileSystemBuilderHelper {
	return &FileSystemBuilderHelper{
		RootInfo: rootInfo,
	}
}

func NewEmptyHelper() *FileSystemBuilderHelper {
	return &FileSystemBuilderHelper{}
}

type FileSystemBuilderHelper struct {
	RootInfo *RootInfo
}

func (z *FileSystemBuilderHelper) HashSeed() []string {
	if z.RootInfo == nil {
		return []string{}
	}
	return []string{
		"r" + z.RootInfo.RootNamespaceId,
		"h" + z.RootInfo.HomeNamespaceId,
	}
}
