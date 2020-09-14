package es_filecompare

import (
	"github.com/watermint/toolbox/essentials/file/es_filesystem"
)

type compareTime struct {
}

func (z compareTime) Compare(source, target es_filesystem.Entry) (same bool, err es_filesystem.FileSystemError) {
	return source.ModTime().Equal(target.ModTime()), nil
}

type compareSize struct {
}

func (z compareSize) Compare(source, target es_filesystem.Entry) (same bool, err es_filesystem.FileSystemError) {
	return source.Size() == target.Size(), nil
}

type compareContent struct {
}

func (z compareContent) Compare(source, target es_filesystem.Entry) (same bool, err es_filesystem.FileSystemError) {
	sourceHash, err := source.ContentHash()
	if err != nil {
		return false, err
	}
	targetHash, err := target.ContentHash()
	if err != nil {
		return false, err
	}
	return sourceHash == targetHash, nil
}

func New(opt ...Opt) FileComparator {
	opts := Opts{}.Apply(opt)
	comparators := make([]FileComparator, 0)
	comparators = append(comparators, &compareSize{})
	if !opts.dontCompareTime {
		comparators = append(comparators, &compareTime{})
	}
	if !opts.dontCompareContent {
		comparators = append(comparators, &compareContent{})
	}
	return &compareImpl{
		comparators: comparators,
	}
}

type compareImpl struct {
	comparators []FileComparator
}

func (z compareImpl) Compare(source, target es_filesystem.Entry) (same bool, err es_filesystem.FileSystemError) {
	for _, c := range z.comparators {
		same, err = c.Compare(source, target)
		if err != nil {
			return
		}
		if !same {
			return false, nil
		}
	}
	return true, nil
}
