package es_filecompare

import "github.com/watermint/toolbox/essentials/file/es_filesystem"

type FolderComparator interface {
	Compare(source, target es_filesystem.Path) (err error)

	CompareAndSummarize(source, target es_filesystem.Path) (missingSource, missingTarget []es_filesystem.Entry, fileDiffs, typeDiffs []EntryDataPair, err error)
}

type EntryDataPair struct {
	SourceData es_filesystem.EntryData `json:"source_data"`
	TargetData es_filesystem.EntryData `json:"target_data"`
}

type EntryPair struct {
	Source es_filesystem.Entry
	Target es_filesystem.Entry
}

type PathDataPair struct {
	SourceData es_filesystem.PathData `json:"source_data"`
	TargetData es_filesystem.PathData `json:"target_data"`
}

type PathPair struct {
	Source es_filesystem.Path
	Target es_filesystem.Path
}

type MissingSource func(base PathPair, target es_filesystem.Entry)
type MissingTarget func(base PathPair, source es_filesystem.Entry)
type TypeDiff func(base PathPair, source, target es_filesystem.Entry)
type FileDiff func(base PathPair, source, target es_filesystem.Entry)
type SameFile func(base PathPair, source, target es_filesystem.Entry)
type CompareFolder func(base PathPair, source, target es_filesystem.Entry)
